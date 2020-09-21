Go-API is a micro API service framework, auto-generate RESTful API based on model yml

# Go-API Features
- Restful JSON API generator
- Auto support for application/x-www-form-urlencoded & application/json
- NamedParams & Params(Query Args/POST Body)
- Routing & Before/After Filter
- Access Log
- MySQL based ORM Model & CRUD methods generator
- Auto create table
- Transaction
- Pagination
- Connection Pool
- Write/Read Splitting
- Multi datasources
- Auto/Customized mapping of Model and datasource/table
- SQL log & Slow SQL log for profiling

# Install Go on Mac
``` bash
sudo rm -fr /usr/local/go
Download & Install MacOS pkg from https://golang.org/dl/
export PATH=$PATH:/usr/local/go/bin
```

# Install Go on CentOS
``` bash
wget https://golang.org/dl/go1.15.1.linux-amd64.tar.gz
sudo tar -C /usr/local/ -xzvf go1.15.1.linux-amd64.tar.gz
sudo vi /etc/profile
export GOROOT=/usr/local/go
export GOPATH=/data/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
```
# Usage
Define Datasources in datasource.yml
``` yml
datasources:
  - name: default
    write: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True
    read: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True

  - name: ds_2
    write: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True
    read: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True

sql_log: false
slow_sql_log: 500

redis_host: 127.0.0.1
redis_port: 6379
redis_password: 
```
Define Models in model.yml
``` yml
models:
  - model: User
    name: string
    created_at: time.Time

  - model: Event
    name: string
    created_at: time.Time
```
Generate Model & Router & Controller go files
``` bash
> go run gen.go
> -- Generate model/User.go
> -- Generate model/Event.go
> -- Generate controller/UserController.go
> -- Generate controller/EventController.go
```
Start Server with go run server.go
``` go
package main

import (
	"fmt"
	"time"

	. "go-api/controller"
	. "go-api/model"
	r "go-api/routing"

	"github.com/valyala/fasthttp"
)

func main() {

	User.Exec("DROP TABLE IF EXISTS user")
	Event.Exec("DROP TABLE IF EXISTS event")
	User.CreateTable()
	Event.CreateTable()
	for i := 0; i < 30; i++ {
		props := map[string]interface{}{"name": "Calvin"}
		User.Create(props)
	}

	var App = r.New()
	App.Register(UserController)
	App.Register(EventController)

	fmt.Printf("[%s] ===== Server Started at 8080.\n", time.Now().Format("2006-01-02 15:04:05"))
	panic(fasthttp.ListenAndServe(":8080", App.HandleRequest))
}

```
Before/After Filter in Controller, return error in Before Handler to Abort normal handlers(e.g., Auth)
``` go
func (c *UserControllerStruct) Register(App *r.Router) {

	// BeforeFilter
	App.Before("/users", func(c *r.Context) error {
		fmt.Println("BeforeFilter - users")
		e := "No Permission."
		c.ResponseErrJSON(e)
		return errors.New(e)
	})

	// AfterFilter
	App.After("/users", func(c *r.Context) error {
		fmt.Println("AfterFilter - users")
		return nil
	})
...
```
Test your API
``` bash
curl -H "Content-Type: application/json" -X POST http://localhost:8080/users -d '{"name": "John"}'
curl -H "Content-Type: application/json" -X GET http://localhost:8080/users?page=1
curl -H "Content-Type: application/json" -X GET http://localhost:8080/users/1
curl -H "Content-Type: application/json" -X PUT http://localhost:8080/users/1 -d '{"name": "John"}'
curl -H "Content-Type: application/json" -X DELETE http://localhost:8080/users/1
```
BenchMark
``` bash
ab -n 100000 -c 100 -k http://127.0.0.1:8080/users
```