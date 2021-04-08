package main

import (
	"net/http"
	"play-go/blue"
	"play-go/blue/context"
)

func main() {
	enginee := blue.New()
	enginee.GET("/index", func(c *context.Context) {
		c.String(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := enginee.Group("/v1")
	{
		v1.GET("/", func(c *context.Context) {
			c.String(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *context.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := enginee.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *context.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *context.Context) {
			result := map[string]string{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			}
			c.JSON(http.StatusOK, result)
		})

	}
	enginee.Run(":8080")
}
