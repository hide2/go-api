package main

import (
	"fmt"
	r "go-api/routing"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	router := r.New()

	// List Users
	router.Get("/users", func(c *r.Context) error {
		fmt.Printf("List Users %v\n", c.Param("id"))
		fmt.Printf("Create New User - Path %s\n", c.Path())
		fmt.Printf("Create New User - Header %s\n", c.Request.Header.Peek("Content-Type"))
		c.JSON(c.Params)
		return nil
	})

	// Create New User
	router.Post("/users", func(c *r.Context) error {
		fmt.Printf("Create New User - Params %v\n", c.Params)
		c.JSON(c.Params)
		return nil
	})

	// Get User
	router.Get("/users/<id>", func(c *r.Context) error {
		fmt.Printf("Get User %v\n", c.NamedParams["id"])
		fmt.Printf("Get User - NamedParams %v\n", c.NamedParams)
		c.JSON(c.NamedParams)
		return nil
	})

	// Update User
	router.Put("/users/<user_id>", func(c *r.Context) error {
		fmt.Printf("Update User %v\n", c.NamedParams["user_id"])
		c.JSON(c.NamedParams)
		return nil
	})

	// Delete User
	router.Delete("/users/<id>", func(c *r.Context) error {
		fmt.Printf("Delete User %v\n", c.NamedParams["id"])
		c.JSON(c.NamedParams)
		return nil
	})

	fmt.Printf("[%s] ===== Go-API Server Started at 8080.\n", time.Now().Format("2006-01-02 15:04:05"))
	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}
