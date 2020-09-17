package controller

import (
	"fmt"
	. "go-api/model"
	r "go-api/routing"
	"strconv"
)

type UserControllerStruct struct {
}

func (c *UserControllerStruct) Register(App *r.Router) {

	// List Users
	App.Get("/users", func(c *r.Context) error {
		page, size := 1, 20
		if c.Params["page"] != nil {
			page, _ = strconv.Atoi(c.Params["page"].(string))
		}
		if c.Params["size"] != nil {
			size, _ = strconv.Atoi(c.Params["size"].(string))
		}
		us, _ := User.Page(page, size).All()
		ujs := make([]map[string]interface{}, 0)
		for _, v := range us {
			u := make(map[string]interface{})
			u["id"] = v.ID
			u["name"] = v.Name
			u["created_at"] = v.CreatedAt
			ujs = append(ujs, u)
		}
		j, _ := ResponseJSON(ujs)
		c.Write(j)
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
