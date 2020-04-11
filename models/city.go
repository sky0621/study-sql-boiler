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

// City is an object representing the database table.
type City struct {
	CityID     int       `boil:"city_id" json:"city_id" toml:"city_id" yaml:"city_id"`
	City       string    `boil:"city" json:"city" toml:"city" yaml:"city"`
	CountryID  int       `boil:"country_id" json:"country_id" toml:"country_id" yaml:"country_id"`
	LastUpdate time.Time `boil:"last_update" json:"last_update" toml:"last_update" yaml:"last_update"`

	R *cityR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cityL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CityColumns = struct {
	CityID     string
	City       string
	CountryID  string
	LastUpdate string
}{
	CityID:     "city_id",
	City:       "city",
	CountryID:  "country_id",
	LastUpdate: "last_update",
}

// Generated where

var CityWhere = struct {
	CityID     whereHelperint
	City       whereHelperstring
	CountryID  whereHelperint
	LastUpdate whereHelpertime_Time
}{
	CityID:     whereHelperint{field: "\"city\".\"city_id\""},
	City:       whereHelperstring{field: "\"city\".\"city\""},
	CountryID:  whereHelperint{field: "\"city\".\"country_id\""},
	LastUpdate: whereHelpertime_Time{field: "\"city\".\"last_update\""},
}

// CityRels is where relationship names are stored.
var CityRels = struct {
	Country   string
	Addresses string
}{
	Country:   "Country",
	Addresses: "Addresses",
}

// cityR is where relationships are stored.
type cityR struct {
	Country   *Country
	Addresses AddressSlice
}

// NewStruct creates a new relationship struct
func (*cityR) NewStruct() *cityR {
	return &cityR{}
}

// cityL is where Load methods for each relationship are stored.
type cityL struct{}

var (
	cityAllColumns            = []string{"city_id", "city", "country_id", "last_update"}
	cityColumnsWithoutDefault = []string{"city", "country_id"}
	cityColumnsWithDefault    = []string{"city_id", "last_update"}
	cityPrimaryKeyColumns     = []string{"city_id"}
)

type (
	// CitySlice is an alias for a slice of pointers to City.
	// This should generally be used opposed to []City.
	CitySlice []*City
	// CityHook is the signature for custom City hook methods
	CityHook func(context.Context, boil.ContextExecutor, *City) error

	cityQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cityType                 = reflect.TypeOf(&City{})
	cityMapping              = queries.MakeStructMapping(cityType)
	cityPrimaryKeyMapping, _ = queries.BindMapping(cityType, cityMapping, cityPrimaryKeyColumns)
	cityInsertCacheMut       sync.RWMutex
	cityInsertCache          = make(map[string]insertCache)
	cityUpdateCacheMut       sync.RWMutex
	cityUpdateCache          = make(map[string]updateCache)
	cityUpsertCacheMut       sync.RWMutex
	cityUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cityBeforeInsertHooks []CityHook
var cityBeforeUpdateHooks []CityHook
var cityBeforeDeleteHooks []CityHook
var cityBeforeUpsertHooks []CityHook

var cityAfterInsertHooks []CityHook
var cityAfterSelectHooks []CityHook
var cityAfterUpdateHooks []CityHook
var cityAfterDeleteHooks []CityHook
var cityAfterUpsertHooks []CityHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *City) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *City) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *City) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *City) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *City) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *City) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *City) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *City) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *City) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cityAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCityHook registers your hook function for all future operations.
func AddCityHook(hookPoint boil.HookPoint, cityHook CityHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cityBeforeInsertHooks = append(cityBeforeInsertHooks, cityHook)
	case boil.BeforeUpdateHook:
		cityBeforeUpdateHooks = append(cityBeforeUpdateHooks, cityHook)
	case boil.BeforeDeleteHook:
		cityBeforeDeleteHooks = append(cityBeforeDeleteHooks, cityHook)
	case boil.BeforeUpsertHook:
		cityBeforeUpsertHooks = append(cityBeforeUpsertHooks, cityHook)
	case boil.AfterInsertHook:
		cityAfterInsertHooks = append(cityAfterInsertHooks, cityHook)
	case boil.AfterSelectHook:
		cityAfterSelectHooks = append(cityAfterSelectHooks, cityHook)
	case boil.AfterUpdateHook:
		cityAfterUpdateHooks = append(cityAfterUpdateHooks, cityHook)
	case boil.AfterDeleteHook:
		cityAfterDeleteHooks = append(cityAfterDeleteHooks, cityHook)
	case boil.AfterUpsertHook:
		cityAfterUpsertHooks = append(cityAfterUpsertHooks, cityHook)
	}
}

