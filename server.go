package main

import (
	"fmt"
	"time"

	. "go-api/controller"
	. "go-api/model"
	r "go-api/routing"

	"github.com/valyala/fasthttp"
)

func main() {

	User.Exec("DROP TABLE IF EXISTS user")
	Event.Exec("DROP TABLE IF EXISTS event")
	User.CreateTable()
	Event.CreateTable()
	for i := 0; i < 30; i++ {
		props := map[string]interface{}{"name": "Calvin"}
		User.Create(props)
	}

	var App = r.New()
	App.Register(UserController)
	App.Register(EventController)

	fmt.Printf("[%s] ===== Server Started at 8080.\n", time.Now().Format("2006-01-02 15:04:05"))
	panic(fasthttp.ListenAndServe(":8080", App.HandleRequest))
}
