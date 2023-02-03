package json

import (
	"encoding/json"
	"io"
	"net/http"
)

type JSONSerializer struct{}

// Send(w io.Writer, status int, payload interface{}) error
// Read(r io.Reader, dest interface{}) error

func (j *JSONSerializer) Send(w http.ResponseWriter, status int, payload interface{}) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}

func (j *JSONSerializer) Read(r io.Reader, dest interface{}) error {
	return json.NewDecoder(r).Decode(dest)
}
