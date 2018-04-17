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
	"os"
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

func (q *BaseQuery) groupBy(fields ...string) {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = v
	}
}

func (q *BaseQuery) setLimit(startIncluded int64, count int64) {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
}

func (q *BaseQuery) orderBy(fieldName string, asc bool) {
	if q.order != "" {
		q.order += ","
	}
	q.order += fieldName + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}
}

func (q *BaseQuery) orderByGroupCount(asc bool) {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}
}

func (q *BaseQuery) setWhere(format string, a ...interface{}) {
	q.where += fmt.Sprintf(format, a...)
}

const OPERATION_TABLE_NAME = "operation"

type OPERATION_FIELD string

const OPERATION_FIELD_ID = OPERATION_FIELD("id")
const OPERATION_FIELD_CREATE_TIME = OPERATION_FIELD("create_time")
const OPERATION_FIELD_OPERATION_TYPE = OPERATION_FIELD("operation_type")
const OPERATION_FIELD_USER_AGENT = OPERATION_FIELD("user_agent")
const OPERATION_FIELD_USER_ID = OPERATION_FIELD("user_id")
const OPERATION_FIELD_API_NAME = OPERATION_FIELD("api_name")
const OPERATION_FIELD_FRIEND_ID = OPERATION_FIELD("friend_id")
const OPERATION_FIELD_TODO_ID = OPERATION_FIELD("todo_id")
const OPERATION_FIELD_TODO_ITEM = OPERATION_FIELD("todo_item")
const OPERATION_FIELD_USER_PROFILE = OPERATION_FIELD("user_profile")

const OPERATION_ALL_FIELDS_STRING = "id,create_time,operation_type,user_agent,user_id,api_name,friend_id,todo_id,todo_item,user_profile"

type Operation struct {
	Id            uint64 //size=20
	CreateTime    time.Time
	OperationType string //size=128
	UserAgent     string //size=256
	UserId        string //size=128
	ApiName       string //size=128
	FriendId      string //size=128
	TodoId        string //size=128
	TodoItem      string //size=1024
	UserProfile   string //size=1024
}

type OperationQuery struct {
	BaseQuery
	dao *OperationDao
}

func NewOperationQuery(dao *OperationDao) *OperationQuery {
	q := &OperationQuery{}
	q.dao = dao

	return q
}

func (q *OperationQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*Operation, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *OperationQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*Operation, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *OperationQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *OperationQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *OperationQuery) ForUpdate() *OperationQuery {
	q.forUpdate = true
	return q
}

func (q *OperationQuery) ForShare() *OperationQuery {
	q.forShare = true
	return q
}

func (q *OperationQuery) GroupBy(fields ...OPERATION_FIELD) *OperationQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *OperationQuery) Limit(startIncluded int64, count int64) *OperationQuery {
	q.setLimit(startIncluded, count)
	return q
}

func (q *OperationQuery) OrderBy(fieldName OPERATION_FIELD, asc bool) *OperationQuery {
	q.orderBy(string(fieldName), asc)
	return q
}

func (q *OperationQuery) OrderByGroupCount(asc bool) *OperationQuery {
	q.orderByGroupCount(asc)
	return q
}

func (q *OperationQuery) w(format string, a ...interface{}) *OperationQuery {
	q.setWhere(format, a...)
	return q
}

func (q *OperationQuery) Left() *OperationQuery  { return q.w(" ( ") }
func (q *OperationQuery) Right() *OperationQuery { return q.w(" ) ") }
func (q *OperationQuery) And() *OperationQuery   { return q.w(" AND ") }
func (q *OperationQuery) Or() *OperationQuery    { return q.w(" OR ") }
func (q *OperationQuery) Not() *OperationQuery   { return q.w(" NOT ") }

