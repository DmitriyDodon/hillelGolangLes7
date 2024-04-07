package server

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type httpServer struct {
	port int
}

func NewHttpServer(port int) *httpServer {
	return &httpServer{
		port: port,
	}
}

func (h httpServer) Start() error {
	http.HandleFunc("/", h.handle)
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), nil)
	if err != nil {
		return err
	}
	return nil
}

func (h httpServer) handle(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		h.handleGet(w, r)
	case "DELETE":
		h.handleDelete(w, r)
	case "POST":
		err = h.handlePost(w, r)
		if err != nil {
			log.Error(err.Error())
		}
	case "PUT":
		err = h.handlePut(w, r)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func (h httpServer) handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Value from query: %#v\n\n", r.URL.Query())
}

func (h httpServer) handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Value from query: %#v\n\n", r.URL.Query())
}

func (h httpServer) handlePost(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("Method: %s\n", r.Method)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Value from body: %s\n\n", string(body))

	return nil
}

func (h httpServer) handlePut(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("Method: %s\n", r.Method)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Value from body: %s\n\n", string(body))

	return nil
}
