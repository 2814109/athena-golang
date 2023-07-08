// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Article is an object representing the database table.
type Article struct {
	ID          int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description string `boil:"description" json:"description" toml:"description" yaml:"description"`
	Status      string `boil:"status" json:"status" toml:"status" yaml:"status"`

	R *articleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L articleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArticleColumns = struct {
	ID          string
	Title       string
	Description string
	Status      string
}{
	ID:          "id",
	Title:       "title",
	Description: "description",
	Status:      "status",
}

var ArticleTableColumns = struct {
	ID          string
	Title       string
	Description string
	Status      string
}{
	ID:          "articles.id",
	Title:       "articles.title",
	Description: "articles.description",
	Status:      "articles.status",
}

// Generated where

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

var ArticleWhere = struct {
	ID          whereHelperint
	Title       whereHelperstring
	Description whereHelperstring
	Status      whereHelperstring
}{
	ID:          whereHelperint{field: "\"articles\".\"id\""},
	Title:       whereHelperstring{field: "\"articles\".\"title\""},
	Description: whereHelperstring{field: "\"articles\".\"description\""},
	Status:      whereHelperstring{field: "\"articles\".\"status\""},
}

// ArticleRels is where relationship names are stored.
var ArticleRels = struct {
	ArticleStatus string
}{
	ArticleStatus: "ArticleStatus",
}

// articleR is where relationships are stored.
type articleR struct {
	ArticleStatus *Status `boil:"ArticleStatus" json:"ArticleStatus" toml:"ArticleStatus" yaml:"ArticleStatus"`
}

// NewStruct creates a new relationship struct
func (*articleR) NewStruct() *articleR {
	return &articleR{}
}

func (r *articleR) GetArticleStatus() *Status {
	if r == nil {
		return nil
	}
	return r.ArticleStatus
}

// articleL is where Load methods for each relationship are stored.
type articleL struct{}

var (
	articleAllColumns            = []string{"id", "title", "description", "status"}
	articleColumnsWithoutDefault = []string{"title", "description", "status"}
	articleColumnsWithDefault    = []string{"id"}
	articlePrimaryKeyColumns     = []string{"id"}
	articleGeneratedColumns      = []string{}
)

type (
	// ArticleSlice is an alias for a slice of pointers to Article.
	// This should almost always be used instead of []Article.
	ArticleSlice []*Article
	// ArticleHook is the signature for custom Article hook methods
	ArticleHook func(context.Context, boil.ContextExecutor, *Article) error

	articleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	articleType                 = reflect.TypeOf(&Article{})
	articleMapping              = queries.MakeStructMapping(articleType)
	articlePrimaryKeyMapping, _ = queries.BindMapping(articleType, articleMapping, articlePrimaryKeyColumns)
	articleInsertCacheMut       sync.RWMutex
	articleInsertCache          = make(map[string]insertCache)
	articleUpdateCacheMut       sync.RWMutex
	articleUpdateCache          = make(map[string]updateCache)
	articleUpsertCacheMut       sync.RWMutex
	articleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var articleAfterSelectHooks []ArticleHook

var articleBeforeInsertHooks []ArticleHook
var articleAfterInsertHooks []ArticleHook

var articleBeforeUpdateHooks []ArticleHook
var articleAfterUpdateHooks []ArticleHook

var articleBeforeDeleteHooks []ArticleHook
var articleAfterDeleteHooks []ArticleHook

var articleBeforeUpsertHooks []ArticleHook
var articleAfterUpsertHooks []ArticleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Article) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Article) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Article) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Article) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Article) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Article) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Article) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Article) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Article) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArticleHook registers your hook function for all future operations.
func AddArticleHook(hookPoint boil.HookPoint, articleHook ArticleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		articleAfterSelectHooks = append(articleAfterSelectHooks, articleHook)
	case boil.BeforeInsertHook:
		articleBeforeInsertHooks = append(articleBeforeInsertHooks, articleHook)
	case boil.AfterInsertHook:
		articleAfterInsertHooks = append(articleAfterInsertHooks, articleHook)
	case boil.BeforeUpdateHook:
		articleBeforeUpdateHooks = append(articleBeforeUpdateHooks, articleHook)
	case boil.AfterUpdateHook:
		articleAfterUpdateHooks = append(articleAfterUpdateHooks, articleHook)
	case boil.BeforeDeleteHook:
		articleBeforeDeleteHooks = append(articleBeforeDeleteHooks, articleHook)
	case boil.AfterDeleteHook:
		articleAfterDeleteHooks = append(articleAfterDeleteHooks, articleHook)
	case boil.BeforeUpsertHook:
		articleBeforeUpsertHooks = append(articleBeforeUpsertHooks, articleHook)
	case boil.AfterUpsertHook:
		articleAfterUpsertHooks = append(articleAfterUpsertHooks, articleHook)
	}
}

