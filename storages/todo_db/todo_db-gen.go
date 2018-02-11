package todo_db

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/sql/wrap"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"strings"
	"time"
)

var _ = sql.ErrNoRows
var _ = mysql.ErrOldProtocol

type BaseQuery struct {
	forUpdate     bool
	forShare      bool
	where         string
	limit         string
	order         string
	groupByFields []string
}

func (q *BaseQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.where != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(q.where)
	}

	if q.groupByFields != nil && len(q.groupByFields) > 0 {
		buf.WriteString(" GROUP BY ")
		buf.WriteString(strings.Join(q.groupByFields, ","))
	}

	if q.order != "" {
		buf.WriteString(" order by ")
		buf.WriteString(q.order)
	}

	if q.limit != "" {
		buf.WriteString(q.limit)
	}

	if q.forUpdate {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forShare {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	return buf.String()
}

const TODO_TABLE_NAME = "todo"

type TODO_FIELD string

const TODO_FIELD_ID = TODO_FIELD("id")
const TODO_FIELD_CREATE_TIME = TODO_FIELD("create_time")
const TODO_FIELD_UPDATE_TIME = TODO_FIELD("update_time")
const TODO_FIELD_UPDATE_VERSION = TODO_FIELD("update_version")
const TODO_FIELD_TODO_ID = TODO_FIELD("todo_id")
const TODO_FIELD_USER_ID = TODO_FIELD("user_id")
const TODO_FIELD_TODO_CATEGORY = TODO_FIELD("todo_category")
const TODO_FIELD_TODO_TITLE = TODO_FIELD("todo_title")
const TODO_FIELD_TODO_DESC = TODO_FIELD("todo_desc")
const TODO_FIELD_TODO_STATUS = TODO_FIELD("todo_status")
const TODO_FIELD_TODO_PRIORITY = TODO_FIELD("todo_priority")

const TODO_ALL_FIELDS_STRING = "id,create_time,update_time,update_version,todo_id,user_id,todo_category,todo_title,todo_desc,todo_status,todo_priority"

var TODO_ALL_FIELDS = []string{
	"id",
	"create_time",
	"update_time",
	"update_version",
	"todo_id",
	"user_id",
	"todo_category",
	"todo_title",
	"todo_desc",
	"todo_status",
	"todo_priority",
}

type Todo struct {
	Id            int64 //size=20
	CreateTime    time.Time
	UpdateTime    time.Time
	UpdateVersion int64  //size=20
	TodoId        string //size=128
	UserId        string //size=128
	TodoCategory  string //size=32
	TodoTitle     string //size=32
	TodoDesc      string //size=1024
	TodoStatus    string //size=32
	TodoPriority  int32  //size=10
}

type TodoQuery struct {
	BaseQuery
	dao *TodoDao
}

func NewTodoQuery(dao *TodoDao) *TodoQuery {
	q := &TodoQuery{}
	q.dao = dao

	return q
}

func (q *TodoQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*Todo, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *TodoQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*Todo, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *TodoQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *TodoQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *TodoQuery) ForUpdate() *TodoQuery {
	q.forUpdate = true
	return q
}

func (q *TodoQuery) ForShare() *TodoQuery {
	q.forShare = true
	return q
}

func (q *TodoQuery) GroupBy(fields ...TODO_FIELD) *TodoQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *TodoQuery) Limit(startIncluded int64, count int64) *TodoQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *TodoQuery) OrderBy(fieldName TODO_FIELD, asc bool) *TodoQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *TodoQuery) OrderByGroupCount(asc bool) *TodoQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *TodoQuery) w(format string, a ...interface{}) *TodoQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *TodoQuery) Left() *TodoQuery  { return q.w(" ( ") }
func (q *TodoQuery) Right() *TodoQuery { return q.w(" ) ") }
func (q *TodoQuery) And() *TodoQuery   { return q.w(" AND ") }
func (q *TodoQuery) Or() *TodoQuery    { return q.w(" OR ") }
func (q *TodoQuery) Not() *TodoQuery   { return q.w(" NOT ") }

