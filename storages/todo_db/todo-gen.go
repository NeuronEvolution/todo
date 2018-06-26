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

type QueryBase struct {
	where              *bytes.Buffer
	whereParams        []interface{}
	groupByFields      []string
	groupByOrders      []bool
	orderByFields      []string
	orderByOrders      []bool
	hasLimit           bool
	limitStartIncluded int64
	limitCount         int64
	forUpdate          bool
	forShare           bool
	updateFields       []string
	updateParams       []interface{}
}

func (q *QueryBase) buildSelectQuery() (queryString string, params []interface{}) {
	query := bytes.NewBufferString("")

	where := q.where.String()
	if where != "" {
		query.WriteString(" WHERE ")
		query.WriteString(where)
		params = append(params, q.whereParams...)
	}

	groupByCount := len(q.groupByFields)
	if groupByCount > 0 {
		groupByItems := make([]string, groupByCount)
		for i, v := range q.groupByFields {
			if q.groupByOrders[i] {
				groupByItems[i] = v + " ASC"
			} else {
				groupByItems[i] = v + " DESC"
			}
		}
		query.WriteString(" GROUP BY ")
		query.WriteString(strings.Join(groupByItems, ","))
	}

	orderByCount := len(q.orderByFields)
	if orderByCount > 0 {
		orderByItems := make([]string, orderByCount)
		for i, v := range q.orderByFields {
			if q.orderByOrders[i] {
				orderByItems[i] = v + " ASC"
			} else {
				orderByItems[i] = v + " DESC"
			}
		}
		query.WriteString(" ORDER BY ")
		query.WriteString(strings.Join(orderByItems, ","))
	}

	if q.hasLimit {
		query.WriteString(fmt.Sprintf(" LIMIT %d,%d", q.limitStartIncluded, q.limitCount))
	}

	if q.forUpdate {
		query.WriteString(" FOR UPDATE")
	}

	if q.forShare {
		query.WriteString(" LOCK IN SHARE MODE")
	}

	return query.String(), params
}

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
	QueryBase
	dao *OperationDao
}

func (dao *OperationDao) Query() *OperationQuery {
	q := &OperationQuery{}
	q.dao = dao
	q.where = bytes.NewBufferString("")
	return q
}

func (q *OperationQuery) Left() *OperationQuery {
	q.where.WriteString(" (")
	return q
}

func (q *OperationQuery) Right() *OperationQuery {
	q.where.WriteString(" )")
	return q
}

func (q *OperationQuery) And() *OperationQuery {
	q.where.WriteString(" AND")
	return q
}

func (q *OperationQuery) Or() *OperationQuery {
	q.where.WriteString(" OR")
	return q
}

func (q *OperationQuery) Not() *OperationQuery {
	q.where.WriteString(" NOT")
	return q
}

