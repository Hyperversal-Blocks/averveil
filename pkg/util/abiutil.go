package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ContractParser parses the contract address and jsonABI into appropriate readable format
func ContractParser(address string, jsonABI interface{}) (common.Address, abi.ABI, error) {
	jsonMarshaledABI, err := json.Marshal(jsonABI)
	if err != nil {
		return common.Address{}, abi.ABI{}, fmt.Errorf("unable to marshal json: %w", err)
	}

	jsonToABI, err := abi.JSON(strings.NewReader(string(jsonMarshaledABI)))
	if err != nil {
		return common.Address{}, abi.ABI{}, fmt.Errorf("unable to parse ABI: %w", err)
	}

	return common.HexToAddress(address), jsonToABI, nil
}

func WriteJson(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	marshalledJson, e := json.MarshalIndent(v, "", "    ")
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(fmt.Errorf("unable to marshall json"))
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(statusCode)
	_, err := w.Write(marshalledJson)
	if err != nil {
		return
	}
}
