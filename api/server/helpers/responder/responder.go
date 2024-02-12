package responder

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type genericResponse struct {
	Error   bool        `json:"error"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, data any, status int, headers ...http.Header) {
	response := genericResponse{
		Error:  false,
		Status: status,
		Data:   data,
	}

	out, err := json.Marshal(response)
	if err != nil {
		slog.Error("Error in Marshal:", err)
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
		slog.Error("Error in Write:", err)
	}
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	response := genericResponse{
		Error:   true,
		Status:  statusCode,
		Message: err.Error(),
	}

	WriteJSON(w, response, statusCode)
}
