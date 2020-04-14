package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/boil"

	. "github.com/volatiletech/sqlboiler/queries/qm"

	. "github.com/sky0621/study-sql-boiler/example/models"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	var tx *sql.Tx
	{
		var db *sql.DB
		var err error
		db, err = sql.Open("postgres", "dbname=sqlboilerexample user=postgres password=localpass sslmode=disable port=11411")
		if err != nil {
			panic(err)
		}
		defer func() {
			if db != nil {
				if err := db.Close(); err != nil {
					panic(err)
				}
			}
		}()

		tx, err = db.BeginTx(ctx, nil)
		if err != nil {
			panic(err)
		}
		defer func() {
			if tx != nil {
				if err := tx.Rollback(); err != nil {
					panic(err)
				}
			}
		}()

		boil.DebugMode = true

		var loc *time.Location
		loc, err = time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		boil.SetLocation(loc)
	}

	insertPilots(ctx, tx)
	records := findAllPilots(ctx, tx)
	updatePilots(ctx, tx, records)
	findAllPilots(ctx, tx)
}

func insertPilots(ctx context.Context, tx *sql.Tx) {
	var p Pilot
	names := map[int]string{1: "Larry", 2: "Boris", 3: "Rupert", 4: "Nigel", 0: "Bill"}
	for k, v := range names {
		p.ID = k
		p.Name = v
		if err := p.Insert(ctx, tx, boil.Infer()); err != nil {
			panic(err)
		}
	}
}

func findAllPilots(ctx context.Context, tx *sql.Tx) PilotSlice {
	records, err := Pilots().All(ctx, tx)
	check(err)
	for _, record := range records {
		ba, err := json.Marshal(record)
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		err = json.Indent(&buf, ba, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(buf.String())
	}
	return records
}

func updatePilots(ctx context.Context, tx *sql.Tx, records PilotSlice) {
	c := PilotColumns
	rowAff, err := records.UpdateAll(ctx, tx, M{
		c.Name: "zero",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rowAff)
}

/****************************************************************
 * クエリ
 ****************************************************************/
func queries(ctx context.Context, db *sql.DB) {

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
		var res *Pilot
		// SELECT * FROM "pilots";
		err := NewQuery(From(TableNames.Pilots)).Bind(ctx, db, res)
		check(err)
		fmt.Println(res)
	}

	{
		records, err := Pilots(SQL("SELECT * FROM pilots WHERE id > 3")).All(ctx, db)
		check(err)
		for _, record := range records {
			ba, err := json.Marshal(record)
			if err != nil {
				panic(err)
			}
			var buf bytes.Buffer
			err = json.Indent(&buf, ba, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(buf.String())
		}
	}

}

/****************************************************************
 * 削除
 ****************************************************************/
func del(ctx context.Context, db *sql.DB) {

	{
		// DELETE FROM "pilots" WHERE ("pilots"."id" = $1);
		// [1]
		cnt, err := Pilots(PilotWhere.ID.EQ(1)).DeleteAll(ctx, db)
		check(err)
		fmt.Println(cnt)
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
