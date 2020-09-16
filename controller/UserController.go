package controller

import (
	"fmt"
	r "go-api/routing"
)

type UserControllerStruct struct {
}

func (c *UserControllerStruct) Register(App *r.Router) {

	// List Users
	App.Get("/users", func(c *r.Context) error {
		fmt.Printf("List Users %v\n", c.Params["page"])
		c.JSON(c.Params)
		return nil
	})

	// Create New User
	App.Post("/users", func(c *r.Context) error {
		fmt.Printf("Create New User - Params %v\n", c.Params)
		c.JSON(c.Params)
		return nil
	})

	// Get User
	App.Get("/users/<id>", func(c *r.Context) error {
		fmt.Printf("Get User %v\n", c.NamedParams["id"])
		c.JSON(c.NamedParams)
		return nil
	})

	// Update User
	App.Put("/users/<user_id>", func(c *r.Context) error {
		fmt.Printf("Update User %v\n", c.NamedParams["user_id"])
		c.JSON(c.NamedParams)
		return nil
	})

	// Delete User
	App.Delete("/users/<id>", func(c *r.Context) error {
		fmt.Printf("Delete User %v\n", c.NamedParams["id"])
		c.JSON(c.NamedParams)
		return nil
	})

	fmt.Println("UserController Registered.")
}

var UserController = &UserControllerStruct{}
