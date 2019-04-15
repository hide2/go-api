package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func MySQL(dns string) (*sql.DB, error) {
	return sql.Open("mysql", dns)
}

// func main() {
// 	db, err := MySQL("test:test@/test")
// 	if err != nil {
// 		panic(err)
// 	}

// 	rows, err := db.Query("select * from users")
// 	if err != nil {
// 		panic(err)
// 	}
// 	users := make(map[int]string)
// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		err = rows.Scan(&uid, &username)
// 		if err != nil {
// 			panic(err)
// 		}
// 		users[uid] = username
// 	}
// 	fmt.Println(users)
// }
