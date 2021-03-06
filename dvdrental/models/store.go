// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Store is an object representing the database table.
type Store struct {
	StoreID        int       `boil:"store_id" json:"store_id" toml:"store_id" yaml:"store_id"`
	ManagerStaffID int       `boil:"manager_staff_id" json:"manager_staff_id" toml:"manager_staff_id" yaml:"manager_staff_id"`
	AddressID      int       `boil:"address_id" json:"address_id" toml:"address_id" yaml:"address_id"`
	LastUpdate     time.Time `boil:"last_update" json:"last_update" toml:"last_update" yaml:"last_update"`

	R *storeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L storeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StoreColumns = struct {
	StoreID        string
	ManagerStaffID string
	AddressID      string
	LastUpdate     string
}{
	StoreID:        "store_id",
	ManagerStaffID: "manager_staff_id",
	AddressID:      "address_id",
	LastUpdate:     "last_update",
}

// Generated where

var StoreWhere = struct {
	StoreID        whereHelperint
	ManagerStaffID whereHelperint
	AddressID      whereHelperint
	LastUpdate     whereHelpertime_Time
}{
	StoreID:        whereHelperint{field: "\"store\".\"store_id\""},
	ManagerStaffID: whereHelperint{field: "\"store\".\"manager_staff_id\""},
	AddressID:      whereHelperint{field: "\"store\".\"address_id\""},
	LastUpdate:     whereHelpertime_Time{field: "\"store\".\"last_update\""},
}

// StoreRels is where relationship names are stored.
var StoreRels = struct {
	Address      string
	ManagerStaff string
}{
	Address:      "Address",
	ManagerStaff: "ManagerStaff",
}

// storeR is where relationships are stored.
type storeR struct {
	Address      *Address
	ManagerStaff *Staff
}

// NewStruct creates a new relationship struct
func (*storeR) NewStruct() *storeR {
	return &storeR{}
}

// storeL is where Load methods for each relationship are stored.
type storeL struct{}

var (
	storeAllColumns            = []string{"store_id", "manager_staff_id", "address_id", "last_update"}
	storeColumnsWithoutDefault = []string{"manager_staff_id", "address_id"}
	storeColumnsWithDefault    = []string{"store_id", "last_update"}
	storePrimaryKeyColumns     = []string{"store_id"}
)

type (
	// StoreSlice is an alias for a slice of pointers to Store.
	// This should generally be used opposed to []Store.
	StoreSlice []*Store
	// StoreHook is the signature for custom Store hook methods
	StoreHook func(context.Context, boil.ContextExecutor, *Store) error

	storeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	storeType                 = reflect.TypeOf(&Store{})
	storeMapping              = queries.MakeStructMapping(storeType)
	storePrimaryKeyMapping, _ = queries.BindMapping(storeType, storeMapping, storePrimaryKeyColumns)
	storeInsertCacheMut       sync.RWMutex
	storeInsertCache          = make(map[string]insertCache)
	storeUpdateCacheMut       sync.RWMutex
	storeUpdateCache          = make(map[string]updateCache)
	storeUpsertCacheMut       sync.RWMutex
	storeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var storeBeforeInsertHooks []StoreHook
var storeBeforeUpdateHooks []StoreHook
var storeBeforeDeleteHooks []StoreHook
var storeBeforeUpsertHooks []StoreHook

var storeAfterInsertHooks []StoreHook
var storeAfterSelectHooks []StoreHook
var storeAfterUpdateHooks []StoreHook
var storeAfterDeleteHooks []StoreHook
var storeAfterUpsertHooks []StoreHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Store) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Store) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Store) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Store) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Store) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Store) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Store) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Store) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Store) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range storeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStoreHook registers your hook function for all future operations.
func AddStoreHook(hookPoint boil.HookPoint, storeHook StoreHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		storeBeforeInsertHooks = append(storeBeforeInsertHooks, storeHook)
	case boil.BeforeUpdateHook:
		storeBeforeUpdateHooks = append(storeBeforeUpdateHooks, storeHook)
	case boil.BeforeDeleteHook:
		storeBeforeDeleteHooks = append(storeBeforeDeleteHooks, storeHook)
	case boil.BeforeUpsertHook:
		storeBeforeUpsertHooks = append(storeBeforeUpsertHooks, storeHook)
	case boil.AfterInsertHook:
		storeAfterInsertHooks = append(storeAfterInsertHooks, storeHook)
	case boil.AfterSelectHook:
		storeAfterSelectHooks = append(storeAfterSelectHooks, storeHook)
	case boil.AfterUpdateHook:
		storeAfterUpdateHooks = append(storeAfterUpdateHooks, storeHook)
	case boil.AfterDeleteHook:
		storeAfterDeleteHooks = append(storeAfterDeleteHooks, storeHook)
	case boil.AfterUpsertHook:
		storeAfterUpsertHooks = append(storeAfterUpsertHooks, storeHook)
	}
}

