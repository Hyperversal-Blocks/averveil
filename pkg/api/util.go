package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