func (q *OperationQuery) Id_Equal(v uint64) *OperationQuery { return q.w("id='" + fmt.Sprint(v) + "'") }
func (q *OperationQuery) Id_NotEqual(v uint64) *OperationQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) Id_Less(v uint64) *OperationQuery { return q.w("id<'" + fmt.Sprint(v) + "'") }
func (q *OperationQuery) Id_LessEqual(v uint64) *OperationQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) Id_Greater(v uint64) *OperationQuery { return q.w("id>'" + fmt.Sprint(v) + "'") }
func (q *OperationQuery) Id_GreaterEqual(v uint64) *OperationQuery {
	return q.w("id>='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) CreateTime_Equal(v time.Time) *OperationQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) CreateTime_NotEqual(v time.Time) *OperationQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) CreateTime_Less(v time.Time) *OperationQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) CreateTime_LessEqual(v time.Time) *OperationQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) CreateTime_Greater(v time.Time) *OperationQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) CreateTime_GreaterEqual(v time.Time) *OperationQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) OperationType_Equal(v string) *OperationQuery {
	return q.w("operation_type='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) OperationType_NotEqual(v string) *OperationQuery {
	return q.w("operation_type<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) UserAgent_Equal(v string) *OperationQuery {
	return q.w("user_agent='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) UserAgent_NotEqual(v string) *OperationQuery {
	return q.w("user_agent<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) UserId_Equal(v string) *OperationQuery {
	return q.w("user_id='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) UserId_NotEqual(v string) *OperationQuery {
	return q.w("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) ApiName_Equal(v string) *OperationQuery {
	return q.w("api_name='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) ApiName_NotEqual(v string) *OperationQuery {
	return q.w("api_name<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) FriendId_Equal(v string) *OperationQuery {
	return q.w("friend_id='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) FriendId_NotEqual(v string) *OperationQuery {
	return q.w("friend_id<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) TodoId_Equal(v string) *OperationQuery {
	return q.w("todo_id='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) TodoId_NotEqual(v string) *OperationQuery {
	return q.w("todo_id<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) TodoItem_Equal(v string) *OperationQuery {
	return q.w("todo_item='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) TodoItem_NotEqual(v string) *OperationQuery {
	return q.w("todo_item<>'" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) UserProfile_Equal(v string) *OperationQuery {
	return q.w("user_profile='" + fmt.Sprint(v) + "'")
}
func (q *OperationQuery) UserProfile_NotEqual(v string) *OperationQuery {
	return q.w("user_profile<>'" + fmt.Sprint(v) + "'")
}

type OperationDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewOperationDao(db *DB) (t *OperationDao, err error) {
	t = &OperationDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *OperationDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}

func (dao *OperationDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO operation (operation_type,user_agent,user_id,api_name,friend_id,todo_id,todo_item,user_profile) VALUES (?,?,?,?,?,?,?,?)")
	return err
}

func (dao *OperationDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM operation WHERE id=?")
	return err
}

func (dao *OperationDao) Insert(ctx context.Context, tx *wrap.Tx, e *Operation) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.OperationType, e.UserAgent, e.UserId, e.ApiName, e.FriendId, e.TodoId, e.TodoItem, e.UserProfile)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *OperationDao) Delete(ctx context.Context, tx *wrap.Tx, id uint64) (err error) {
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

func (dao *OperationDao) scanRow(row *wrap.Row) (*Operation, error) {
	e := &Operation{}
	err := row.Scan(&e.Id, &e.CreateTime, &e.OperationType, &e.UserAgent, &e.UserId, &e.ApiName, &e.FriendId, &e.TodoId, &e.TodoItem, &e.UserProfile)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *OperationDao) scanRows(rows *wrap.Rows) (list []*Operation, err error) {
	list = make([]*Operation, 0)
	for rows.Next() {
		e := Operation{}
		err = rows.Scan(&e.Id, &e.CreateTime, &e.OperationType, &e.UserAgent, &e.UserId, &e.ApiName, &e.FriendId, &e.TodoId, &e.TodoItem, &e.UserProfile)
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

func (dao *OperationDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*Operation, error) {
	querySql := "SELECT " + OPERATION_ALL_FIELDS_STRING + " FROM operation " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *OperationDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*Operation, err error) {
	querySql := "SELECT " + OPERATION_ALL_FIELDS_STRING + " FROM operation " + query
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

func (dao *OperationDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM operation " + query
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

func (dao *OperationDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM operation " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *OperationDao) GetQuery() *OperationQuery {
	return NewOperationQuery(dao)
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

type Todo struct {
	Id            uint64 //size=20
	CreateTime    time.Time
	UpdateTime    time.Time
	UpdateVersion int64  //size=20
	TodoId        string //size=128
	UserId        string //size=128
	TodoCategory  string //size=128
	TodoTitle     string //size=128
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
	q.setLimit(startIncluded, count)
	return q
}

func (q *TodoQuery) OrderBy(fieldName TODO_FIELD, asc bool) *TodoQuery {
	q.orderBy(string(fieldName), asc)
	return q
}

func (q *TodoQuery) OrderByGroupCount(asc bool) *TodoQuery {
	q.orderByGroupCount(asc)
	return q
}

func (q *TodoQuery) w(format string, a ...interface{}) *TodoQuery {
	q.setWhere(format, a...)
	return q
}

func (q *TodoQuery) Left() *TodoQuery  { return q.w(" ( ") }
func (q *TodoQuery) Right() *TodoQuery { return q.w(" ) ") }
func (q *TodoQuery) And() *TodoQuery   { return q.w(" AND ") }
func (q *TodoQuery) Or() *TodoQuery    { return q.w(" OR ") }
func (q *TodoQuery) Not() *TodoQuery   { return q.w(" NOT ") }

func (q *TodoQuery) Id_Equal(v uint64) *TodoQuery        { return q.w("id='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_NotEqual(v uint64) *TodoQuery     { return q.w("id<>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_Less(v uint64) *TodoQuery         { return q.w("id<'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_LessEqual(v uint64) *TodoQuery    { return q.w("id<='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_Greater(v uint64) *TodoQuery      { return q.w("id>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) Id_GreaterEqual(v uint64) *TodoQuery { return q.w("id>='" + fmt.Sprint(v) + "'") }
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
func (q *TodoQuery) UserId_Equal(v string) *TodoQuery    { return q.w("user_id='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) UserId_NotEqual(v string) *TodoQuery { return q.w("user_id<>'" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoCategory_Equal(v string) *TodoQuery {
	return q.w("todo_category='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoCategory_NotEqual(v string) *TodoQuery {
	return q.w("todo_category<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_Equal(v string) *TodoQuery {
	return q.w("todo_title='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoTitle_NotEqual(v string) *TodoQuery {
	return q.w("todo_title<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoDesc_Equal(v string) *TodoQuery { return q.w("todo_desc='" + fmt.Sprint(v) + "'") }
func (q *TodoQuery) TodoDesc_NotEqual(v string) *TodoQuery {
	return q.w("todo_desc<>'" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_Equal(v string) *TodoQuery {
	return q.w("todo_status='" + fmt.Sprint(v) + "'")
}
func (q *TodoQuery) TodoStatus_NotEqual(v string) *TodoQuery {
	return q.w("todo_status<>'" + fmt.Sprint(v) + "'")
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

type TodoUpdate struct {
	dao    *TodoDao
	keys   []string
	values []interface{}
}

func NewTodoUpdate(dao *TodoDao) *TodoUpdate {
	q := &TodoUpdate{}
	q.dao = dao
	q.keys = make([]string, 0)
	q.values = make([]interface{}, 0)

	return q
}

func (u *TodoUpdate) Update(ctx context.Context, tx *wrap.Tx, id uint64) (err error) {
	if len(u.keys) == 0 {
		err = fmt.Errorf("TodoUpdate没有设置更新字段")
		u.dao.logger.Error("TodoUpdate", zap.Error(err))
		return err
	}
	s := "UPDATE todo SET " + strings.Join(u.keys, ",") + " WHERE id=?"
	v := append(u.values, id)
	if tx == nil {
		_, err = u.dao.db.Exec(ctx, s, v)
	} else {
		_, err = tx.Exec(ctx, s, v)
	}

	if err != nil {
		return err
	}

	return nil
}

func (u *TodoUpdate) TodoId(v string) *TodoUpdate {
	u.keys = append(u.keys, "todo_id=?")
	u.values = append(u.values, v)
	return u
}

func (u *TodoUpdate) UserId(v string) *TodoUpdate {
	u.keys = append(u.keys, "user_id=?")
	u.values = append(u.values, v)
	return u
}

func (u *TodoUpdate) TodoCategory(v string) *TodoUpdate {
	u.keys = append(u.keys, "todo_category=?")
	u.values = append(u.values, v)
	return u
}

func (u *TodoUpdate) TodoTitle(v string) *TodoUpdate {
	u.keys = append(u.keys, "todo_title=?")
	u.values = append(u.values, v)
	return u
}

func (u *TodoUpdate) TodoDesc(v string) *TodoUpdate {
	u.keys = append(u.keys, "todo_desc=?")
	u.values = append(u.values, v)
	return u
}

func (u *TodoUpdate) TodoStatus(v string) *TodoUpdate {
	u.keys = append(u.keys, "todo_status=?")
	u.values = append(u.values, v)
	return u
}

func (u *TodoUpdate) TodoPriority(v int32) *TodoUpdate {
	u.keys = append(u.keys, "todo_priority=?")
	u.values = append(u.values, v)
	return u
}

type TodoDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
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

func (dao *TodoDao) Delete(ctx context.Context, tx *wrap.Tx, id uint64) (err error) {
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

func (dao *TodoDao) GetUpdate() *TodoUpdate {
	return NewTodoUpdate(dao)
}

const USER_PROFILE_TABLE_NAME = "user_profile"

type USER_PROFILE_FIELD string

const USER_PROFILE_FIELD_ID = USER_PROFILE_FIELD("id")
const USER_PROFILE_FIELD_CREATE_TIME = USER_PROFILE_FIELD("create_time")
const USER_PROFILE_FIELD_UPDATE_TIME = USER_PROFILE_FIELD("update_time")
const USER_PROFILE_FIELD_UPDATE_VERSION = USER_PROFILE_FIELD("update_version")
const USER_PROFILE_FIELD_USER_ID = USER_PROFILE_FIELD("user_id")
const USER_PROFILE_FIELD_USER_NAME = USER_PROFILE_FIELD("user_name")
const USER_PROFILE_FIELD_TODO_VISIBILITY = USER_PROFILE_FIELD("todo_visibility")

const USER_PROFILE_ALL_FIELDS_STRING = "id,create_time,update_time,update_version,user_id,user_name,todo_visibility"

type UserProfile struct {
	Id             uint64 //size=20
	CreateTime     time.Time
	UpdateTime     time.Time
	UpdateVersion  int64  //size=20
	UserId         string //size=128
	UserName       string //size=128
	TodoVisibility string //size=32
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
	q.setLimit(startIncluded, count)
	return q
}

func (q *UserProfileQuery) OrderBy(fieldName USER_PROFILE_FIELD, asc bool) *UserProfileQuery {
	q.orderBy(string(fieldName), asc)
	return q
}

func (q *UserProfileQuery) OrderByGroupCount(asc bool) *UserProfileQuery {
	q.orderByGroupCount(asc)
	return q
}

func (q *UserProfileQuery) w(format string, a ...interface{}) *UserProfileQuery {
	q.setWhere(format, a...)
	return q
}

func (q *UserProfileQuery) Left() *UserProfileQuery  { return q.w(" ( ") }
func (q *UserProfileQuery) Right() *UserProfileQuery { return q.w(" ) ") }
func (q *UserProfileQuery) And() *UserProfileQuery   { return q.w(" AND ") }
func (q *UserProfileQuery) Or() *UserProfileQuery    { return q.w(" OR ") }
func (q *UserProfileQuery) Not() *UserProfileQuery   { return q.w(" NOT ") }

func (q *UserProfileQuery) Id_Equal(v uint64) *UserProfileQuery {
	return q.w("id='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) Id_NotEqual(v uint64) *UserProfileQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) Id_Less(v uint64) *UserProfileQuery { return q.w("id<'" + fmt.Sprint(v) + "'") }
func (q *UserProfileQuery) Id_LessEqual(v uint64) *UserProfileQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) Id_Greater(v uint64) *UserProfileQuery {
	return q.w("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) Id_GreaterEqual(v uint64) *UserProfileQuery {
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
func (q *UserProfileQuery) UserName_Equal(v string) *UserProfileQuery {
	return q.w("user_name='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) UserName_NotEqual(v string) *UserProfileQuery {
	return q.w("user_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoVisibility_Equal(v string) *UserProfileQuery {
	return q.w("todo_visibility='" + fmt.Sprint(v) + "'")
}
func (q *UserProfileQuery) TodoVisibility_NotEqual(v string) *UserProfileQuery {
	return q.w("todo_visibility<>'" + fmt.Sprint(v) + "'")
}

type UserProfileUpdate struct {
	dao    *UserProfileDao
	keys   []string
	values []interface{}
}

func NewUserProfileUpdate(dao *UserProfileDao) *UserProfileUpdate {
	q := &UserProfileUpdate{}
	q.dao = dao
	q.keys = make([]string, 0)
	q.values = make([]interface{}, 0)

	return q
}

func (u *UserProfileUpdate) Update(ctx context.Context, tx *wrap.Tx, id uint64) (err error) {
	if len(u.keys) == 0 {
		err = fmt.Errorf("UserProfileUpdate没有设置更新字段")
		u.dao.logger.Error("UserProfileUpdate", zap.Error(err))
		return err
	}
	s := "UPDATE user_profile SET " + strings.Join(u.keys, ",") + " WHERE id=?"
	v := append(u.values, id)
	if tx == nil {
		_, err = u.dao.db.Exec(ctx, s, v)
	} else {
		_, err = tx.Exec(ctx, s, v)
	}

	if err != nil {
		return err
	}

	return nil
}

func (u *UserProfileUpdate) UserId(v string) *UserProfileUpdate {
	u.keys = append(u.keys, "user_id=?")
	u.values = append(u.values, v)
	return u
}

func (u *UserProfileUpdate) UserName(v string) *UserProfileUpdate {
	u.keys = append(u.keys, "user_name=?")
	u.values = append(u.values, v)
	return u
}

func (u *UserProfileUpdate) TodoVisibility(v string) *UserProfileUpdate {
	u.keys = append(u.keys, "todo_visibility=?")
	u.values = append(u.values, v)
	return u
}

type UserProfileDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
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

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserProfileDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_profile (update_version,user_id,user_name,todo_visibility) VALUES (?,?,?,?)")
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

	result, err := stmt.Exec(ctx, e.UpdateVersion, e.UserId, e.UserName, e.TodoVisibility)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *UserProfileDao) Delete(ctx context.Context, tx *wrap.Tx, id uint64) (err error) {
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
	err := row.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.UserId, &e.UserName, &e.TodoVisibility)
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
		err = rows.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.UserId, &e.UserName, &e.TodoVisibility)
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

func (dao *UserProfileDao) GetUpdate() *UserProfileUpdate {
	return NewUserProfileUpdate(dao)
}

type DB struct {
	wrap.DB
	Operation   *OperationDao
	Todo        *TodoDao
	UserProfile *UserProfileDao
}

func NewDB() (d *DB, err error) {
	d = &DB{}

	connectionString := os.Getenv("DB")
	if connectionString == "" {
		return nil, fmt.Errorf("DB env nil")
	}
	connectionString += "/todo?parseTime=true"
	db, err := wrap.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	d.DB = *db

	err = d.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	d.Operation, err = NewOperationDao(d)
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
