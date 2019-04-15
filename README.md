Run
```
go run *.go
ab -n 1000000 -c 10000 -k http://127.0.0.1:9798/
```

Example
```go
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