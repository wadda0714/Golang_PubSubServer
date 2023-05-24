// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package userdb

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Message is an object representing the database table.
type Message struct {
	Roomid    string    `boil:"roomid" json:"roomid" toml:"roomid" yaml:"roomid"`
	Messageid string    `boil:"messageid" json:"messageid" toml:"messageid" yaml:"messageid"`
	Content   string    `boil:"content" json:"content" toml:"content" yaml:"content"`
	Userid    int       `boil:"userid" json:"userid" toml:"userid" yaml:"userid"`
	Timestamp time.Time `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`

	R *messageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L messageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MessageColumns = struct {
	Roomid    string
	Messageid string
	Content   string
	Userid    string
	Timestamp string
}{
	Roomid:    "roomid",
	Messageid: "messageid",
	Content:   "content",
	Userid:    "userid",
	Timestamp: "timestamp",
}

var MessageTableColumns = struct {
	Roomid    string
	Messageid string
	Content   string
	Userid    string
	Timestamp string
}{
	Roomid:    "messages.roomid",
	Messageid: "messages.messageid",
	Content:   "messages.content",
	Userid:    "messages.userid",
	Timestamp: "messages.timestamp",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MessageWhere = struct {
	Roomid    whereHelperstring
	Messageid whereHelperstring
	Content   whereHelperstring
	Userid    whereHelperint
	Timestamp whereHelpertime_Time
}{
	Roomid:    whereHelperstring{field: "\"messages\".\"roomid\""},
	Messageid: whereHelperstring{field: "\"messages\".\"messageid\""},
	Content:   whereHelperstring{field: "\"messages\".\"content\""},
	Userid:    whereHelperint{field: "\"messages\".\"userid\""},
	Timestamp: whereHelpertime_Time{field: "\"messages\".\"timestamp\""},
}

// MessageRels is where relationship names are stored.
var MessageRels = struct {
	UseridUser string
}{
	UseridUser: "UseridUser",
}

// messageR is where relationships are stored.
type messageR struct {
	UseridUser *User `boil:"UseridUser" json:"UseridUser" toml:"UseridUser" yaml:"UseridUser"`
}

// NewStruct creates a new relationship struct
func (*messageR) NewStruct() *messageR {
	return &messageR{}
}

func (r *messageR) GetUseridUser() *User {
	if r == nil {
		return nil
	}
	return r.UseridUser
}

// messageL is where Load methods for each relationship are stored.
type messageL struct{}

var (
	messageAllColumns            = []string{"roomid", "messageid", "content", "userid", "timestamp"}
	messageColumnsWithoutDefault = []string{"roomid", "messageid", "content", "userid", "timestamp"}
	messageColumnsWithDefault    = []string{}
	messagePrimaryKeyColumns     = []string{"messageid"}
	messageGeneratedColumns      = []string{}
)

type (
	// MessageSlice is an alias for a slice of pointers to Message.
	// This should almost always be used instead of []Message.
	MessageSlice []*Message

	messageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	messageType                 = reflect.TypeOf(&Message{})
	messageMapping              = queries.MakeStructMapping(messageType)
	messagePrimaryKeyMapping, _ = queries.BindMapping(messageType, messageMapping, messagePrimaryKeyColumns)
	messageInsertCacheMut       sync.RWMutex
	messageInsertCache          = make(map[string]insertCache)
	messageUpdateCacheMut       sync.RWMutex
	messageUpdateCache          = make(map[string]updateCache)
	messageUpsertCacheMut       sync.RWMutex
	messageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single message record from the query.
func (q messageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Message, error) {
	o := &Message{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "userdb: failed to execute a one query for messages")
	}

	return o, nil
}

// All returns all Message records from the query.
func (q messageQuery) All(ctx context.Context, exec boil.ContextExecutor) (MessageSlice, error) {
	var o []*Message

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "userdb: failed to assign all query results to Message slice")
	}

	return o, nil
}

// Count returns the count of all Message records in the query.
func (q messageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "userdb: failed to count messages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q messageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "userdb: failed to check if messages exists")
	}

	return count > 0, nil
}

