
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
		ms, _ := User.Page(page, size).All()
		c.ResponseJSON(ms)
		return nil
	})

	// Create New User
	App.Post("/users", func(c *r.Context) error {
		props := c.Params
		m, _ := User.Create(props)
		c.ResponseJSON(m)
		return nil
	})

	// Get User
	App.Get("/users/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		m, _ := User.Find(int64(id))
		c.ResponseJSON(m)
		return nil
	})

	// Update User
	App.Put("/users/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		props := c.Params
		conds := map[string]interface{}{"id": int64(id)}
		User.Update(props, conds)
		m, _ := User.Find(int64(id))
		c.ResponseJSON(m)
		return nil
	})

	// Delete User
	App.Delete("/users/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		User.Destroy(int64(id))
		m := make(map[string]interface{})
		c.ResponseJSON(m)
		return nil
	})

	fmt.Println("UserController Registered.")
}

var UserController = &UserControllerStruct{}
