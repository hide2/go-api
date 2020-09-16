package main

import (
	"fmt"
	"time"

	. "go-api/controller"
	r "go-api/routing"

	"github.com/valyala/fasthttp"
)

func main() {

	var App = r.New()
	InitUserController(App)
	InitEventController(App)

	fmt.Printf("[%s] ===== Go-API Server Started at 8080.\n", time.Now().Format("2006-01-02 15:04:05"))
	panic(fasthttp.ListenAndServe(":8080", App.HandleRequest))
}
