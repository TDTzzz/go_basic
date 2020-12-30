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

	r.GET("/hello/1", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})
	r.GET("/hello/2", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s", c.Query("name"), c.Path)
	})

	//r.GET("/hello/:name/doc", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})
	//r.GET("/hello/sssss", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//})

	//r.GET("/assets/*filepath/sss", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{
	//		"filepath": c.Param("filepath")})
	//})

	r.Run(":9999")
}
