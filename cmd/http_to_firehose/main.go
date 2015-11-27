package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/gorilla/handlers"
	"github.com/tj/go-config"
	"github.com/tj/http_to_firehose/client"
	"github.com/tj/http_to_firehose/server"
	"github.com/tj/http_to_firehose/server/basicauth"
)

type Options struct {
	Address    string `help:"Bind address"`
	StreamName string `help:"Firehose stream name"`
	Backlog    int    `help:"Firehose record backlog size"`
	Username   string `help:"Basic auth username"`
	Password   string `help:"Basic auth password"`
}

func main() {
	var options = Options{
		Address: ":3000",
		Backlog: 100,
	}

	config.MustResolve(&options)

	c := &client.Client{
		Firehose:   firehose.New(session.New()),
		StreamName: options.StreamName,
		Backlog:    make(chan []byte, options.Backlog),
	}

	s := &server.Server{
		Client: c,
	}

	c.Start()

	var h http.Handler = s

	if options.Username != "" && options.Password != "" {
		h = basicauth.BasicAuth{
			Username: options.Username,
			Password: options.Password,
			Handler:  h,
		}
	}

	h = handlers.LoggingHandler(os.Stderr, h)

	log.Fatalln(http.ListenAndServe(options.Address, h))
}
