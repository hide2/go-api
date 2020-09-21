package generator

import (
	"bytes"
	"flag"
	"fmt"
	. "go-api/lib"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

const GOCONTROLLERTEMPLATE = `
package controller

import (
	"fmt"
	. "go-api/model"
	r "go-api/routing"
	"strconv"
)

type {{.Model}}ControllerStruct struct {
}

func (c *{{.Model}}ControllerStruct) Register(App *r.Router) {
	// List {{.Model}}s
	App.Get("/{{.Paths}}", func(c *r.Context) error {
		page, size := 1, 20
		if c.Params["page"] != nil {
			page, _ = strconv.Atoi(c.Params["page"].(string))
		}
		if c.Params["size"] != nil {
			size, _ = strconv.Atoi(c.Params["size"].(string))
		}
		ms, _ := {{.Model}}.Page(page, size).All()
		c.ResponseJSON(ms)
		return nil
	})

	// Create New {{.Model}}
	App.Post("/{{.Paths}}", func(c *r.Context) error {
		props := c.Params
		m, _ := {{.Model}}.Create(props)
		c.ResponseJSON(m)
		return nil
	})

	// Get {{.Model}}
	App.Get("/{{.Paths}}/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		m, _ := {{.Model}}.Find(int64(id))
		c.ResponseJSON(m)
		return nil
	})

	// Update {{.Model}}
	App.Put("/{{.Paths}}/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		props := c.Params
		conds := map[string]interface{}{"id": int64(id)}
		{{.Model}}.Update(props, conds)
		m, _ := {{.Model}}.Find(int64(id))
		c.ResponseJSON(m)
		return nil
	})

	// Delete {{.Model}}
	App.Delete("/{{.Paths}}/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		{{.Model}}.Destroy(int64(id))
		m := make(map[string]interface{})
		c.ResponseJSON(m)
		return nil
	})

	fmt.Println("{{.Model}}Controller Registered.")
}

var {{.Model}}Controller = &{{.Model}}ControllerStruct{}
`

type ControllerTemplateAttr struct {
	Model string
	Paths string
	Attrs []string
	Keys  []string
}

func GenController() {
	flag.Parse()
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	mf, _ := ioutil.ReadFile(*inputConfigFile)
	ms := make(map[string][]yaml.MapSlice)
	merr := yaml.Unmarshal(mf, &ms)
	if merr != nil {
		fmt.Println("error:", merr)
	}
	for _, j := range ms["models"] {
		var modelname, paths, filename string
		attrs := make([]string, 0)
		keys := make([]string, 0)
		attrs = append(attrs, "ID")
		keys = append(keys, "id")
		for _, v := range j {
			if v.Key != "model" {
				attrs = append(attrs, Camelize(v.Key.(string)))
				keys = append(keys, v.Key.(string))
			} else {
				modelname = v.Value.(string)
				paths = strings.ToLower(modelname) + "s"
				filename = "controller/" + modelname + "Controller.go"
			}
		}
		fmt.Println("-- Generate", filename)
		t, err := template.New("GOCONTROLLERTEMPLATE").Parse(GOCONTROLLERTEMPLATE)
		if err != nil {
			fmt.Println(err)
			return
		}

		m := ControllerTemplateAttr{modelname, paths, attrs, keys}
		var b bytes.Buffer
		t.Execute(&b, m)
		// fmt.Println(b.String())

		// Write to file
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println("create file: ", err)
			return
		}
		err = t.Execute(f, m)
		if err != nil {
			fmt.Print("execute: ", err)
			return
		}
		f.Close()

	}
}