func (q *TodoQuery) Id_Equal(v int64) *TodoQuery        { return q.w("id='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_NotEqual(v int64) *TodoQuery     { return q.w("id<>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_Less(v int64) *TodoQuery         { return q.w("id<'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_LessEqual(v int64) *TodoQuery    { return q.w("id<='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_Greater(v int64) *TodoQuery      { return q.w("id>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_GreaterEqual(v int64) *TodoQuery { return q.w("id>='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) CreateTime_Equal(v time.Time) *TodoQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) CreateTime_NotEqual(v time.Time) *TodoQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) CreateTime_Less(v time.Time) *TodoQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) CreateTime_LessEqual(v time.Time) *TodoQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) CreateTime_Greater(v time.Time) *TodoQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) CreateTime_GreaterEqual(v time.Time) *TodoQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateTime_Equal(v time.Time) *TodoQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateTime_NotEqual(v time.Time) *TodoQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateTime_Less(v time.Time) *TodoQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateTime_LessEqual(v time.Time) *TodoQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateTime_Greater(v time.Time) *TodoQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateTime_GreaterEqual(v time.Time) *TodoQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateVersion_Equal(v int64) *TodoQuery {
	return q.w("update_version='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateVersion_NotEqual(v int64) *TodoQuery {
	return q.w("update_version<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateVersion_Less(v int64) *TodoQuery {
	return q.w("update_version<'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateVersion_LessEqual(v int64) *TodoQuery {
	return q.w("update_version<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateVersion_Greater(v int64) *TodoQuery {
	return q.w("update_version>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UpdateVersion_GreaterEqual(v int64) *TodoQuery {
	return q.w("update_version>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoId_Equal(v string) *TodoQuery    { return q.w("todo_id='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoId_NotEqual(v string) *TodoQuery { return q.w("todo_id<>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoId_Less(v string) *TodoQuery     { return q.w("todo_id<'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoId_LessEqual(v string) *TodoQuery {
	return q.w("todo_id<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoId_Greater(v string) *TodoQuery { return q.w("todo_id>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoId_GreaterEqual(v string) *TodoQuery {
	return q.w("todo_id>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UserId_Equal(v string) *TodoQuery    { return q.w("user_id='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) UserId_NotEqual(v string) *TodoQuery { return q.w("user_id<>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) UserId_Less(v string) *TodoQuery     { return q.w("user_id<'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) UserId_LessEqual(v string) *TodoQuery {
	return q.w("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) UserId_Greater(v string) *TodoQuery { return q.w("user_id>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) UserId_GreaterEqual(v string) *TodoQuery {
	return q.w("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoCategory_Equal(v string) *TodoQuery {
	return q.w("todo_category='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoCategory_NotEqual(v string) *TodoQuery {
	return q.w("todo_category<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoCategory_Less(v string) *TodoQuery {
	return q.w("todo_category<'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoCategory_LessEqual(v string) *TodoQuery {
	return q.w("todo_category<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoCategory_Greater(v string) *TodoQuery {
	return q.w("todo_category>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoCategory_GreaterEqual(v string) *TodoQuery {
	return q.w("todo_category>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_Equal(v string) *TodoQuery {
	return q.w("todo_title='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_NotEqual(v string) *TodoQuery {
	return q.w("todo_title<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_Less(v string) *TodoQuery {
	return q.w("todo_title<'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_LessEqual(v string) *TodoQuery {
	return q.w("todo_title<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_Greater(v string) *TodoQuery {
	return q.w("todo_title>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_GreaterEqual(v string) *TodoQuery {
	return q.w("todo_title>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoDesc_Equal(v string) *TodoQuery { return q.w("todo_desc='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoDesc_NotEqual(v string) *TodoQuery {
	return q.w("todo_desc<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoDesc_Less(v string) *TodoQuery { return q.w("todo_desc<'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoDesc_LessEqual(v string) *TodoQuery {
	return q.w("todo_desc<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoDesc_Greater(v string) *TodoQuery {
	return q.w("todo_desc>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoDesc_GreaterEqual(v string) *TodoQuery {
	return q.w("todo_desc>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_Equal(v string) *TodoQuery {
	return q.w("todo_status='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_NotEqual(v string) *TodoQuery {
	return q.w("todo_status<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_Less(v string) *TodoQuery {
	return q.w("todo_status<'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_LessEqual(v string) *TodoQuery {
	return q.w("todo_status<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_Greater(v string) *TodoQuery {
	return q.w("todo_status>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_GreaterEqual(v string) *TodoQuery {
	return q.w("todo_status>='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoPriority_Equal(v int32) *TodoQuery {
	return q.w("todo_priority='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoPriority_NotEqual(v int32) *TodoQuery {
	return q.w("todo_priority<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoPriority_Less(v int32) *TodoQuery {
	return q.w("todo_priority<'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoPriority_LessEqual(v int32) *TodoQuery {
	return q.w("todo_priority<='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoPriority_Greater(v int32) *TodoQuery {
	return q.w("todo_priority>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoPriority_GreaterEqual(v int32) *TodoQuery {
	return q.w("todo_priority>='" + fmt.Sprint(v) + "'")
}

type TodoDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewTodoDao(db *DB) (t *TodoDao, err error) {
	t = &TodoDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *TodoDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}

func (dao *TodoDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO todo (update_version,todo_id,user_id,todo_category,todo_title,todo_desc,todo_status,todo_priority) VALUES (?,?,?,?,?,?,?,?)")
	return err
}

func (dao *TodoDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE todo SET update_version=update_version+1,todo_id=?,user_id=?,todo_category=?,todo_title=?,todo_desc=?,todo_status=?,todo_priority=? WHERE id=? AND update_version=?")
	return err
}

func (dao *TodoDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM todo WHERE id=?")
	return err
}

func (dao *TodoDao) Insert(ctx context.Context, tx *wrap.Tx, e *Todo) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UpdateVersion, e.TodoId, e.UserId, e.TodoCategory, e.TodoTitle, e.TodoDesc, e.TodoStatus, e.TodoPriority)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *TodoDao) Update(ctx context.Context, tx *wrap.Tx, e *Todo) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, e.TodoId, e.UserId, e.TodoCategory, e.TodoTitle, e.TodoDesc, e.TodoStatus, e.TodoPriority, e.Id, e.UpdateVersion)
	if err != nil {
		return err
	}

	return nil
}

func (dao *TodoDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *TodoDao) scanRow(row *wrap.Row) (*Todo, error) {
	e := &Todo{}
	err := row.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.TodoId, &e.UserId, &e.TodoCategory, &e.TodoTitle, &e.TodoDesc, &e.TodoStatus, &e.TodoPriority)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *TodoDao) scanRows(rows *wrap.Rows) (list []*Todo, err error) {
	list = make([]*Todo, 0)
	for rows.Next() {
		e := Todo{}
		err = rows.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.TodoId, &e.UserId, &e.TodoCategory, &e.TodoTitle, &e.TodoDesc, &e.TodoStatus, &e.TodoPriority)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *TodoDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*Todo, error) {
	querySql := "SELECT " + TODO_ALL_FIELDS_STRING + " FROM todo " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *TodoDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*Todo, err error) {
	querySql := "SELECT " + TODO_ALL_FIELDS_STRING + " FROM todo " + query
	var rows *wrap.Rows
	if tx == nil {
		rows, err = dao.db.Query(ctx, querySql)
	} else {
		rows, err = tx.Query(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *TodoDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM todo " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *TodoDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM todo " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *TodoDao) GetQuery() *TodoQuery {
	return NewTodoQuery(dao)
}

const USER_PROFILE_TABLE_NAME = "user_profile"

type USER_PROFILE_FIELD string

const USER_PROFILE_FIELD_ID = USER_PROFILE_FIELD("id")
const USER_PROFILE_FIELD_CREATE_TIME = USER_PROFILE_FIELD("create_time")
const USER_PROFILE_FIELD_UPDATE_TIME = USER_PROFILE_FIELD("update_time")
const USER_PROFILE_FIELD_UPDATE_VERSION = USER_PROFILE_FIELD("update_version")
const USER_PROFILE_FIELD_USER_ID = USER_PROFILE_FIELD("user_id")
const USER_PROFILE_FIELD_USER_NAME = USER_PROFILE_FIELD("user_name")
const USER_PROFILE_FIELD_TODO_PUBLIC_VISIBLE = USER_PROFILE_FIELD("todo_public_visible")

const USER_PROFILE_ALL_FIELDS_STRING = "id,create_time,update_time,update_version,user_id,user_name,todo_public_visible"

var USER_PROFILE_ALL_FIELDS = []string{
	"id",
	"create_time",
	"update_time",
	"update_version",
	"user_id",
	"user_name",
	"todo_public_visible",
}

type UserProfile struct {
	Id                int64 //size=20
	CreateTime        time.Time
	UpdateTime        time.Time
	UpdateVersion     int64  //size=20
	UserId            string //size=128
	UserName          string //size=128
	TodoPublicVisible int32  //size=11
}

type UserProfileQuery struct {
	BaseQuery
	dao *UserProfileDao
}

func NewUserProfileQuery(dao *UserProfileDao) *UserProfileQuery {
	q := &UserProfileQuery{}
	q.dao = dao

	return q
}

func (q *UserProfileQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*UserProfile, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *UserProfileQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*UserProfile, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *UserProfileQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *UserProfileQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *UserProfileQuery) ForUpdate() *UserProfileQuery {
	q.forUpdate = true
	return q
}

func (q *UserProfileQuery) ForShare() *UserProfileQuery {
	q.forShare = true
	return q
}

func (q *UserProfileQuery) GroupBy(fields ...USER_PROFILE_FIELD) *UserProfileQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *UserProfileQuery) Limit(startIncluded int64, count int64) *UserProfileQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *UserProfileQuery) OrderBy(fieldName USER_PROFILE_FIELD, asc bool) *UserProfileQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserProfileQuery) OrderByGroupCount(asc bool) *UserProfileQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserProfileQuery) w(format string, a ...interface{}) *UserProfileQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *UserProfileQuery) Left() *UserProfileQuery  { return q.w(" ( ") }
func (q *UserProfileQuery) Right() *UserProfileQuery { return q.w(" ) ") }
func (q *UserProfileQuery) And() *UserProfileQuery   { return q.w(" AND ") }
func (q *UserProfileQuery) Or() *UserProfileQuery    { return q.w(" OR ") }
func (q *UserProfileQuery) Not() *UserProfileQuery   { return q.w(" NOT ") }

func (q *UserProfileQuery) Id_Equal(v int64) *UserProfileQuery { return q.w("id='" + fmt.Sprint(v) + "'") }
func (q *UserProfileQuery) Id_NotEqual(v int64) *UserProfileQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) Id_Less(v int64) *UserProfileQuery { return q.w("id<'" + fmt.Sprint(v) + "'") }
func (q *UserProfileQuery) Id_LessEqual(v int64) *UserProfileQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) Id_Greater(v int64) *UserProfileQuery {
	return q.w("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) Id_GreaterEqual(v int64) *UserProfileQuery {
	return q.w("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) CreateTime_Equal(v time.Time) *UserProfileQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) CreateTime_NotEqual(v time.Time) *UserProfileQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) CreateTime_Less(v time.Time) *UserProfileQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) CreateTime_LessEqual(v time.Time) *UserProfileQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) CreateTime_Greater(v time.Time) *UserProfileQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) CreateTime_GreaterEqual(v time.Time) *UserProfileQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateTime_Equal(v time.Time) *UserProfileQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateTime_NotEqual(v time.Time) *UserProfileQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateTime_Less(v time.Time) *UserProfileQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateTime_LessEqual(v time.Time) *UserProfileQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateTime_Greater(v time.Time) *UserProfileQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateTime_GreaterEqual(v time.Time) *UserProfileQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateVersion_Equal(v int64) *UserProfileQuery {
	return q.w("update_version='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateVersion_NotEqual(v int64) *UserProfileQuery {
	return q.w("update_version<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateVersion_Less(v int64) *UserProfileQuery {
	return q.w("update_version<'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateVersion_LessEqual(v int64) *UserProfileQuery {
	return q.w("update_version<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateVersion_Greater(v int64) *UserProfileQuery {
	return q.w("update_version>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UpdateVersion_GreaterEqual(v int64) *UserProfileQuery {
	return q.w("update_version>='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserId_Equal(v string) *UserProfileQuery {
	return q.w("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserId_NotEqual(v string) *UserProfileQuery {
	return q.w("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserId_Less(v string) *UserProfileQuery {
	return q.w("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserId_LessEqual(v string) *UserProfileQuery {
	return q.w("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserId_Greater(v string) *UserProfileQuery {
	return q.w("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserId_GreaterEqual(v string) *UserProfileQuery {
	return q.w("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserName_Equal(v string) *UserProfileQuery {
	return q.w("user_name='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserName_NotEqual(v string) *UserProfileQuery {
	return q.w("user_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserName_Less(v string) *UserProfileQuery {
	return q.w("user_name<'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserName_LessEqual(v string) *UserProfileQuery {
	return q.w("user_name<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserName_Greater(v string) *UserProfileQuery {
	return q.w("user_name>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserName_GreaterEqual(v string) *UserProfileQuery {
	return q.w("user_name>='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoPublicVisible_Equal(v int32) *UserProfileQuery {
	return q.w("todo_public_visible='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoPublicVisible_NotEqual(v int32) *UserProfileQuery {
	return q.w("todo_public_visible<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoPublicVisible_Less(v int32) *UserProfileQuery {
	return q.w("todo_public_visible<'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoPublicVisible_LessEqual(v int32) *UserProfileQuery {
	return q.w("todo_public_visible<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoPublicVisible_Greater(v int32) *UserProfileQuery {
	return q.w("todo_public_visible>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoPublicVisible_GreaterEqual(v int32) *UserProfileQuery {
	return q.w("todo_public_visible>='" + fmt.Sprint(v) + "'")
}

type UserProfileDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewUserProfileDao(db *DB) (t *UserProfileDao, err error) {
	t = &UserProfileDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *UserProfileDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserProfileDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_profile (update_version,user_id,user_name,todo_public_visible) VALUES (?,?,?,?)")
	return err
}

func (dao *UserProfileDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE user_profile SET update_version=update_version+1,user_id=?,user_name=?,todo_public_visible=? WHERE id=? AND update_version=?")
	return err
}

func (dao *UserProfileDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM user_profile WHERE id=?")
	return err
}

func (dao *UserProfileDao) Insert(ctx context.Context, tx *wrap.Tx, e *UserProfile) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UpdateVersion, e.UserId, e.UserName, e.TodoPublicVisible)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *UserProfileDao) Update(ctx context.Context, tx *wrap.Tx, e *UserProfile) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, e.UserId, e.UserName, e.TodoPublicVisible, e.Id, e.UpdateVersion)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserProfileDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserProfileDao) scanRow(row *wrap.Row) (*UserProfile, error) {
	e := &UserProfile{}
	err := row.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.UserId, &e.UserName, &e.TodoPublicVisible)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *UserProfileDao) scanRows(rows *wrap.Rows) (list []*UserProfile, err error) {
	list = make([]*UserProfile, 0)
	for rows.Next() {
		e := UserProfile{}
		err = rows.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.UserId, &e.UserName, &e.TodoPublicVisible)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *UserProfileDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*UserProfile, error) {
	querySql := "SELECT " + USER_PROFILE_ALL_FIELDS_STRING + " FROM user_profile " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserProfileDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserProfile, err error) {
	querySql := "SELECT " + USER_PROFILE_ALL_FIELDS_STRING + " FROM user_profile " + query
	var rows *wrap.Rows
	if tx == nil {
		rows, err = dao.db.Query(ctx, querySql)
	} else {
		rows, err = tx.Query(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *UserProfileDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM user_profile " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *UserProfileDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM user_profile " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *UserProfileDao) GetQuery() *UserProfileQuery {
	return NewUserProfileQuery(dao)
}

type DB struct {
	wrap.DB
	Todo        *TodoDao
	UserProfile *UserProfileDao
}

func NewDB(connectionString string) (d *DB, err error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connectionString nil")
	}

	d = &DB{}

	db, err := wrap.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	d.DB = *db

	err = d.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	d.Todo, err = NewTodoDao(d)
	if err != nil {
		return nil, err
	}

	d.UserProfile, err = NewUserProfileDao(d)
	if err != nil {
		return nil, err
	}

	return d, nil
}