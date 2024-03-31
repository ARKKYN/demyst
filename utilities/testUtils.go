package utilities

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func CreateMockServer(status int, response interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		if response != nil {
			json.NewEncoder(w).Encode(response)
		}
	}))
	return server
}
