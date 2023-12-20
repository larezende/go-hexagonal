package server

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"go-hexagonal/adapters/web/handler"
	"go-hexagonal/application"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver(service application.ProductServiceInterface) *WebServer {
	return &WebServer{
		Service: service,
	}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductWebHandler(r, n, w.Service)
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
