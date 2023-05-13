package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nironwp/hexagonal-architecture/adapter/server/handler"
	"github.com/nironwp/hexagonal-architecture/application"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer(service application.ProductServiceInterface) *WebServer {
	return &WebServer{
		Service: service,
	}
}

func (w WebServer) Server() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)

	http.Handle("/", r)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
