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

	. "github.com/sky0621/study-sql-boiler/models"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=dvdrental user=postgres password=localpass sslmode=disable port=11311")
	if err != nil {
		panic(err)
	}

	{
		l, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		boil.SetLocation(l)
	}

	ctx := context.Background()

	//var mods []QueryMod

	//mods = append(mods, Select("film_id", "title", "description"))
	//mods = append(mods, Where("film_id < ?", 10))
	//mods = append(mods, Limit(3))
	//mods = append(mods, Offset(2))
	//findFilms(ctx, db, mods)

	//{
	//	c := ActorColumns
	//	w := ActorWhere
	//
	//	mods = append(mods, Select(c.ActorID, c.FirstName))
	//	mods = append(mods, w.ActorID.GTE(3))
	//	mods = append(mods, Limit(3))
	//	findActor(ctx, db, mods)
	//}

	{
		fas, err := FilmActors().One(ctx, db)
		if err != nil {
			panic(err)
		}
		as, err := fas.Actor().All(ctx, db)
		if err != nil {
			panic(err)
		}
		for _, a := range as {
			fmt.Println(a)
		}

		//c := CountryColumns
		//
		//mods = append(mods, Select(c.CountryID))
		//res, err := Countries(mods...).One(ctx, db)
		//if err != nil {
		//	panic(err)
		//}
		//res.j
	}
}

func findFilms(ctx context.Context, db *sql.DB, mods []QueryMod) {
	records, err := Films(mods...).All(ctx, db)
	if err != nil {
		panic(err)
	}
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

func findActor(ctx context.Context, db *sql.DB, mods []QueryMod) {
	r := ActorRels

	records, err := Actors(append(mods, Load(r.FilmActors))...).All(ctx, db)
	if err != nil {
		panic(err)
	}
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