// One returns a single store record from the query.
func (q storeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Store, error) {
	o := &Store{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for store")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Store records from the query.
func (q storeQuery) All(ctx context.Context, exec boil.ContextExecutor) (StoreSlice, error) {
	var o []*Store

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Store slice")
	}

	if len(storeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Store records in the query.
func (q storeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count store rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q storeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if store exists")
	}

	return count > 0, nil
}

// Address pointed to by the foreign key.
func (o *Store) Address(mods ...qm.QueryMod) addressQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"address_id\" = ?", o.AddressID),
	}

	queryMods = append(queryMods, mods...)

	query := Addresses(queryMods...)
	queries.SetFrom(query.Query, "\"address\"")

	return query
}

// ManagerStaff pointed to by the foreign key.
func (o *Store) ManagerStaff(mods ...qm.QueryMod) staffQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"staff_id\" = ?", o.ManagerStaffID),
	}

	queryMods = append(queryMods, mods...)

	query := Staffs(queryMods...)
	queries.SetFrom(query.Query, "\"staff\"")

	return query
}

// LoadAddress allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (storeL) LoadAddress(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStore interface{}, mods queries.Applicator) error {
	var slice []*Store
	var object *Store

	if singular {
		object = maybeStore.(*Store)
	} else {
		slice = *maybeStore.(*[]*Store)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &storeR{}
		}
		args = append(args, object.AddressID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &storeR{}
			}

			for _, a := range args {
				if a == obj.AddressID {
					continue Outer
				}
			}

			args = append(args, obj.AddressID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`address`), qm.WhereIn(`address.address_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Address")
	}

	var resultSlice []*Address
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Address")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for address")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for address")
	}

	if len(storeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Address = foreign
		if foreign.R == nil {
			foreign.R = &addressR{}
		}
		foreign.R.Stores = append(foreign.R.Stores, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AddressID == foreign.AddressID {
				local.R.Address = foreign
				if foreign.R == nil {
					foreign.R = &addressR{}
				}
				foreign.R.Stores = append(foreign.R.Stores, local)
				break
			}
		}
	}

	return nil
}

// LoadManagerStaff allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (storeL) LoadManagerStaff(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStore interface{}, mods queries.Applicator) error {
	var slice []*Store
	var object *Store

	if singular {
		object = maybeStore.(*Store)
	} else {
		slice = *maybeStore.(*[]*Store)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &storeR{}
		}
		args = append(args, object.ManagerStaffID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &storeR{}
			}

			for _, a := range args {
				if a == obj.ManagerStaffID {
					continue Outer
				}
			}

			args = append(args, obj.ManagerStaffID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`staff`), qm.WhereIn(`staff.staff_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Staff")
	}

	var resultSlice []*Staff
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Staff")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for staff")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for staff")
	}

	if len(storeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ManagerStaff = foreign
		if foreign.R == nil {
			foreign.R = &staffR{}
		}
		foreign.R.ManagerStaffStore = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ManagerStaffID == foreign.StaffID {
				local.R.ManagerStaff = foreign
				if foreign.R == nil {
					foreign.R = &staffR{}
				}
				foreign.R.ManagerStaffStore = local
				break
			}
		}
	}

	return nil
}

