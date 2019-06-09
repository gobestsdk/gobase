package httpserver

import (
	"net/http"
	"strings"
)

const MethodSearch = "SEARCH"
const MethodCreate = "CREATE"

func AccessControlAllowMethods() string {
	var method = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPost,
		http.MethodOptions,
		MethodCreate,
		MethodSearch,
	}
	return strings.Join(method, ",")
}
