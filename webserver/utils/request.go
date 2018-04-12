package utils

import (
	"net/http"
)

// QueryParams retrieves the requests querystring values for a given key
func QueryParams(key string, r *http.Request) []string {
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys) <= 0 {
		return []string{}
	}
	return keys
}

// QueryParam retrieves the requests querystring value for a given key
func QueryParam(key string, r *http.Request) (string, bool) {
	params := QueryParams(key, r)
	if len(params) <= 0 {
		return "", false
	}
	return params[0], true
}
