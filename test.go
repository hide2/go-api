package main

import (
	"fmt"
	r "go-api/fasthttp-routing"

	"github.com/valyala/fasthttp"
)

func main() {
	router := r.New()

	// List Users
	router.Get("/users", func(c *r.Context) error {
		fmt.Printf("List Users %v\n", c.Param("id"))
		return nil
	})

	// Create New User
	router.Post("/users", func(c *r.Context) error {
		fmt.Printf("\nCreate New User - URI %s\n", c.RequestURI())
		fmt.Printf("Create New User - Header %s\n", c.Request.Header.Peek("Content-Type"))
		fmt.Printf("Create New User - Path %s\n", c.Path())
		fmt.Printf("Create New User - Params %v\n", c.Params)
		return nil
	})

	// Get User
	router.Get("/users/<id>", func(c *r.Context) error {
		fmt.Printf("Get User %v\n", c.NamedParams["id"])
		fmt.Printf("Get User - NamedParams %v\n", c.NamedParams)
		return nil
	})

	// Update User
	router.Put("/users/<id>", func(c *r.Context) error {
		fmt.Printf("Update User %v %v %v\n", c.Param("id"), c.Get("aaa"), c.Get("bbb"))
		return nil
	})

	// Delete User
	router.Delete("/users/<id>", func(c *r.Context) error {
		fmt.Printf("Delete User %v\n", c.Param("id"))
		return nil
	})

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}
