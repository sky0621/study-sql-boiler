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

func testRentals(t *testing.T) {
	t.Parallel()

	query := Rentals()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRentalsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
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

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRentalsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Rentals().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRentalsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RentalSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRentalsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RentalExists(ctx, tx, o.RentalID)
	if err != nil {
		t.Errorf("Unable to check if Rental exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RentalExists to return true, but got false.")
	}
}

func testRentalsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	rentalFound, err := FindRental(ctx, tx, o.RentalID)
	if err != nil {
		t.Error(err)
	}

	if rentalFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRentalsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Rentals().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRentalsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Rentals().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRentalsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	rentalOne := &Rental{}
	rentalTwo := &Rental{}
	if err = randomize.Struct(seed, rentalOne, rentalDBTypes, false, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}
	if err = randomize.Struct(seed, rentalTwo, rentalDBTypes, false, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rentalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rentalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Rentals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRentalsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	rentalOne := &Rental{}
	rentalTwo := &Rental{}
	if err = randomize.Struct(seed, rentalOne, rentalDBTypes, false, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}
	if err = randomize.Struct(seed, rentalTwo, rentalDBTypes, false, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rentalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rentalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func rentalBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func rentalAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Rental) error {
	*o = Rental{}
	return nil
}

func testRentalsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Rental{}
	o := &Rental{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, rentalDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Rental object: %s", err)
	}

	AddRentalHook(boil.BeforeInsertHook, rentalBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	rentalBeforeInsertHooks = []RentalHook{}

	AddRentalHook(boil.AfterInsertHook, rentalAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	rentalAfterInsertHooks = []RentalHook{}

	AddRentalHook(boil.AfterSelectHook, rentalAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	rentalAfterSelectHooks = []RentalHook{}

	AddRentalHook(boil.BeforeUpdateHook, rentalBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	rentalBeforeUpdateHooks = []RentalHook{}

	AddRentalHook(boil.AfterUpdateHook, rentalAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	rentalAfterUpdateHooks = []RentalHook{}

	AddRentalHook(boil.BeforeDeleteHook, rentalBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	rentalBeforeDeleteHooks = []RentalHook{}

	AddRentalHook(boil.AfterDeleteHook, rentalAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	rentalAfterDeleteHooks = []RentalHook{}

	AddRentalHook(boil.BeforeUpsertHook, rentalBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	rentalBeforeUpsertHooks = []RentalHook{}

	AddRentalHook(boil.AfterUpsertHook, rentalAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	rentalAfterUpsertHooks = []RentalHook{}
}

func testRentalsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRentalsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(rentalColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRentalToManyPayments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Rental
	var b, c Payment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, paymentDBTypes, false, paymentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, paymentDBTypes, false, paymentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RentalID = a.RentalID
	c.RentalID = a.RentalID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Payments().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.RentalID == b.RentalID {
			bFound = true
		}
		if v.RentalID == c.RentalID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RentalSlice{&a}
	if err = a.L.LoadPayments(ctx, tx, false, (*[]*Rental)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Payments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Payments = nil
	if err = a.L.LoadPayments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Payments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testRentalToManyAddOpPayments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Rental
	var b, c, d, e Payment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, rentalDBTypes, false, strmangle.SetComplement(rentalPrimaryKeyColumns, rentalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Payment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, paymentDBTypes, false, strmangle.SetComplement(paymentPrimaryKeyColumns, paymentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Payment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPayments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.RentalID != first.RentalID {
			t.Error("foreign key was wrong value", a.RentalID, first.RentalID)
		}
		if a.RentalID != second.RentalID {
			t.Error("foreign key was wrong value", a.RentalID, second.RentalID)
		}

		if first.R.Rental != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Rental != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Payments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Payments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Payments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testRentalToOneCustomerUsingCustomer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Rental
	var foreign Customer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, rentalDBTypes, false, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, customerDBTypes, false, customerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Customer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CustomerID = foreign.CustomerID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Customer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.CustomerID != foreign.CustomerID {
		t.Errorf("want: %v, got %v", foreign.CustomerID, check.CustomerID)
	}

	slice := RentalSlice{&local}
	if err = local.L.LoadCustomer(ctx, tx, false, (*[]*Rental)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Customer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Customer = nil
	if err = local.L.LoadCustomer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Customer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRentalToOneInventoryUsingInventory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Rental
	var foreign Inventory

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, rentalDBTypes, false, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, inventoryDBTypes, false, inventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Inventory struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.InventoryID = foreign.InventoryID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Inventory().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.InventoryID != foreign.InventoryID {
		t.Errorf("want: %v, got %v", foreign.InventoryID, check.InventoryID)
	}

	slice := RentalSlice{&local}
	if err = local.L.LoadInventory(ctx, tx, false, (*[]*Rental)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Inventory == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Inventory = nil
	if err = local.L.LoadInventory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Inventory == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRentalToOneStaffUsingStaff(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Rental
	var foreign Staff

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, rentalDBTypes, false, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, staffDBTypes, false, staffColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Staff struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StaffID = foreign.StaffID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Staff().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.StaffID != foreign.StaffID {
		t.Errorf("want: %v, got %v", foreign.StaffID, check.StaffID)
	}

	slice := RentalSlice{&local}
	if err = local.L.LoadStaff(ctx, tx, false, (*[]*Rental)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Staff == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Staff = nil
	if err = local.L.LoadStaff(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Staff == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRentalToOneSetOpCustomerUsingCustomer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Rental
	var b, c Customer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, rentalDBTypes, false, strmangle.SetComplement(rentalPrimaryKeyColumns, rentalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, customerDBTypes, false, strmangle.SetComplement(customerPrimaryKeyColumns, customerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Customer{&b, &c} {
		err = a.SetCustomer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Customer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Rentals[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CustomerID != x.CustomerID {
			t.Error("foreign key was wrong value", a.CustomerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CustomerID))
		reflect.Indirect(reflect.ValueOf(&a.CustomerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CustomerID != x.CustomerID {
			t.Error("foreign key was wrong value", a.CustomerID, x.CustomerID)
		}
	}
}
func testRentalToOneSetOpInventoryUsingInventory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Rental
	var b, c Inventory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, rentalDBTypes, false, strmangle.SetComplement(rentalPrimaryKeyColumns, rentalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, inventoryDBTypes, false, strmangle.SetComplement(inventoryPrimaryKeyColumns, inventoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, inventoryDBTypes, false, strmangle.SetComplement(inventoryPrimaryKeyColumns, inventoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Inventory{&b, &c} {
		err = a.SetInventory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Inventory != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Rentals[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.InventoryID != x.InventoryID {
			t.Error("foreign key was wrong value", a.InventoryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.InventoryID))
		reflect.Indirect(reflect.ValueOf(&a.InventoryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.InventoryID != x.InventoryID {
			t.Error("foreign key was wrong value", a.InventoryID, x.InventoryID)
		}
	}
}
func testRentalToOneSetOpStaffUsingStaff(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Rental
	var b, c Staff

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, rentalDBTypes, false, strmangle.SetComplement(rentalPrimaryKeyColumns, rentalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, staffDBTypes, false, strmangle.SetComplement(staffPrimaryKeyColumns, staffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, staffDBTypes, false, strmangle.SetComplement(staffPrimaryKeyColumns, staffColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Staff{&b, &c} {
		err = a.SetStaff(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Staff != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Rentals[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StaffID != x.StaffID {
			t.Error("foreign key was wrong value", a.StaffID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StaffID))
		reflect.Indirect(reflect.ValueOf(&a.StaffID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StaffID != x.StaffID {
			t.Error("foreign key was wrong value", a.StaffID, x.StaffID)
		}
	}
}

func testRentalsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
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

func testRentalsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RentalSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRentalsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Rentals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	rentalDBTypes = map[string]string{`RentalID`: `integer`, `RentalDate`: `timestamp without time zone`, `InventoryID`: `integer`, `CustomerID`: `integer`, `ReturnDate`: `timestamp without time zone`, `StaffID`: `integer`, `LastUpdate`: `timestamp without time zone`}
	_             = bytes.MinRead
)

func testRentalsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(rentalPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(rentalAllColumns) == len(rentalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRentalsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(rentalAllColumns) == len(rentalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Rental{}
	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rentalDBTypes, true, rentalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(rentalAllColumns, rentalPrimaryKeyColumns) {
		fields = rentalAllColumns
	} else {
		fields = strmangle.SetComplement(
			rentalAllColumns,
			rentalPrimaryKeyColumns,
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

	slice := RentalSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRentalsUpsert(t *testing.T) {
	t.Parallel()

	if len(rentalAllColumns) == len(rentalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Rental{}
	if err = randomize.Struct(seed, &o, rentalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Rental: %s", err)
	}

	count, err := Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, rentalDBTypes, false, rentalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Rental struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Rental: %s", err)
	}

	count, err = Rentals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
