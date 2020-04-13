package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/boil"

	. "github.com/sky0621/study-sql-boiler/example/models"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=sqlboilerexample user=postgres password=localpass sslmode=disable port=11411")
	if err != nil {
		panic(err)
	}

	{
		l, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		boil.SetLocation(l)

		boil.DebugMode = true
	}

	ctx := context.Background()

	{
		cnt, err := Pilots().Count(ctx, db)
		if err != nil {
			panic(err)
		}
		fmt.Println(cnt)
	}

}
