Go-API is a micro API service framework, auto-generate RESTful API based on model yml

# Go-API Features
- Restful JSON API generator
- Auto support for application/x-www-form-urlencoded & application/json
- NamedParams & Params(Query Args/POST Body)
- Routing & Filter & Auth
- Access Log
- MySQL ORM & Redis Support
- Auto create table
- Model & CRUD methods generator
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
go run gen.go
```
Start Server with go run server.go
```
package main

import (
    "encoding/json"
    "fmt"
    "github.com/valyala/fasthttp"
)

func main() {
    app := App()
    c := Redis("localhost:6379")
    db, _ := MySQL("test:test@/test")

    app.Use("/ping", func(ctx *fasthttp.RequestCtx, next func(error)) {
        pong, _ := c.Ping().Result()
        ctx.SetBody([]byte(pong))
    })

    app.Use("/key", func(ctx *fasthttp.RequestCtx, next func(error)) {
        c.Set("key", "123456", 0)
        val, _ := c.Get("key").Result()
        ctx.SetBody([]byte(val))
    })

    app.Use("/db", func(ctx *fasthttp.RequestCtx, next func(error)) {
        rows, err := db.Query("select * from users")
        if err != nil {
            panic(err)
        }
        users := make(map[int]string)
        for rows.Next() {
            var uid int
            var username string
            err = rows.Scan(&uid, &username)
            if err != nil {
                panic(err)
            }
            users[uid] = username
        }
        data, err := json.Marshal(users)
        if err != nil {
            panic(err)
        }
        ctx.SetBody([]byte(data))
    })

    app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
        ctx.SetBody([]byte("666"))
    })

    server := &fasthttp.Server{
        Handler:     app.Handler,
        Concurrency: 1024 * 1024,
    }

    fmt.Println("Server started at :9798")
    server.ListenAndServe(":9798")
}
```

BenchMark
```
ab -n 1000000 -c 10000 -k http://127.0.0.1:8080/
```