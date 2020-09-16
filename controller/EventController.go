package controller

import (
	"fmt"
	r "go-api/routing"
)

type EventControllerStruct struct {
}

func (c *EventControllerStruct) Register(App *r.Router) {
	// List Events
	App.Get("/events", func(c *r.Context) error {
		fmt.Printf("List Events %v\n", c.Params["page"])
		c.JSON(c.Params)
		return nil
	})

	fmt.Println("EventController Registered.")
}

var EventController = &EventControllerStruct{}
