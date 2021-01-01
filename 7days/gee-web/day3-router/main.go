package main

import (
	"go_basic/7days/gee-web/day3-router/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Ha Ha Ha</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})

	r.GET("/hello/v1", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})
	r.GET("/hello/v1/list", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})
	r.GET("/hello/v1/detail", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})
	r.GET("/hello/v2", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})
	r.GET("/hello/v2/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})
	r.Run(":9999")
}