func (q *OperationQuery) IdEqual(v uint64) *OperationQuery {
	q.where.WriteString(" id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) IdNotEqual(v uint64) *OperationQuery {
	q.where.WriteString(" id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) IdLess(v uint64) *OperationQuery {
	q.where.WriteString(" id<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) IdLessEqual(v uint64) *OperationQuery {
	q.where.WriteString(" id<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) IdGreater(v uint64) *OperationQuery {
	q.where.WriteString(" id>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) IdGreaterEqual(v uint64) *OperationQuery {
	q.where.WriteString(" id>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) IdIn(items []uint64) *OperationQuery {
	q.where.WriteString(" id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) CreateTimeEqual(v time.Time) *OperationQuery {
	q.where.WriteString(" create_time=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) CreateTimeNotEqual(v time.Time) *OperationQuery {
	q.where.WriteString(" create_time<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) CreateTimeLess(v time.Time) *OperationQuery {
	q.where.WriteString(" create_time<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) CreateTimeLessEqual(v time.Time) *OperationQuery {
	q.where.WriteString(" create_time<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) CreateTimeGreater(v time.Time) *OperationQuery {
	q.where.WriteString(" create_time>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) CreateTimeGreaterEqual(v time.Time) *OperationQuery {
	q.where.WriteString(" create_time>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) OperationTypeEqual(v string) *OperationQuery {
	q.where.WriteString(" operation_type=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) OperationTypeNotEqual(v string) *OperationQuery {
	q.where.WriteString(" operation_type<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) OperationTypeIn(items []string) *OperationQuery {
	q.where.WriteString(" operation_type IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) UserAgentEqual(v string) *OperationQuery {
	q.where.WriteString(" user_agent=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) UserAgentNotEqual(v string) *OperationQuery {
	q.where.WriteString(" user_agent<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) UserAgentIn(items []string) *OperationQuery {
	q.where.WriteString(" user_agent IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) UserIdEqual(v string) *OperationQuery {
	q.where.WriteString(" user_id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) UserIdNotEqual(v string) *OperationQuery {
	q.where.WriteString(" user_id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) UserIdIn(items []string) *OperationQuery {
	q.where.WriteString(" user_id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) ApiNameEqual(v string) *OperationQuery {
	q.where.WriteString(" api_name=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) ApiNameNotEqual(v string) *OperationQuery {
	q.where.WriteString(" api_name<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) ApiNameIn(items []string) *OperationQuery {
	q.where.WriteString(" api_name IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) FriendIdEqual(v string) *OperationQuery {
	q.where.WriteString(" friend_id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) FriendIdNotEqual(v string) *OperationQuery {
	q.where.WriteString(" friend_id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) FriendIdIn(items []string) *OperationQuery {
	q.where.WriteString(" friend_id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) TodoIdEqual(v string) *OperationQuery {
	q.where.WriteString(" todo_id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) TodoIdNotEqual(v string) *OperationQuery {
	q.where.WriteString(" todo_id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) TodoIdIn(items []string) *OperationQuery {
	q.where.WriteString(" todo_id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) TodoItemEqual(v string) *OperationQuery {
	q.where.WriteString(" todo_item=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) TodoItemNotEqual(v string) *OperationQuery {
	q.where.WriteString(" todo_item<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) TodoItemIn(items []string) *OperationQuery {
	q.where.WriteString(" todo_item IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) UserProfileEqual(v string) *OperationQuery {
	q.where.WriteString(" user_profile=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) UserProfileNotEqual(v string) *OperationQuery {
	q.where.WriteString(" user_profile<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *OperationQuery) UserProfileIn(items []string) *OperationQuery {
	q.where.WriteString(" user_profile IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *OperationQuery) GroupByOperationType(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "operation_type")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) GroupByUserAgent(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "user_agent")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) GroupByUserId(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "user_id")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) GroupByApiName(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "api_name")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) GroupByFriendId(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "friend_id")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) GroupByTodoId(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "todo_id")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) GroupByTodoItem(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "todo_item")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) GroupByUserProfile(asc bool) *OperationQuery {
	q.groupByFields = append(q.groupByFields, "user_profile")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *OperationQuery) OrderById(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByCreateTime(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "create_time")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByOperationType(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "operation_type")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByUserAgent(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "user_agent")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByUserId(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "user_id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByApiName(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "api_name")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByFriendId(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "friend_id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByTodoId(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "todo_id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByTodoItem(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "todo_item")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByUserProfile(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "user_profile")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) OrderByGroupCount(asc bool) *OperationQuery {
	q.orderByFields = append(q.orderByFields, "count(*)")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *OperationQuery) Limit(startIncluded int64, count int64) *OperationQuery {
	q.hasLimit = true
	q.limitStartIncluded = startIncluded
	q.limitCount = count
	return q
}

func (q *OperationQuery) ForUpdate() *OperationQuery {
	q.forUpdate = true
	return q
}

func (q *OperationQuery) ForShare() *OperationQuery {
	q.forShare = true
	return q
}

func (q *OperationQuery) Select(ctx context.Context, tx *wrap.Tx) (e *Operation, err error) {
	if !q.hasLimit {
		q.limitCount = 1
		q.hasLimit = true
	}

	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT id,create_time,operation_type,user_agent,user_id,api_name,friend_id,todo_id,todo_item,user_profile FROM operation ")
	query.WriteString(queryString)
	e = &Operation{}
	row := q.dao.db.QueryRow(ctx, tx, query.String(), params...)
	err = row.Scan(&e.Id, &e.CreateTime, &e.OperationType, &e.UserAgent, &e.UserId, &e.ApiName, &e.FriendId, &e.TodoId, &e.TodoItem, &e.UserProfile)
	if err == wrap.ErrNoRows {
		return nil, nil
	}

	return e, err
}

func (q *OperationQuery) SelectList(ctx context.Context, tx *wrap.Tx) (list []*Operation, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT id,create_time,operation_type,user_agent,user_id,api_name,friend_id,todo_id,todo_item,user_profile FROM operation ")
	query.WriteString(queryString)
	rows, err := q.dao.db.Query(ctx, tx, query.String(), params...)
	if err != nil {
		return nil, err
	}
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

func (q *OperationQuery) SelectCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT COUNT(*) FROM operation ")
	query.WriteString(queryString)
	row := q.dao.db.QueryRow(ctx, tx, query.String(), params...)
	err = row.Scan(&count)

	return count, err
}

func (q *OperationQuery) SelectGroupBy(ctx context.Context, tx *wrap.Tx, withCount bool) (rows *wrap.Rows, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT ")
	query.WriteString(strings.Join(q.groupByFields, ","))
	if withCount {
		query.WriteString(",Count(*) ")
	}
	query.WriteString(" FROM operation ")
	query.WriteString(queryString)

	return q.dao.db.Query(ctx, tx, query.String(), params...)
}

func (q *OperationQuery) SetOperationType(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "operation_type")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) SetUserAgent(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "user_agent")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) SetUserId(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "user_id")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) SetApiName(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "api_name")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) SetFriendId(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "friend_id")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) SetTodoId(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "todo_id")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) SetTodoItem(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "todo_item")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) SetUserProfile(v string) *OperationQuery {
	q.updateFields = append(q.updateFields, "user_profile")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *OperationQuery) Update(ctx context.Context, tx *wrap.Tx) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	var params []interface{}
	params = append(params, q.updateParams)
	query.WriteString("UPDATE operation SET ")
	updateItems := make([]string, len(q.updateFields))
	for i, v := range q.updateFields {
		updateItems[i] = v + "=?"
	}
	query.WriteString(strings.Join(updateItems, ","))
	where := q.where.String()
	if where != "" {
		query.WriteString(" WHERE ")
		query.WriteString(where)
		params = append(params, q.whereParams)
	}

	return q.dao.db.Exec(ctx, tx, query.String(), params...)
}

func (q *OperationQuery) Delete(ctx context.Context, tx *wrap.Tx) (result *wrap.Result, err error) {
	query := "DELETE FROM operation WHERE " + q.where.String()
	return q.dao.db.Exec(ctx, tx, query, q.whereParams...)
}

type OperationDao struct {
	logger *zap.Logger
	db     *DB
}

func NewOperationDao(db *DB) (t *OperationDao, err error) {
	t = &OperationDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t, nil
}

func (dao *OperationDao) Insert(ctx context.Context, tx *wrap.Tx, e *Operation) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	query.WriteString("INSERT INTO operation (operation_type,user_agent,user_id,api_name,friend_id,todo_id,todo_item,user_profile) VALUES (?,?,?,?,?,?,?,?)")
	params := []interface{}{e.OperationType, e.UserAgent, e.UserId, e.ApiName, e.FriendId, e.TodoId, e.TodoItem, e.UserProfile}
	return dao.db.Exec(ctx, tx, query.String(), params...)
}

func (dao *OperationDao) BatchInsert(ctx context.Context, tx *wrap.Tx, list []*Operation) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	query.WriteString("INSERT INTO operation (operation_type,user_agent,user_id,api_name,friend_id,todo_id,todo_item,user_profile) VALUES ")
	query.WriteString(wrap.RepeatWithSeparator("(?,?,?,?,?,?,?,?)", len(list), ","))
	params := make([]interface{}, len(list)*8)
	offset := 0
	for _, e := range list {
		params[offset+0] = e.OperationType
		params[offset+1] = e.UserAgent
		params[offset+2] = e.UserId
		params[offset+3] = e.ApiName
		params[offset+4] = e.FriendId
		params[offset+5] = e.TodoId
		params[offset+6] = e.TodoItem
		params[offset+7] = e.UserProfile
		offset += 8
	}

	return dao.db.Exec(ctx, tx, query.String(), params...)
}

func (dao *OperationDao) DeleteById(ctx context.Context, tx *wrap.Tx, id uint64) (result *wrap.Result, err error) {
	query := "DELETE FROM Operation WHERE id=?"
	return dao.db.Exec(ctx, tx, query, id)
}

func (dao *OperationDao) UpdateById(ctx context.Context, tx *wrap.Tx, e *Operation) (result *wrap.Result, err error) {
	query := "UPDATE operation SET operation_type=?,user_agent=?,user_id=?,api_name=?,friend_id=?,todo_id=?,todo_item=?,user_profile=? WHERE id=?"
	params := []interface{}{e.OperationType, e.UserAgent, e.UserId, e.ApiName, e.FriendId, e.TodoId, e.TodoItem, e.UserProfile, e.Id}
	return dao.db.Exec(ctx, tx, query, params...)
}

func (dao *OperationDao) SelectById(ctx context.Context, tx *wrap.Tx, id int64) (e *Operation, err error) {
	query := "SELECT id,create_time,operation_type,user_agent,user_id,api_name,friend_id,todo_id,todo_item,user_profile FROM operation WHERE id=?"
	row := dao.db.QueryRow(ctx, tx, query, id)
	e = &Operation{}
	err = row.Scan(&e.Id, &e.CreateTime, &e.OperationType, &e.UserAgent, &e.UserId, &e.ApiName, &e.FriendId, &e.TodoId, &e.TodoItem, &e.UserProfile)
	if err == wrap.ErrNoRows {
		return nil, nil
	}
	return e, err
}

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
	QueryBase
	dao *TodoDao
}

func (dao *TodoDao) Query() *TodoQuery {
	q := &TodoQuery{}
	q.dao = dao
	q.where = bytes.NewBufferString("")
	return q
}

func (q *TodoQuery) Left() *TodoQuery {
	q.where.WriteString(" (")
	return q
}

func (q *TodoQuery) Right() *TodoQuery {
	q.where.WriteString(" )")
	return q
}

func (q *TodoQuery) And() *TodoQuery {
	q.where.WriteString(" AND")
	return q
}

func (q *TodoQuery) Or() *TodoQuery {
	q.where.WriteString(" OR")
	return q
}

func (q *TodoQuery) Not() *TodoQuery {
	q.where.WriteString(" NOT")
	return q
}

func (q *TodoQuery) IdEqual(v uint64) *TodoQuery {
	q.where.WriteString(" id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) IdNotEqual(v uint64) *TodoQuery {
	q.where.WriteString(" id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) IdLess(v uint64) *TodoQuery {
	q.where.WriteString(" id<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) IdLessEqual(v uint64) *TodoQuery {
	q.where.WriteString(" id<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) IdGreater(v uint64) *TodoQuery {
	q.where.WriteString(" id>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) IdGreaterEqual(v uint64) *TodoQuery {
	q.where.WriteString(" id>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) IdIn(items []uint64) *TodoQuery {
	q.where.WriteString(" id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) CreateTimeEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" create_time=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) CreateTimeNotEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" create_time<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) CreateTimeLess(v time.Time) *TodoQuery {
	q.where.WriteString(" create_time<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) CreateTimeLessEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" create_time<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) CreateTimeGreater(v time.Time) *TodoQuery {
	q.where.WriteString(" create_time>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) CreateTimeGreaterEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" create_time>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateTimeEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" update_time=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateTimeNotEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" update_time<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateTimeLess(v time.Time) *TodoQuery {
	q.where.WriteString(" update_time<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateTimeLessEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" update_time<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateTimeGreater(v time.Time) *TodoQuery {
	q.where.WriteString(" update_time>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateTimeGreaterEqual(v time.Time) *TodoQuery {
	q.where.WriteString(" update_time>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateVersionEqual(v int64) *TodoQuery {
	q.where.WriteString(" update_version=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateVersionNotEqual(v int64) *TodoQuery {
	q.where.WriteString(" update_version<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateVersionLess(v int64) *TodoQuery {
	q.where.WriteString(" update_version<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateVersionLessEqual(v int64) *TodoQuery {
	q.where.WriteString(" update_version<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateVersionGreater(v int64) *TodoQuery {
	q.where.WriteString(" update_version>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateVersionGreaterEqual(v int64) *TodoQuery {
	q.where.WriteString(" update_version>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UpdateVersionIn(items []int64) *TodoQuery {
	q.where.WriteString(" update_version IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) TodoIdEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoIdNotEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoIdIn(items []string) *TodoQuery {
	q.where.WriteString(" todo_id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) UserIdEqual(v string) *TodoQuery {
	q.where.WriteString(" user_id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UserIdNotEqual(v string) *TodoQuery {
	q.where.WriteString(" user_id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) UserIdIn(items []string) *TodoQuery {
	q.where.WriteString(" user_id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) TodoCategoryEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_category=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoCategoryNotEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_category<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoCategoryIn(items []string) *TodoQuery {
	q.where.WriteString(" todo_category IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) TodoTitleEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_title=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoTitleNotEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_title<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoTitleIn(items []string) *TodoQuery {
	q.where.WriteString(" todo_title IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) TodoDescEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_desc=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoDescNotEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_desc<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoDescIn(items []string) *TodoQuery {
	q.where.WriteString(" todo_desc IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) TodoStatusEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_status=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoStatusNotEqual(v string) *TodoQuery {
	q.where.WriteString(" todo_status<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoStatusIn(items []string) *TodoQuery {
	q.where.WriteString(" todo_status IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) TodoPriorityEqual(v int32) *TodoQuery {
	q.where.WriteString(" todo_priority=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoPriorityNotEqual(v int32) *TodoQuery {
	q.where.WriteString(" todo_priority<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoPriorityLess(v int32) *TodoQuery {
	q.where.WriteString(" todo_priority<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoPriorityLessEqual(v int32) *TodoQuery {
	q.where.WriteString(" todo_priority<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoPriorityGreater(v int32) *TodoQuery {
	q.where.WriteString(" todo_priority>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoPriorityGreaterEqual(v int32) *TodoQuery {
	q.where.WriteString(" todo_priority>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *TodoQuery) TodoPriorityIn(items []int32) *TodoQuery {
	q.where.WriteString(" todo_priority IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *TodoQuery) GroupByUserId(asc bool) *TodoQuery {
	q.groupByFields = append(q.groupByFields, "user_id")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *TodoQuery) GroupByTodoCategory(asc bool) *TodoQuery {
	q.groupByFields = append(q.groupByFields, "todo_category")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *TodoQuery) GroupByTodoTitle(asc bool) *TodoQuery {
	q.groupByFields = append(q.groupByFields, "todo_title")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *TodoQuery) GroupByTodoDesc(asc bool) *TodoQuery {
	q.groupByFields = append(q.groupByFields, "todo_desc")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *TodoQuery) GroupByTodoStatus(asc bool) *TodoQuery {
	q.groupByFields = append(q.groupByFields, "todo_status")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *TodoQuery) GroupByTodoPriority(asc bool) *TodoQuery {
	q.groupByFields = append(q.groupByFields, "todo_priority")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *TodoQuery) OrderById(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByCreateTime(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "create_time")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByUpdateTime(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "update_time")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByTodoId(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "todo_id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByUserId(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "user_id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByTodoCategory(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "todo_category")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByTodoTitle(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "todo_title")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByTodoDesc(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "todo_desc")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByTodoStatus(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "todo_status")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByTodoPriority(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "todo_priority")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) OrderByGroupCount(asc bool) *TodoQuery {
	q.orderByFields = append(q.orderByFields, "count(*)")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *TodoQuery) Limit(startIncluded int64, count int64) *TodoQuery {
	q.hasLimit = true
	q.limitStartIncluded = startIncluded
	q.limitCount = count
	return q
}

func (q *TodoQuery) ForUpdate() *TodoQuery {
	q.forUpdate = true
	return q
}

func (q *TodoQuery) ForShare() *TodoQuery {
	q.forShare = true
	return q
}

func (q *TodoQuery) Select(ctx context.Context, tx *wrap.Tx) (e *Todo, err error) {
	if !q.hasLimit {
		q.limitCount = 1
		q.hasLimit = true
	}

	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT id,create_time,update_time,update_version,todo_id,user_id,todo_category,todo_title,todo_desc,todo_status,todo_priority FROM todo ")
	query.WriteString(queryString)
	e = &Todo{}
	row := q.dao.db.QueryRow(ctx, tx, query.String(), params...)
	err = row.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.TodoId, &e.UserId, &e.TodoCategory, &e.TodoTitle, &e.TodoDesc, &e.TodoStatus, &e.TodoPriority)
	if err == wrap.ErrNoRows {
		return nil, nil
	}

	return e, err
}

func (q *TodoQuery) SelectList(ctx context.Context, tx *wrap.Tx) (list []*Todo, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT id,create_time,update_time,update_version,todo_id,user_id,todo_category,todo_title,todo_desc,todo_status,todo_priority FROM todo ")
	query.WriteString(queryString)
	rows, err := q.dao.db.Query(ctx, tx, query.String(), params...)
	if err != nil {
		return nil, err
	}
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

func (q *TodoQuery) SelectCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT COUNT(*) FROM todo ")
	query.WriteString(queryString)
	row := q.dao.db.QueryRow(ctx, tx, query.String(), params...)
	err = row.Scan(&count)

	return count, err
}

func (q *TodoQuery) SelectGroupBy(ctx context.Context, tx *wrap.Tx, withCount bool) (rows *wrap.Rows, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT ")
	query.WriteString(strings.Join(q.groupByFields, ","))
	if withCount {
		query.WriteString(",Count(*) ")
	}
	query.WriteString(" FROM todo ")
	query.WriteString(queryString)

	return q.dao.db.Query(ctx, tx, query.String(), params...)
}

func (q *TodoQuery) SetUpdateVersion(v int64) *TodoQuery {
	q.updateFields = append(q.updateFields, "update_version")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) SetTodoId(v string) *TodoQuery {
	q.updateFields = append(q.updateFields, "todo_id")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) SetUserId(v string) *TodoQuery {
	q.updateFields = append(q.updateFields, "user_id")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) SetTodoCategory(v string) *TodoQuery {
	q.updateFields = append(q.updateFields, "todo_category")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) SetTodoTitle(v string) *TodoQuery {
	q.updateFields = append(q.updateFields, "todo_title")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) SetTodoDesc(v string) *TodoQuery {
	q.updateFields = append(q.updateFields, "todo_desc")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) SetTodoStatus(v string) *TodoQuery {
	q.updateFields = append(q.updateFields, "todo_status")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) SetTodoPriority(v int32) *TodoQuery {
	q.updateFields = append(q.updateFields, "todo_priority")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *TodoQuery) Update(ctx context.Context, tx *wrap.Tx) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	var params []interface{}
	params = append(params, q.updateParams)
	query.WriteString("UPDATE todo SET ")
	updateItems := make([]string, len(q.updateFields))
	for i, v := range q.updateFields {
		updateItems[i] = v + "=?"
	}
	query.WriteString(strings.Join(updateItems, ","))
	where := q.where.String()
	if where != "" {
		query.WriteString(" WHERE ")
		query.WriteString(where)
		params = append(params, q.whereParams)
	}

	return q.dao.db.Exec(ctx, tx, query.String(), params...)
}

func (q *TodoQuery) Delete(ctx context.Context, tx *wrap.Tx) (result *wrap.Result, err error) {
	query := "DELETE FROM todo WHERE " + q.where.String()
	return q.dao.db.Exec(ctx, tx, query, q.whereParams...)
}

type TodoDao struct {
	logger *zap.Logger
	db     *DB
}

func NewTodoDao(db *DB) (t *TodoDao, err error) {
	t = &TodoDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t, nil
}

func (dao *TodoDao) Insert(ctx context.Context, tx *wrap.Tx, e *Todo, onDuplicatedKeyUpdate bool) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	query.WriteString("INSERT INTO todo (update_version,todo_id,user_id,todo_category,todo_title,todo_desc,todo_status,todo_priority) VALUES (?,?,?,?,?,?,?,?)")
	if onDuplicatedKeyUpdate {
		query.WriteString(" ON DUPLICATED KEY UPDATE update_version=VALUES(update_version),user_id=VALUES(user_id),todo_category=VALUES(todo_category),todo_title=VALUES(todo_title),todo_desc=VALUES(todo_desc),todo_status=VALUES(todo_status),todo_priority=VALUES(todo_priority)")
	}
	params := []interface{}{e.UpdateVersion, e.TodoId, e.UserId, e.TodoCategory, e.TodoTitle, e.TodoDesc, e.TodoStatus, e.TodoPriority}
	return dao.db.Exec(ctx, tx, query.String(), params...)
}

func (dao *TodoDao) BatchInsert(ctx context.Context, tx *wrap.Tx, list []*Todo, onDuplicatedKeyUpdate bool) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	query.WriteString("INSERT INTO todo (update_version,todo_id,user_id,todo_category,todo_title,todo_desc,todo_status,todo_priority) VALUES ")
	query.WriteString(wrap.RepeatWithSeparator("(?,?,?,?,?,?,?,?)", len(list), ","))
	if onDuplicatedKeyUpdate {
		query.WriteString(" ON DUPLICATED KEY UPDATE update_version=VALUES(update_version),user_id=VALUES(user_id),todo_category=VALUES(todo_category),todo_title=VALUES(todo_title),todo_desc=VALUES(todo_desc),todo_status=VALUES(todo_status),todo_priority=VALUES(todo_priority)")
	}
	params := make([]interface{}, len(list)*8)
	offset := 0
	for _, e := range list {
		params[offset+0] = e.UpdateVersion
		params[offset+1] = e.TodoId
		params[offset+2] = e.UserId
		params[offset+3] = e.TodoCategory
		params[offset+4] = e.TodoTitle
		params[offset+5] = e.TodoDesc
		params[offset+6] = e.TodoStatus
		params[offset+7] = e.TodoPriority
		offset += 8
	}

	return dao.db.Exec(ctx, tx, query.String(), params...)
}

func (dao *TodoDao) DeleteById(ctx context.Context, tx *wrap.Tx, id uint64) (result *wrap.Result, err error) {
	query := "DELETE FROM Todo WHERE id=?"
	return dao.db.Exec(ctx, tx, query, id)
}

func (dao *TodoDao) UpdateById(ctx context.Context, tx *wrap.Tx, e *Todo) (result *wrap.Result, err error) {
	query := "UPDATE todo SET update_version=?,todo_id=?,user_id=?,todo_category=?,todo_title=?,todo_desc=?,todo_status=?,todo_priority=? WHERE id=?"
	params := []interface{}{e.UpdateVersion, e.TodoId, e.UserId, e.TodoCategory, e.TodoTitle, e.TodoDesc, e.TodoStatus, e.TodoPriority, e.Id}
	return dao.db.Exec(ctx, tx, query, params...)
}

func (dao *TodoDao) SelectById(ctx context.Context, tx *wrap.Tx, id int64) (e *Todo, err error) {
	query := "SELECT id,create_time,update_time,update_version,todo_id,user_id,todo_category,todo_title,todo_desc,todo_status,todo_priority FROM todo WHERE id=?"
	row := dao.db.QueryRow(ctx, tx, query, id)
	e = &Todo{}
	err = row.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.TodoId, &e.UserId, &e.TodoCategory, &e.TodoTitle, &e.TodoDesc, &e.TodoStatus, &e.TodoPriority)
	if err == wrap.ErrNoRows {
		return nil, nil
	}
	return e, err
}

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
	QueryBase
	dao *UserProfileDao
}

func (dao *UserProfileDao) Query() *UserProfileQuery {
	q := &UserProfileQuery{}
	q.dao = dao
	q.where = bytes.NewBufferString("")
	return q
}

func (q *UserProfileQuery) Left() *UserProfileQuery {
	q.where.WriteString(" (")
	return q
}

func (q *UserProfileQuery) Right() *UserProfileQuery {
	q.where.WriteString(" )")
	return q
}

func (q *UserProfileQuery) And() *UserProfileQuery {
	q.where.WriteString(" AND")
	return q
}

func (q *UserProfileQuery) Or() *UserProfileQuery {
	q.where.WriteString(" OR")
	return q
}

func (q *UserProfileQuery) Not() *UserProfileQuery {
	q.where.WriteString(" NOT")
	return q
}

func (q *UserProfileQuery) IdEqual(v uint64) *UserProfileQuery {
	q.where.WriteString(" id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) IdNotEqual(v uint64) *UserProfileQuery {
	q.where.WriteString(" id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) IdLess(v uint64) *UserProfileQuery {
	q.where.WriteString(" id<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) IdLessEqual(v uint64) *UserProfileQuery {
	q.where.WriteString(" id<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) IdGreater(v uint64) *UserProfileQuery {
	q.where.WriteString(" id>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) IdGreaterEqual(v uint64) *UserProfileQuery {
	q.where.WriteString(" id>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) IdIn(items []uint64) *UserProfileQuery {
	q.where.WriteString(" id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *UserProfileQuery) CreateTimeEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" create_time=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) CreateTimeNotEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" create_time<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) CreateTimeLess(v time.Time) *UserProfileQuery {
	q.where.WriteString(" create_time<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) CreateTimeLessEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" create_time<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) CreateTimeGreater(v time.Time) *UserProfileQuery {
	q.where.WriteString(" create_time>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) CreateTimeGreaterEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" create_time>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateTimeEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" update_time=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateTimeNotEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" update_time<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateTimeLess(v time.Time) *UserProfileQuery {
	q.where.WriteString(" update_time<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateTimeLessEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" update_time<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateTimeGreater(v time.Time) *UserProfileQuery {
	q.where.WriteString(" update_time>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateTimeGreaterEqual(v time.Time) *UserProfileQuery {
	q.where.WriteString(" update_time>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateVersionEqual(v int64) *UserProfileQuery {
	q.where.WriteString(" update_version=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateVersionNotEqual(v int64) *UserProfileQuery {
	q.where.WriteString(" update_version<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateVersionLess(v int64) *UserProfileQuery {
	q.where.WriteString(" update_version<?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateVersionLessEqual(v int64) *UserProfileQuery {
	q.where.WriteString(" update_version<=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateVersionGreater(v int64) *UserProfileQuery {
	q.where.WriteString(" update_version>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateVersionGreaterEqual(v int64) *UserProfileQuery {
	q.where.WriteString(" update_version>=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UpdateVersionIn(items []int64) *UserProfileQuery {
	q.where.WriteString(" update_version IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *UserProfileQuery) UserIdEqual(v string) *UserProfileQuery {
	q.where.WriteString(" user_id=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UserIdNotEqual(v string) *UserProfileQuery {
	q.where.WriteString(" user_id<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UserIdIn(items []string) *UserProfileQuery {
	q.where.WriteString(" user_id IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *UserProfileQuery) UserNameEqual(v string) *UserProfileQuery {
	q.where.WriteString(" user_name=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UserNameNotEqual(v string) *UserProfileQuery {
	q.where.WriteString(" user_name<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) UserNameIn(items []string) *UserProfileQuery {
	q.where.WriteString(" user_name IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *UserProfileQuery) TodoVisibilityEqual(v string) *UserProfileQuery {
	q.where.WriteString(" todo_visibility=?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) TodoVisibilityNotEqual(v string) *UserProfileQuery {
	q.where.WriteString(" todo_visibility<>?")
	q.whereParams = append(q.whereParams, v)
	return q
}

func (q *UserProfileQuery) TodoVisibilityIn(items []string) *UserProfileQuery {
	q.where.WriteString(" todo_visibility IN(")
	q.where.WriteString(wrap.RepeatWithSeparator("?", len(items), ","))
	q.where.WriteString(")")
	q.whereParams = append(q.whereParams, items)
	return q
}

func (q *UserProfileQuery) GroupByUserName(asc bool) *UserProfileQuery {
	q.groupByFields = append(q.groupByFields, "user_name")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *UserProfileQuery) GroupByTodoVisibility(asc bool) *UserProfileQuery {
	q.groupByFields = append(q.groupByFields, "todo_visibility")
	q.groupByOrders = append(q.groupByOrders, asc)
	return q
}

func (q *UserProfileQuery) OrderById(asc bool) *UserProfileQuery {
	q.orderByFields = append(q.orderByFields, "id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *UserProfileQuery) OrderByCreateTime(asc bool) *UserProfileQuery {
	q.orderByFields = append(q.orderByFields, "create_time")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *UserProfileQuery) OrderByUpdateTime(asc bool) *UserProfileQuery {
	q.orderByFields = append(q.orderByFields, "update_time")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *UserProfileQuery) OrderByUserId(asc bool) *UserProfileQuery {
	q.orderByFields = append(q.orderByFields, "user_id")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *UserProfileQuery) OrderByUserName(asc bool) *UserProfileQuery {
	q.orderByFields = append(q.orderByFields, "user_name")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *UserProfileQuery) OrderByTodoVisibility(asc bool) *UserProfileQuery {
	q.orderByFields = append(q.orderByFields, "todo_visibility")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *UserProfileQuery) OrderByGroupCount(asc bool) *UserProfileQuery {
	q.orderByFields = append(q.orderByFields, "count(*)")
	q.orderByOrders = append(q.orderByOrders, asc)
	return q
}

func (q *UserProfileQuery) Limit(startIncluded int64, count int64) *UserProfileQuery {
	q.hasLimit = true
	q.limitStartIncluded = startIncluded
	q.limitCount = count
	return q
}

func (q *UserProfileQuery) ForUpdate() *UserProfileQuery {
	q.forUpdate = true
	return q
}

func (q *UserProfileQuery) ForShare() *UserProfileQuery {
	q.forShare = true
	return q
}

func (q *UserProfileQuery) Select(ctx context.Context, tx *wrap.Tx) (e *UserProfile, err error) {
	if !q.hasLimit {
		q.limitCount = 1
		q.hasLimit = true
	}

	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT id,create_time,update_time,update_version,user_id,user_name,todo_visibility FROM user_profile ")
	query.WriteString(queryString)
	e = &UserProfile{}
	row := q.dao.db.QueryRow(ctx, tx, query.String(), params...)
	err = row.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.UserId, &e.UserName, &e.TodoVisibility)
	if err == wrap.ErrNoRows {
		return nil, nil
	}

	return e, err
}

func (q *UserProfileQuery) SelectList(ctx context.Context, tx *wrap.Tx) (list []*UserProfile, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT id,create_time,update_time,update_version,user_id,user_name,todo_visibility FROM user_profile ")
	query.WriteString(queryString)
	rows, err := q.dao.db.Query(ctx, tx, query.String(), params...)
	if err != nil {
		return nil, err
	}
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

func (q *UserProfileQuery) SelectCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT COUNT(*) FROM user_profile ")
	query.WriteString(queryString)
	row := q.dao.db.QueryRow(ctx, tx, query.String(), params...)
	err = row.Scan(&count)

	return count, err
}

func (q *UserProfileQuery) SelectGroupBy(ctx context.Context, tx *wrap.Tx, withCount bool) (rows *wrap.Rows, err error) {
	queryString, params := q.buildSelectQuery()
	query := bytes.NewBufferString("")
	query.WriteString("SELECT ")
	query.WriteString(strings.Join(q.groupByFields, ","))
	if withCount {
		query.WriteString(",Count(*) ")
	}
	query.WriteString(" FROM user_profile ")
	query.WriteString(queryString)

	return q.dao.db.Query(ctx, tx, query.String(), params...)
}

func (q *UserProfileQuery) SetUpdateVersion(v int64) *UserProfileQuery {
	q.updateFields = append(q.updateFields, "update_version")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *UserProfileQuery) SetUserId(v string) *UserProfileQuery {
	q.updateFields = append(q.updateFields, "user_id")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *UserProfileQuery) SetUserName(v string) *UserProfileQuery {
	q.updateFields = append(q.updateFields, "user_name")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *UserProfileQuery) SetTodoVisibility(v string) *UserProfileQuery {
	q.updateFields = append(q.updateFields, "todo_visibility")
	q.updateParams = append(q.updateParams, v)
	return q
}

func (q *UserProfileQuery) Update(ctx context.Context, tx *wrap.Tx) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	var params []interface{}
	params = append(params, q.updateParams)
	query.WriteString("UPDATE user_profile SET ")
	updateItems := make([]string, len(q.updateFields))
	for i, v := range q.updateFields {
		updateItems[i] = v + "=?"
	}
	query.WriteString(strings.Join(updateItems, ","))
	where := q.where.String()
	if where != "" {
		query.WriteString(" WHERE ")
		query.WriteString(where)
		params = append(params, q.whereParams)
	}

	return q.dao.db.Exec(ctx, tx, query.String(), params...)
}

func (q *UserProfileQuery) Delete(ctx context.Context, tx *wrap.Tx) (result *wrap.Result, err error) {
	query := "DELETE FROM user_profile WHERE " + q.where.String()
	return q.dao.db.Exec(ctx, tx, query, q.whereParams...)
}

type UserProfileDao struct {
	logger *zap.Logger
	db     *DB
}

func NewUserProfileDao(db *DB) (t *UserProfileDao, err error) {
	t = &UserProfileDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t, nil
}

func (dao *UserProfileDao) Insert(ctx context.Context, tx *wrap.Tx, e *UserProfile, onDuplicatedKeyUpdate bool) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	query.WriteString("INSERT INTO user_profile (update_version,user_id,user_name,todo_visibility) VALUES (?,?,?,?)")
	if onDuplicatedKeyUpdate {
		query.WriteString(" ON DUPLICATED KEY UPDATE update_version=VALUES(update_version),user_name=VALUES(user_name),todo_visibility=VALUES(todo_visibility)")
	}
	params := []interface{}{e.UpdateVersion, e.UserId, e.UserName, e.TodoVisibility}
	return dao.db.Exec(ctx, tx, query.String(), params...)
}

func (dao *UserProfileDao) BatchInsert(ctx context.Context, tx *wrap.Tx, list []*UserProfile, onDuplicatedKeyUpdate bool) (result *wrap.Result, err error) {
	query := bytes.NewBufferString("")
	query.WriteString("INSERT INTO user_profile (update_version,user_id,user_name,todo_visibility) VALUES ")
	query.WriteString(wrap.RepeatWithSeparator("(?,?,?,?)", len(list), ","))
	if onDuplicatedKeyUpdate {
		query.WriteString(" ON DUPLICATED KEY UPDATE update_version=VALUES(update_version),user_name=VALUES(user_name),todo_visibility=VALUES(todo_visibility)")
	}
	params := make([]interface{}, len(list)*4)
	offset := 0
	for _, e := range list {
		params[offset+0] = e.UpdateVersion
		params[offset+1] = e.UserId
		params[offset+2] = e.UserName
		params[offset+3] = e.TodoVisibility
		offset += 4
	}

	return dao.db.Exec(ctx, tx, query.String(), params...)
}

func (dao *UserProfileDao) DeleteById(ctx context.Context, tx *wrap.Tx, id uint64) (result *wrap.Result, err error) {
	query := "DELETE FROM UserProfile WHERE id=?"
	return dao.db.Exec(ctx, tx, query, id)
}

func (dao *UserProfileDao) UpdateById(ctx context.Context, tx *wrap.Tx, e *UserProfile) (result *wrap.Result, err error) {
	query := "UPDATE user_profile SET update_version=?,user_id=?,user_name=?,todo_visibility=? WHERE id=?"
	params := []interface{}{e.UpdateVersion, e.UserId, e.UserName, e.TodoVisibility, e.Id}
	return dao.db.Exec(ctx, tx, query, params...)
}

func (dao *UserProfileDao) SelectById(ctx context.Context, tx *wrap.Tx, id int64) (e *UserProfile, err error) {
	query := "SELECT id,create_time,update_time,update_version,user_id,user_name,todo_visibility FROM user_profile WHERE id=?"
	row := dao.db.QueryRow(ctx, tx, query, id)
	e = &UserProfile{}
	err = row.Scan(&e.Id, &e.CreateTime, &e.UpdateTime, &e.UpdateVersion, &e.UserId, &e.UserName, &e.TodoVisibility)
	if err == wrap.ErrNoRows {
		return nil, nil
	}
	return e, err
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
