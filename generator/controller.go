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
		us, _ := {{.Model}}.Page(page, size).All()
		ujs := make([]map[string]interface{}, 0)
		for _, v := range us {
			u := make(map[string]interface{})
			{{- range $i, $k := .Keys }}
			u["{{$k}}"] = v.{{index $.Attrs $i}}
			{{- end }}
			ujs = append(ujs, u)
		}
		j, _ := ResponseJSON(ujs)
		c.Write(j)
		return nil
	})

	// Create New {{.Model}}
	App.Post("/{{.Paths}}", func(c *r.Context) error {
		props := c.Params
		u, _ := {{.Model}}.Create(props)
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

	// Get {{.Model}}
	App.Get("/{{.Paths}}/<id>", func(c *r.Context) error {
		id, _ := strconv.Atoi(c.NamedParams["id"].(string))
		u, _ := {{.Model}}.Find(int64(id))
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

	// Update {{.Model}}
	App.Put("/{{.Paths}}/<id>", func(c *r.Context) error {
		c.JSON(c.NamedParams)
		return nil
	})

	// Delete {{.Model}}
	App.Delete("/{{.Paths}}/<id>", func(c *r.Context) error {
		c.JSON(c.NamedParams)
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
