package main

import (
	"github.com/SeanChan0901/gee-web/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context){
		c.String(http.StatusOK, "Hello Gee-Web\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context){
		names := []string{"gee-web"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
