package controller

import (
	"fmt"
	. "go-api/model"
	r "go-api/routing"
)

type UserControllerStruct struct {
}

func (c *UserControllerStruct) Register(App *r.Router) {

	// List Users
	App.Get("/users", func(c *r.Context) error {
		size := 20
		p := 1
		if c.Params["page"] != nil {
			p = c.Params["page"].(int)
		}
		fmt.Printf("List Users %v %v\n", p, size)
		us, _ := User.Page(p, size).All()
		js := make([]map[string]interface{}, 0)
		for _, u := range us {
			j := make(map[string]interface{})
			j["id"] = u.ID
			j["name"] = u.Name
			j["created_at"] = u.CreatedAt
			js = append(js, j)
		}
		c.JSON(js)
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
