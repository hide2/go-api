package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	app := App()
	app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
		ctx.SetBody([]byte("666"))
	})
	server := &fasthttp.Server{
		Handler:     app.Handler,
		Concurrency: 1024 * 1024,
	}

	fmt.Println("API started at :9798")
	server.ListenAndServe(":9798")
}
