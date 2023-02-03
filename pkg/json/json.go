package json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONSerializer struct{}

// Send(w io.Writer, status int, payload interface{}) error
// Read(r io.Reader, dest interface{}) error

func (j *JSONSerializer) Send(w http.ResponseWriter, status int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	fmt.Println(w.Header().Get("Content-Type"))
	return json.NewEncoder(w).Encode(payload)
}

func (j *JSONSerializer) Read(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(dest)
}
