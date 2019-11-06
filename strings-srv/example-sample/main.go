package main

import (
	"github.com/hailinluo/gokit-examples/strings-srv/example-sample/endpoint"
	"github.com/hailinluo/gokit-examples/strings-srv/example-sample/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hailinluo/gokit-examples/strings-srv/example-sample/transport"
	"log"
	"net/http"
)

func main() {
	svc := service.StringService{}

	uppercaseHandler := httptransport.NewServer(
		endpoint.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		endpoint.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
