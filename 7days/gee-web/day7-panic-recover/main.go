package main

import (
	"go_basic/7days/gee-web/day7-panic-recover/gee"
	"net/http"
)

func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"tdtzzz"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
