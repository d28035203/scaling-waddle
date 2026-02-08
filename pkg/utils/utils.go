package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody decodes a JSON request body into dest.
func ParseBody(r *http.Request, dest interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.Unmarshal(body, dest)
}