// SetAddress of the store to the related item.
// Sets o.R.Address to related.
// Adds o to related.R.Stores.
func (o *Store) SetAddress(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Address) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"store\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"address_id"}),
		strmangle.WhereClause("\"", "\"", 2, storePrimaryKeyColumns),
	)
	values := []interface{}{related.AddressID, o.StoreID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AddressID = related.AddressID
	if o.R == nil {
		o.R = &storeR{
			Address: related,
		}
	} else {
		o.R.Address = related
	}

	if related.R == nil {
		related.R = &addressR{
			Stores: StoreSlice{o},
		}
	} else {
		related.R.Stores = append(related.R.Stores, o)
	}

	return nil
}

// SetManagerStaff of the store to the related item.
// Sets o.R.ManagerStaff to related.
// Adds o to related.R.ManagerStaffStore.
func (o *Store) SetManagerStaff(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Staff) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"store\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"manager_staff_id"}),
		strmangle.WhereClause("\"", "\"", 2, storePrimaryKeyColumns),
	)
	values := []interface{}{related.StaffID, o.StoreID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ManagerStaffID = related.StaffID
	if o.R == nil {
		o.R = &storeR{
			ManagerStaff: related,
		}
	} else {
		o.R.ManagerStaff = related
	}

	if related.R == nil {
		related.R = &staffR{
			ManagerStaffStore: o,
		}
	} else {
		related.R.ManagerStaffStore = o
	}

	return nil
}

// Stores retrieves all the records using an executor.
func Stores(mods ...qm.QueryMod) storeQuery {
	mods = append(mods, qm.From("\"store\""))
	return storeQuery{NewQuery(mods...)}
}

// FindStore retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStore(ctx context.Context, exec boil.ContextExecutor, storeID int, selectCols ...string) (*Store, error) {
	storeObj := &Store{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"store\" where \"store_id\"=$1", sel,
	)

	q := queries.Raw(query, storeID)

	err := q.Bind(ctx, exec, storeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from store")
	}

	return storeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Store) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no store provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(storeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	storeInsertCacheMut.RLock()
	cache, cached := storeInsertCache[key]
	storeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			storeAllColumns,
			storeColumnsWithDefault,
			storeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(storeType, storeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(storeType, storeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"store\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"store\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into store")
	}

	if !cached {
		storeInsertCacheMut.Lock()
		storeInsertCache[key] = cache
		storeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Store.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Store) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	storeUpdateCacheMut.RLock()
	cache, cached := storeUpdateCache[key]
	storeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			storeAllColumns,
			storePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update store, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"store\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, storePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(storeType, storeMapping, append(wl, storePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update store row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for store")
	}

	if !cached {
		storeUpdateCacheMut.Lock()
		storeUpdateCache[key] = cache
		storeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q storeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for store")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for store")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StoreSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"store\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, storePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in store slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all store")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Store) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no store provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(storeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	storeUpsertCacheMut.RLock()
	cache, cached := storeUpsertCache[key]
	storeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			storeAllColumns,
			storeColumnsWithDefault,
			storeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			storeAllColumns,
			storePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert store, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(storePrimaryKeyColumns))
			copy(conflict, storePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"store\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(storeType, storeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(storeType, storeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert store")
	}

	if !cached {
		storeUpsertCacheMut.Lock()
		storeUpsertCache[key] = cache
		storeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Store record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Store) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Store provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), storePrimaryKeyMapping)
	sql := "DELETE FROM \"store\" WHERE \"store_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from store")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for store")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q storeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no storeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from store")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for store")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StoreSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(storeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"store\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, storePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from store slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for store")
	}

	if len(storeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Store) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStore(ctx, exec, o.StoreID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StoreSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StoreSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"store\".* FROM \"store\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, storePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StoreSlice")
	}

	*o = slice

	return nil
}

// StoreExists checks if the Store row exists.
func StoreExists(ctx context.Context, exec boil.ContextExecutor, storeID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"store\" where \"store_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, storeID)
	}
	row := exec.QueryRowContext(ctx, sql, storeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if store exists")
	}

	return exists, nil
}
