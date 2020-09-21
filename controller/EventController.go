
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
		ms, _ := Event.Page(page, size).All()
		c.ResponseJSON(ms)
		return nil
	})

	// Create New Event
	App.Post("/events", func(c *r.Context) error {
		props := c.Params
		m, _ := Event.Create(props)
		c.ResponseJSON(m)
		return nil
	})

	// Get Event
	App.Get("/events/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		m, _ := Event.Find(int64(id))
		c.ResponseJSON(m)
		return nil
	})

	// Update Event
	App.Put("/events/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		props := c.Params
		conds := map[string]interface{}{"id": int64(id)}
		Event.Update(props, conds)
		m, _ := Event.Find(int64(id))
		c.ResponseJSON(m)
		return nil
	})

	// Delete Event
	App.Delete("/events/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		Event.Destroy(int64(id))
		m := make(map[string]interface{})
		c.ResponseJSON(m)
		return nil
	})

	fmt.Println("EventController Registered.")
}

var EventController = &EventControllerStruct{}