// One returns a single article record from the query.
func (q articleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Article, error) {
	o := &Article{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for articles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Article records from the query.
func (q articleQuery) All(ctx context.Context, exec boil.ContextExecutor) (ArticleSlice, error) {
	var o []*Article

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Article slice")
	}

	if len(articleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Article records in the query.
func (q articleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count articles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q articleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if articles exists")
	}

	return count > 0, nil
}

// ArticleStatus pointed to by the foreign key.
func (o *Article) ArticleStatus(mods ...qm.QueryMod) statusQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"label\" = ?", o.Status),
	}

	queryMods = append(queryMods, mods...)

	return Statuses(queryMods...)
}

// LoadArticleStatus allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (articleL) LoadArticleStatus(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArticle interface{}, mods queries.Applicator) error {
	var slice []*Article
	var object *Article

	if singular {
		var ok bool
		object, ok = maybeArticle.(*Article)
		if !ok {
			object = new(Article)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeArticle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeArticle))
			}
		}
	} else {
		s, ok := maybeArticle.(*[]*Article)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeArticle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeArticle))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &articleR{}
		}
		args = append(args, object.Status)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &articleR{}
			}

			for _, a := range args {
				if a == obj.Status {
					continue Outer
				}
			}

			args = append(args, obj.Status)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`statuses`),
		qm.WhereIn(`statuses.label in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Status")
	}

	var resultSlice []*Status
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Status")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for statuses")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for statuses")
	}

	if len(statusAfterSelectHooks) != 0 {
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
		object.R.ArticleStatus = foreign
		if foreign.R == nil {
			foreign.R = &statusR{}
		}
		foreign.R.Articles = append(foreign.R.Articles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Status == foreign.Label {
				local.R.ArticleStatus = foreign
				if foreign.R == nil {
					foreign.R = &statusR{}
				}
				foreign.R.Articles = append(foreign.R.Articles, local)
				break
			}
		}
	}

	return nil
}

// SetArticleStatus of the article to the related item.
// Sets o.R.ArticleStatus to related.
// Adds o to related.R.Articles.
func (o *Article) SetArticleStatus(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Status) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"articles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"status"}),
		strmangle.WhereClause("\"", "\"", 2, articlePrimaryKeyColumns),
	)
	values := []interface{}{related.Label, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Status = related.Label
	if o.R == nil {
		o.R = &articleR{
			ArticleStatus: related,
		}
	} else {
		o.R.ArticleStatus = related
	}

	if related.R == nil {
		related.R = &statusR{
			Articles: ArticleSlice{o},
		}
	} else {
		related.R.Articles = append(related.R.Articles, o)
	}

	return nil
}

// Articles retrieves all the records using an executor.
func Articles(mods ...qm.QueryMod) articleQuery {
	mods = append(mods, qm.From("\"articles\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"articles\".*"})
	}

	return articleQuery{q}
}

// FindArticle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArticle(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Article, error) {
	articleObj := &Article{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"articles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, articleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from articles")
	}

	if err = articleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return articleObj, err
	}

	return articleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Article) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no articles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(articleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	articleInsertCacheMut.RLock()
	cache, cached := articleInsertCache[key]
	articleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			articleAllColumns,
			articleColumnsWithDefault,
			articleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(articleType, articleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(articleType, articleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"articles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"articles\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into articles")
	}

	if !cached {
		articleInsertCacheMut.Lock()
		articleInsertCache[key] = cache
		articleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Article.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Article) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	articleUpdateCacheMut.RLock()
	cache, cached := articleUpdateCache[key]
	articleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			articleAllColumns,
			articlePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update articles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"articles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, articlePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(articleType, articleMapping, append(wl, articlePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update articles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for articles")
	}

	if !cached {
		articleUpdateCacheMut.Lock()
		articleUpdateCache[key] = cache
		articleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q articleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for articles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArticleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), articlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"articles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, articlePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in article slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all article")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Article) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no articles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(articleColumnsWithDefault, o)

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

	articleUpsertCacheMut.RLock()
	cache, cached := articleUpsertCache[key]
	articleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			articleAllColumns,
			articleColumnsWithDefault,
			articleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			articleAllColumns,
			articlePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert articles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(articlePrimaryKeyColumns))
			copy(conflict, articlePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"articles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(articleType, articleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(articleType, articleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert articles")
	}

	if !cached {
		articleUpsertCacheMut.Lock()
		articleUpsertCache[key] = cache
		articleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Article record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Article) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Article provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), articlePrimaryKeyMapping)
	sql := "DELETE FROM \"articles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for articles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q articleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no articleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from articles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for articles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArticleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(articleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), articlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"articles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, articlePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from article slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for articles")
	}

	if len(articleAfterDeleteHooks) != 0 {
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
func (o *Article) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindArticle(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArticleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ArticleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), articlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"articles\".* FROM \"articles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, articlePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ArticleSlice")
	}

	*o = slice

	return nil
}

// ArticleExists checks if the Article row exists.
func ArticleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"articles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if articles exists")
	}

	return exists, nil
}

// Exists checks if the Article row exists.
func (o *Article) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ArticleExists(ctx, exec, o.ID)
}
