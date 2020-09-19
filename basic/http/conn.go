package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		hello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!!!")
}
