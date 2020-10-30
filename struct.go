package request

import "net/http"

type Response struct {
	Status     string
	StatusCode int
	Body       string
	Header     http.Header
}
