package responder

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type genericResponse struct {
	Error   bool        `json:"error"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Bind(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, data any, status int, headers ...http.Header) error {
	response := genericResponse{
		Error:  false,
		Status: status,
		Data:   data,
	}

	out, err := json.Marshal(response)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	response := genericResponse{
		Error:   true,
		Status:  statusCode,
		Message: err.Error(),
	}

	return WriteJSON(w, response, statusCode)
}