// One returns a single city record from the query.
func (q cityQuery) One(ctx context.Context, exec boil.ContextExecutor) (*City, error) {
	o := &City{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for city")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all City records from the query.
func (q cityQuery) All(ctx context.Context, exec boil.ContextExecutor) (CitySlice, error) {
	var o []*City

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to City slice")
	}

	if len(cityAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all City records in the query.
func (q cityQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count city rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cityQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if city exists")
	}

	return count > 0, nil
}

// Country pointed to by the foreign key.
func (o *City) Country(mods ...qm.QueryMod) countryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"country_id\" = ?", o.CountryID),
	}

	queryMods = append(queryMods, mods...)

	query := Countries(queryMods...)
	queries.SetFrom(query.Query, "\"country\"")

	return query
}

// Addresses retrieves all the address's Addresses with an executor.
func (o *City) Addresses(mods ...qm.QueryMod) addressQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"address\".\"city_id\"=?", o.CityID),
	)

	query := Addresses(queryMods...)
	queries.SetFrom(query.Query, "\"address\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"address\".*"})
	}

	return query
}

// LoadCountry allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (cityL) LoadCountry(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCity interface{}, mods queries.Applicator) error {
	var slice []*City
	var object *City

	if singular {
		object = maybeCity.(*City)
	} else {
		slice = *maybeCity.(*[]*City)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &cityR{}
		}
		args = append(args, object.CountryID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cityR{}
			}

			for _, a := range args {
				if a == obj.CountryID {
					continue Outer
				}
			}

			args = append(args, obj.CountryID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`country`), qm.WhereIn(`country.country_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Country")
	}

	var resultSlice []*Country
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Country")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for country")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for country")
	}

	if len(cityAfterSelectHooks) != 0 {
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
		object.R.Country = foreign
		if foreign.R == nil {
			foreign.R = &countryR{}
		}
		foreign.R.Cities = append(foreign.R.Cities, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CountryID == foreign.CountryID {
				local.R.Country = foreign
				if foreign.R == nil {
					foreign.R = &countryR{}
				}
				foreign.R.Cities = append(foreign.R.Cities, local)
				break
			}
		}
	}

	return nil
}

// LoadAddresses allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (cityL) LoadAddresses(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCity interface{}, mods queries.Applicator) error {
	var slice []*City
	var object *City

	if singular {
		object = maybeCity.(*City)
	} else {
		slice = *maybeCity.(*[]*City)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &cityR{}
		}
		args = append(args, object.CityID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cityR{}
			}

			for _, a := range args {
				if a == obj.CityID {
					continue Outer
				}
			}

			args = append(args, obj.CityID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`address`), qm.WhereIn(`address.city_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load address")
	}

	var resultSlice []*Address
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice address")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on address")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for address")
	}

	if len(addressAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Addresses = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &addressR{}
			}
			foreign.R.City = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CityID == foreign.CityID {
				local.R.Addresses = append(local.R.Addresses, foreign)
				if foreign.R == nil {
					foreign.R = &addressR{}
				}
				foreign.R.City = local
				break
			}
		}
	}

	return nil
}

