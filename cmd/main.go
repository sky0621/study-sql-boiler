package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sky0621/study-sql-boiler/models"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=dvdrental user=postgres password=localpass sslmode=disable port=11311")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	actors, err := models.Actors().All(ctx, db)
	if err != nil {
		panic(err)
	}
	for _, actor := range actors {
		fmt.Println(actor)
	}

}
