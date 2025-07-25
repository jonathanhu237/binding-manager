package apiserver

import (
	"encoding/json"
	"maps"
	"net/http"
)

func (as *ApiServer) writeJson(w http.ResponseWriter, status int, data any, headers http.Header) error {
	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
