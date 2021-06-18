package main

import (
	"fmt"
	"github.com/SeanChan0901/gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	stu1 := &student{Name: "Tom", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r := gee.Default()
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Gee-Web\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"gee-web"}
		c.String(http.StatusOK, names[100])
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "css.tmpl", nil)
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=gee-web
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/students", func(c *gee.Context) {
			c.HTML(http.StatusOK, "arr.tmpl", gee.H{
				"title":      "gee",
				"stuArr": [2]*student{stu1, stu2},
			})
		})
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/gee-web
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
