package main

import (
	// "fmt"
	"github.com/go-redis/redis"
)

func Redis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// func main() {
// 	c := Redis("localhost:6379")
// 	pong, err := c.Ping().Result()
// 	fmt.Println(pong, err)

// 	err = c.Set("key", "123", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, _ := c.Get("key").Result()
// 	fmt.Println(val)

// 	val, _ = c.Get("key22").Result()
// 	fmt.Println(val)
// }
