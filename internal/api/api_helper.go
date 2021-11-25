package api

import (
	"encoding/json"
	"net/http"
)

// ResponseStatus constants
const (
	StatusOK   = "ok"
	StatusFail = "nok"
)

// Response implements standard JSON response payload structure.
type Response struct {
	Status string          `json:"status"`
	Error  *ResponseError  `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

// ResponseError implements the standard Error structure to return in response payloads.
type ResponseError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e ResponseError) Error() string {
	j, err := json.Marshal(e)
	if err != nil {
		return "ResponseError: " + err.Error()
	}
	return string(j)
}

// fail ends an unsuccessful JSON response with the standard
// failure format for services.
func fail(w http.ResponseWriter, status int, details ...string) {
	msg := http.StatusText(status)
	r := &Response{
		Status: StatusFail,
		Error: &ResponseError{
			Code:    status,
			Message: msg,
			Details: details,
		},
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

// send sends a successful JSON response
func send(w http.ResponseWriter, status int, result interface{}) {
	rj, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	r := &Response{
		Status: StatusOK,
		Result: rj,
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