// UseridUser pointed to by the foreign key.
func (o *Message) UseridUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Userid),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUseridUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (messageL) LoadUseridUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMessage interface{}, mods queries.Applicator) error {
	var slice []*Message
	var object *Message

	if singular {
		var ok bool
		object, ok = maybeMessage.(*Message)
		if !ok {
			object = new(Message)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMessage))
			}
		}
	} else {
		s, ok := maybeMessage.(*[]*Message)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMessage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMessage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &messageR{}
		}
		args = append(args, object.Userid)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &messageR{}
			}

			for _, a := range args {
				if a == obj.Userid {
					continue Outer
				}
			}

			args = append(args, obj.Userid)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.UseridUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UseridMessages = append(foreign.R.UseridMessages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Userid == foreign.ID {
				local.R.UseridUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UseridMessages = append(foreign.R.UseridMessages, local)
				break
			}
		}
	}

	return nil
}

// SetUseridUser of the message to the related item.
// Sets o.R.UseridUser to related.
// Adds o to related.R.UseridMessages.
func (o *Message) SetUseridUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"userid"}),
		strmangle.WhereClause("\"", "\"", 2, messagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Messageid}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Userid = related.ID
	if o.R == nil {
		o.R = &messageR{
			UseridUser: related,
		}
	} else {
		o.R.UseridUser = related
	}

	if related.R == nil {
		related.R = &userR{
			UseridMessages: MessageSlice{o},
		}
	} else {
		related.R.UseridMessages = append(related.R.UseridMessages, o)
	}

	return nil
}

// Messages retrieves all the records using an executor.
func Messages(mods ...qm.QueryMod) messageQuery {
	mods = append(mods, qm.From("\"messages\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"messages\".*"})
	}

	return messageQuery{q}
}

// FindMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMessage(ctx context.Context, exec boil.ContextExecutor, messageid string, selectCols ...string) (*Message, error) {
	messageObj := &Message{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"messages\" where \"messageid\"=$1", sel,
	)

	q := queries.Raw(query, messageid)

	err := q.Bind(ctx, exec, messageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "userdb: unable to select from messages")
	}

	return messageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Message) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("userdb: no messages provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(messageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	messageInsertCacheMut.RLock()
	cache, cached := messageInsertCache[key]
	messageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			messageAllColumns,
			messageColumnsWithDefault,
			messageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(messageType, messageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"messages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"messages\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "userdb: unable to insert into messages")
	}

	if !cached {
		messageInsertCacheMut.Lock()
		messageInsertCache[key] = cache
		messageInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Message.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Message) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	messageUpdateCacheMut.RLock()
	cache, cached := messageUpdateCache[key]
	messageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("userdb: unable to update messages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"messages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, messagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, append(wl, messagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "userdb: unable to update messages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "userdb: failed to get rows affected by update for messages")
	}

	if !cached {
		messageUpdateCacheMut.Lock()
		messageUpdateCache[key] = cache
		messageUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q messageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "userdb: unable to update all for messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "userdb: unable to retrieve rows affected for messages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MessageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("userdb: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, messagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "userdb: unable to update all in message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "userdb: unable to retrieve rows affected all in update all message")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Message) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("userdb: no messages provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(messageColumnsWithDefault, o)

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

	messageUpsertCacheMut.RLock()
	cache, cached := messageUpsertCache[key]
	messageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			messageAllColumns,
			messageColumnsWithDefault,
			messageColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("userdb: unable to upsert messages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(messagePrimaryKeyColumns))
			copy(conflict, messagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"messages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(messageType, messageMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "userdb: unable to upsert messages")
	}

	if !cached {
		messageUpsertCacheMut.Lock()
		messageUpsertCache[key] = cache
		messageUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Message record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Message) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("userdb: no Message provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), messagePrimaryKeyMapping)
	sql := "DELETE FROM \"messages\" WHERE \"messageid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "userdb: unable to delete from messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "userdb: failed to get rows affected by delete for messages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q messageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("userdb: no messageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "userdb: unable to delete all from messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "userdb: failed to get rows affected by deleteall for messages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MessageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "userdb: unable to delete all from message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "userdb: failed to get rows affected by deleteall for messages")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Message) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMessage(ctx, exec, o.Messageid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"messages\".* FROM \"messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "userdb: unable to reload all in MessageSlice")
	}

	*o = slice

	return nil
}

// MessageExists checks if the Message row exists.
func MessageExists(ctx context.Context, exec boil.ContextExecutor, messageid string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"messages\" where \"messageid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, messageid)
	}
	row := exec.QueryRowContext(ctx, sql, messageid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "userdb: unable to check if messages exists")
	}

	return exists, nil
}

// Exists checks if the Message row exists.
func (o *Message) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MessageExists(ctx, exec, o.Messageid)
}
