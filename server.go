package main

import (
	"fmt"
	"time"

	. "go-api/controller"
	r "go-api/routing"

	"github.com/valyala/fasthttp"
)

func main() {

	// User.CreateTable()
	// Event.CreateTable()

	var App = r.New()
	App.Register(UserController)
	App.Register(EventController)

	fmt.Printf("[%s] ===== Server Started at 8080.\n", time.Now().Format("2006-01-02 15:04:05"))
	panic(fasthttp.ListenAndServe(":8080", App.HandleRequest))
}
