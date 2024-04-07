package server

import (
	"fmt"
)

type request struct {
	method      string
	body        string
	queryParams map[string]string
}

func NewRequest(method string, body string, queryParams map[string]string) *request {
	return &request{
		method:      method,
		body:        body,
		queryParams: queryParams,
	}
}

func (r request) handleGet() {
	fmt.Printf("Method: %s\n", r.method)
	fmt.Printf("Value from query: %#v\n\n", r.queryParams)
}

func (r request) handleDelete() {
	fmt.Printf("Method: %s\n", r.method)
	fmt.Printf("Value from query: %#v\n\n", r.queryParams)
}

func (r request) handlePost() {
	fmt.Printf("Method: %s\n", r.method)
	fmt.Printf("Value from body: %s\n\n", string(r.body))
}

func (r request) handlePut() {
	fmt.Printf("Method: %s\n", r.method)
	fmt.Printf("Value from body: %s\n\n", string(r.body))
}