// SetCountry of the city to the related item.
// Sets o.R.Country to related.
// Adds o to related.R.Cities.
func (o *City) SetCountry(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Country) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"city\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"country_id"}),
		strmangle.WhereClause("\"", "\"", 2, cityPrimaryKeyColumns),
	)
	values := []interface{}{related.CountryID, o.CityID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CountryID = related.CountryID
	if o.R == nil {
		o.R = &cityR{
			Country: related,
		}
	} else {
		o.R.Country = related
	}

	if related.R == nil {
		related.R = &countryR{
			Cities: CitySlice{o},
		}
	} else {
		related.R.Cities = append(related.R.Cities, o)
	}

	return nil
}

// AddAddresses adds the given related objects to the existing relationships
// of the city, optionally inserting them as new records.
// Appends related to o.R.Addresses.
// Sets related.R.City appropriately.
func (o *City) AddAddresses(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Address) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CityID = o.CityID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"address\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"city_id"}),
				strmangle.WhereClause("\"", "\"", 2, addressPrimaryKeyColumns),
			)
			values := []interface{}{o.CityID, rel.AddressID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CityID = o.CityID
		}
	}

	if o.R == nil {
		o.R = &cityR{
			Addresses: related,
		}
	} else {
		o.R.Addresses = append(o.R.Addresses, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &addressR{
				City: o,
			}
		} else {
			rel.R.City = o
		}
	}
	return nil
}

// Cities retrieves all the records using an executor.
func Cities(mods ...qm.QueryMod) cityQuery {
	mods = append(mods, qm.From("\"city\""))
	return cityQuery{NewQuery(mods...)}
}

// FindCity retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCity(ctx context.Context, exec boil.ContextExecutor, cityID int, selectCols ...string) (*City, error) {
	cityObj := &City{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"city\" where \"city_id\"=$1", sel,
	)

	q := queries.Raw(query, cityID)

	err := q.Bind(ctx, exec, cityObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from city")
	}

	return cityObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *City) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no city provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cityColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cityInsertCacheMut.RLock()
	cache, cached := cityInsertCache[key]
	cityInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cityAllColumns,
			cityColumnsWithDefault,
			cityColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cityType, cityMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cityType, cityMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"city\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"city\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into city")
	}

	if !cached {
		cityInsertCacheMut.Lock()
		cityInsertCache[key] = cache
		cityInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the City.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *City) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cityUpdateCacheMut.RLock()
	cache, cached := cityUpdateCache[key]
	cityUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cityAllColumns,
			cityPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update city, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"city\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cityPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cityType, cityMapping, append(wl, cityPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update city row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for city")
	}

	if !cached {
		cityUpdateCacheMut.Lock()
		cityUpdateCache[key] = cache
		cityUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cityQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for city")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for city")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CitySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"city\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, cityPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in city slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all city")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *City) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no city provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cityColumnsWithDefault, o)

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

	cityUpsertCacheMut.RLock()
	cache, cached := cityUpsertCache[key]
	cityUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cityAllColumns,
			cityColumnsWithDefault,
			cityColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cityAllColumns,
			cityPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert city, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cityPrimaryKeyColumns))
			copy(conflict, cityPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"city\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(cityType, cityMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cityType, cityMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert city")
	}

	if !cached {
		cityUpsertCacheMut.Lock()
		cityUpsertCache[key] = cache
		cityUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single City record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *City) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no City provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cityPrimaryKeyMapping)
	sql := "DELETE FROM \"city\" WHERE \"city_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from city")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for city")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cityQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cityQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from city")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for city")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CitySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cityBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"city\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cityPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from city slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for city")
	}

	if len(cityAfterDeleteHooks) != 0 {
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
func (o *City) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCity(ctx, exec, o.CityID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CitySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CitySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"city\".* FROM \"city\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cityPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CitySlice")
	}

	*o = slice

	return nil
}

// CityExists checks if the City row exists.
func CityExists(ctx context.Context, exec boil.ContextExecutor, cityID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"city\" where \"city_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, cityID)
	}
	row := exec.QueryRowContext(ctx, sql, cityID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if city exists")
	}

	return exists, nil
}
