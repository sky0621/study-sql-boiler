package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/boil"

	. "github.com/volatiletech/sqlboiler/queries/qm"

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
		// SELECT COUNT(*) FROM "pilots";
		cnt, err := Pilots().Count(ctx, db)
		if err != nil {
			panic(err)
		}
		fmt.Println(cnt)
	}

	{
		// SELECT * FROM "pilots" LIMIT 5;
		results, err := Pilots(Limit(5)).All(ctx, db)
		check(err)
		fmt.Println(results)
	}

	{
		// DELETE FROM "pilots" WHERE ("pilots"."id" = $1);
		// [1]
		cnt, err := Pilots(PilotWhere.ID.EQ(1)).DeleteAll(ctx, db)
		check(err)
		fmt.Println(cnt)
	}

	{
		var res *Pilot
		// SELECT * FROM "pilots";
		err := NewQuery(From(TableNames.Pilots)).Bind(ctx, db, res)
		check(err)
		fmt.Println(res)
	}
}

func check(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no records")
		} else {
			panic(err)
		}
	}
}
