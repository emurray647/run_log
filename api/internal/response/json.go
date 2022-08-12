package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type outputType struct {
	Err    *errorType  `json:"error"`
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type errorType struct {
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	output := outputType{
		Err:    nil,
		Status: statusCode,
		Data:   v,
	}
	b, err := json.Marshal(output)
	if err != nil {
		return fmt.Errorf("failed to marshal to json: %w", err)
	}

	w.Write(b)

	return nil
}

func WriteJSONError(w http.ResponseWriter, statusCode int, message string) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	output := outputType{
		Err: &errorType{
			Message: message,
		},
		Status: statusCode,
		Data:   nil,
	}
	b, err := json.Marshal(output)
	if err != nil {
		return fmt.Errorf("failed to marshal to json: %w", err)
	}

	w.Write(b)

	return nil
}
