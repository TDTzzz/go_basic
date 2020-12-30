package main

import (
	"go_basic/7days/gee-web/day5-middleware/gee"
	"log"
	"net/http"
	"time"
)

func main() {

	r := gee.New()
	r.Use(gee.Logger())
	//r.GET("/", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello</h1>")
	//})
	v2 := r.Group("/v2")
	v2.Use(midForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}

func midForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
