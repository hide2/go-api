package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)


func main() {
	app := App()
	c := Redis("localhost:6379")

	app.Use("/ping", func(ctx *fasthttp.RequestCtx, next func(error)) {
		pong, _ := c.Ping().Result()
		ctx.SetBody([]byte(pong))
	})

	app.Use("/key", func(ctx *fasthttp.RequestCtx, next func(error)) {
		c.Set("key", "123456", 0)
		val, _ := c.Get("key").Result()
		ctx.SetBody([]byte(val))
	})

	app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
		ctx.SetBody([]byte("666"))
	})

	server := &fasthttp.Server{
		Handler:     app.Handler,
		Concurrency: 1024 * 1024,
	}

	fmt.Println("Server started at :9798")
	server.ListenAndServe(":9798")
}
