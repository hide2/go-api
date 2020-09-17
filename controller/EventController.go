
package controller

import (
	"fmt"
	. "go-api/model"
	r "go-api/routing"
	"strconv"
)

type EventControllerStruct struct {
}

func (c *EventControllerStruct) Register(App *r.Router) {
	// List Events
	App.Get("/events", func(c *r.Context) error {
		page, size := 1, 20
		if c.Params["page"] != nil {
			page, _ = strconv.Atoi(c.Params["page"].(string))
		}
		if c.Params["size"] != nil {
			size, _ = strconv.Atoi(c.Params["size"].(string))
		}
		us, _ := Event.Page(page, size).All()
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

	// Create New Event
	App.Post("/events", func(c *r.Context) error {
		props := c.Params
		u, _ := Event.Create(props)
		ujs := make(map[string]interface{})
		if u != nil {
			ujs["id"] = u.ID
			ujs["name"] = u.Name
			ujs["created_at"] = u.CreatedAt
		}
		j, _ := ResponseJSON(ujs)
		c.Write(j)
		return nil
	})

	// Get Event
	App.Get("/events/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		u, _ := Event.Find(int64(id))
		ujs := make(map[string]interface{})
		if u != nil {
			ujs["id"] = u.ID
			ujs["name"] = u.Name
			ujs["created_at"] = u.CreatedAt
		}
		j, _ := ResponseJSON(ujs)
		c.Write(j)
		return nil
	})

	// Update Event
	App.Put("/events/<id>", func(c *r.Context) error {
		c.JSON(c.NamedParams)
		return nil
	})

	// Delete Event
	App.Delete("/events/<id>", func(c *r.Context) error {
		c.JSON(c.NamedParams)
		return nil
	})

	fmt.Println("EventController Registered.")
}

var EventController = &EventControllerStruct{}
