package main

import (
	"fmt"
	"go_basic/7days/gee-web/day1-http-base/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.Run(":9999")
}
