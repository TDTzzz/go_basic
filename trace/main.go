package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/async", serviceHandler)
	mux.HandleFunc("/service", serviceHandler)
	mux.HandleFunc("/db", dbHandler)
	fmt.Printf("Go to http://localhost:%d/home to start a request!\n", port)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// Acts as our index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="/home"> Click here to start a request </a>`))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//
	span := opentracing.StartSpan("/home")
	defer span.Finish()

	w.Write([]byte("Request started"))
	go func() {
		http.Get("http://localhost:8080/async")
	}()
	_, err := http.Get("http://localhost:8080/service")
	if err != nil {
		ext.Error.Set(span, true)
		span.LogEventWithPayload("GET SERVICE ERR", err)
	}
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	w.Write([]byte("Request done!"))
}

// Mocks a service endpoint that makes a DB call
func serviceHandler(w http.ResponseWriter, r *http.Request) {
	// ...
	http.Get("http://localhost:8080/db")
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	// ...
}

// Mocks a DB call
func dbHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	// here would be the actual call to a DB.
}
