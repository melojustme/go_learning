package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//链接数据库
	// connStr := "user=melo dbname=test sslmode=verify-full"
	connStr := "postgres://postgres:5763@192.168.99.100:32769/test?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//查询
	SelectBird(db)
	//
}

type Birds struct {
	Id          int
	Bird        string
	Description string
}

//查询
func SelectBird(db *sql.DB) {
	b := Birds{}
	id := 2
	rows, err := db.Query(`SELECT * FROM "public"."birds" WHERE id = $1`, id)
	fmt.Println(rows, err)
	for rows.Next() {
		err := rows.Scan(&b.Id, &b.Bird, &b.Description)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(b)
}
