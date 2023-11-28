package swarm

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type swarm struct {
	addr string
}

type HealthStatusResponse struct {
	Status          string `json:"status"`
	Version         string `json:"version"`
	APIVersion      string `json:"apiVersion"`
	DebugAPIVersion string `json:"debugApiVersion"`
}

type ReferenceResponse struct {
	Reference string `json:"reference"`
}

type DebugPostageAllBatchesResponse struct {
	Batches []PostageBatchShort `json:"batches"`
}

type PostageBatchShort struct {
	BatchID       string `json:"batchID"`
	Value         string `json:"value"` // BigInt is represented as a string in JSON
	Start         int    `json:"start"`
	Depth         int    `json:"depth"`
	BucketDepth   int    `json:"bucketDepth"`
	ImmutableFlag bool   `json:"immutableFlag"`
	BatchTTL      int    `json:"batchTTL"`
	Owner         string `json:"owner"`
	StorageRadius int    `json:"storageRadius"`
}

type PostageBatch struct {
	BatchID       string `json:"batchID"`
	Utilization   int    `json:"utilization"`
	Usable        bool   `json:"usable"`
	Label         string `json:"label"`
	Depth         int    `json:"depth"`
	Amount        string `json:"amount"` // BigInt is represented as a string in JSON
	BucketDepth   int    `json:"bucketDepth"`
	BlockNumber   int    `json:"blockNumber"`
	ImmutableFlag bool   `json:"immutableFlag"`
	Exists        bool   `json:"exists"`
	BatchTTL      int    `json:"batchTTL"`
}

func (s *swarm) CheckNodeHealthAndReadiness() (*HealthStatusResponse, error) {
	url := "http://localhost:1635/health"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	m := new(HealthStatusResponse)

	_ = json.Unmarshal(body, &m)

	return m, nil
}

func (s *swarm) BuyPostageStamp(baseURL string, bearerToken string, amount string, depth int, label string, immutable bool) (string, error) {
	// Construct the URL with path parameters
	url := fmt.Sprintf("%s/stamps/%s/%d", baseURL, amount, depth)

	// Create a client
	client := &http.Client{}

	// Create a request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	// Add query parameters if needed
	q := req.URL.Query()
	if label != "" {
		q.Add("label", label)
	}
	if immutable {
		q.Add("immutable", "true")
	}
	req.URL.RawQuery = q.Encode()

	// Add headers
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check if the status code is in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("server returned non-2xx status: %d - body: %s", resp.StatusCode, body)
	}

	// Return the response body
	return string(body), nil
}

func (s *swarm) TopUpPostageBatch(baseURL, bearerToken, batchID string, amount int) (string, error) {
	// Construct the URL with the batch ID and amount as path parameters
	url := fmt.Sprintf("%s/stamps/topup/%s/%d", baseURL, batchID, amount)

	// Create an HTTP client
	client := &http.Client{}

	// Create a PATCH request
	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		return "", err
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check if the status code is in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("server returned non-2xx status: %d - body: %s", resp.StatusCode, body)
	}

	// Return the response body as a string
	return string(body), nil
}

func (s *swarm) GetPostageBatchStatus(baseURL, bearerToken, batchID string) (*PostageBatch, error) {
	// Construct the URL with the batch ID as a path parameter
	url := fmt.Sprintf("%s/stamps/%s", baseURL, batchID)

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if the status code is in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("server returned non-2xx status: %d - body: %s", resp.StatusCode, body)
	}

	// Unmarshal the JSON response into the PostageBatch struct
	var batch PostageBatch
	err = json.Unmarshal(body, &batch)
	if err != nil {
		return nil, err
	}

	return &batch, nil
}

func (s *swarm) Upload(baseURL, bearerToken, swarmPostageBatchId string, data io.Reader) (*ReferenceResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("%s/bytes", baseURL)

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, err
	}

	// Set the necessary headers
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("swarm-postage-batch-id", swarmPostageBatchId)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if the status code is 201 (Created)
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("server returned non-201 status: %d - body: %s", resp.StatusCode, body)
	}

	// Unmarshal the JSON response into the ReferenceResponse struct
	var referenceResponse ReferenceResponse
	err = json.Unmarshal(body, &referenceResponse)
	if err != nil {
		return nil, err
	}

	return &referenceResponse, nil
}

func (s *swarm) GetAllBatches(baseURL, bearerToken string) (*DebugPostageAllBatchesResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("%s/batches", baseURL)

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if the status code is 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-200 status: %d - body: %s", resp.StatusCode, body)
	}

	// Parse the JSON response
	var response DebugPostageAllBatchesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *swarm) GetBalance() error {
	// TODO implement me
	panic("implement me")
}

type Swarm interface {
	CheckNodeHealthAndReadiness() (*HealthStatusResponse, error)
	BuyPostageStamp(baseURL string, bearerToken string, amount string, depth int, label string, immutable bool) (string, error)
	GetPostageBatchStatus(baseURL, bearerToken, batchID string) (*PostageBatch, error)
	TopUpPostageBatch(baseURL, bearerToken, batchID string, amount int) (string, error)
	Upload(baseURL, bearerToken, swarmPostageBatchId string, data io.Reader) (*ReferenceResponse, error)
	GetAllBatches(baseURL, bearerToken string) (*DebugPostageAllBatchesResponse, error)
	GetBalance() error
}

func New(addr string) Swarm {
	return &swarm{addr: addr}
}

// TODO: https://github.com/ethersphere/bee/tree/master/openapi
