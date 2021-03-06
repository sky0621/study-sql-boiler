// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEventOnes(t *testing.T) {
	t.Parallel()

	query := EventOnes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEventOnesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventOnesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EventOnes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventOnesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventOneSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventOnesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EventOneExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EventOne exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EventOneExists to return true, but got false.")
	}
}

func testEventOnesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	eventOneFound, err := FindEventOne(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if eventOneFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEventOnesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EventOnes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEventOnesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EventOnes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEventOnesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	eventOneOne := &EventOne{}
	eventOneTwo := &EventOne{}
	if err = randomize.Struct(seed, eventOneOne, eventOneDBTypes, false, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}
	if err = randomize.Struct(seed, eventOneTwo, eventOneDBTypes, false, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventOneOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventOneTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EventOnes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEventOnesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	eventOneOne := &EventOne{}
	eventOneTwo := &EventOne{}
	if err = randomize.Struct(seed, eventOneOne, eventOneDBTypes, false, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}
	if err = randomize.Struct(seed, eventOneTwo, eventOneDBTypes, false, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventOneOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventOneTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func eventOneBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func eventOneAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EventOne) error {
	*o = EventOne{}
	return nil
}

func testEventOnesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EventOne{}
	o := &EventOne{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, eventOneDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EventOne object: %s", err)
	}

	AddEventOneHook(boil.BeforeInsertHook, eventOneBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	eventOneBeforeInsertHooks = []EventOneHook{}

	AddEventOneHook(boil.AfterInsertHook, eventOneAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	eventOneAfterInsertHooks = []EventOneHook{}

	AddEventOneHook(boil.AfterSelectHook, eventOneAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	eventOneAfterSelectHooks = []EventOneHook{}

	AddEventOneHook(boil.BeforeUpdateHook, eventOneBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	eventOneBeforeUpdateHooks = []EventOneHook{}

	AddEventOneHook(boil.AfterUpdateHook, eventOneAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	eventOneAfterUpdateHooks = []EventOneHook{}

	AddEventOneHook(boil.BeforeDeleteHook, eventOneBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	eventOneBeforeDeleteHooks = []EventOneHook{}

	AddEventOneHook(boil.AfterDeleteHook, eventOneAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	eventOneAfterDeleteHooks = []EventOneHook{}

	AddEventOneHook(boil.BeforeUpsertHook, eventOneBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	eventOneBeforeUpsertHooks = []EventOneHook{}

	AddEventOneHook(boil.AfterUpsertHook, eventOneAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	eventOneAfterUpsertHooks = []EventOneHook{}
}

func testEventOnesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventOnesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(eventOneColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventOnesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEventOnesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventOneSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEventOnesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EventOnes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	eventOneDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `Day`: `enum.workday('monday','tuesday','wednesday','thursday','friday')`}
	_               = bytes.MinRead
)

func testEventOnesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(eventOnePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(eventOneAllColumns) == len(eventOnePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOnePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEventOnesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(eventOneAllColumns) == len(eventOnePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EventOne{}
	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventOneDBTypes, true, eventOnePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(eventOneAllColumns, eventOnePrimaryKeyColumns) {
		fields = eventOneAllColumns
	} else {
		fields = strmangle.SetComplement(
			eventOneAllColumns,
			eventOnePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := EventOneSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEventOnesUpsert(t *testing.T) {
	t.Parallel()

	if len(eventOneAllColumns) == len(eventOnePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EventOne{}
	if err = randomize.Struct(seed, &o, eventOneDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EventOne: %s", err)
	}

	count, err := EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, eventOneDBTypes, false, eventOnePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventOne struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EventOne: %s", err)
	}

	count, err = EventOnes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
