package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}
	server := http.Server{Addr: "localhost:8080", Handler: handler}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Heloo world")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hai")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "image")
	})

	mux.HandleFunc("/images/thumbnail/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "thumbnail")
	})

	server := http.Server{Addr: "localhost:8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
