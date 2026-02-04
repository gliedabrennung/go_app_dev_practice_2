package handlers

import (
	"encoding/json"
	"net/http"
)

func decodeJSON[T any](r *http.Request) (T, error) {
	var payload T

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		return payload, err
	}

	return payload, nil
}
