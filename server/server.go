// Package server implements an HTTP server which then
// transfers the request bodies to the Firehose client.
package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client interface.
type client interface {
	Put(b []byte) error
}

// Server relaying request bodies to Firehose.
type Server struct {
	Client client
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %s", err)
		return
	}

	if err := s.Client.Put(b); err != nil {
		log.Printf("error sending record:", err)
	}

	w.WriteHeader(202)
	fmt.Fprintln(w, http.StatusText(202))
}
