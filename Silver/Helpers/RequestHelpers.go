package Helpers

import "net/http"

var (
	CurrentResponseWriter http.ResponseWriter
	CurrentRequest        *http.Request
)
