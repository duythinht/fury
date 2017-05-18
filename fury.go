package fury

import (
	"fmt"
	"reflect"

	"github.com/gocql/gocql"
)

func println(msg ...interface{}) {
	fmt.Println(msg...)
}

type Query struct {
	sess   *gocql.Session
	stmt   string
	values []interface{}
}

type Rows struct {
	*gocql.Iter
	cursorIndex int
}

func QueryUsing(sess *gocql.Session) *Query {
	return &Query{sess: sess}
}

func (q *Query) CQL(stmt string, values ...interface{}) *Query {
	q.stmt = stmt
	q.values = values
	return q
}

func (q *Query) Rows() Rows {
	return Rows{q.sess.Query(q.stmt, q.values...).Iter(), 0}
}

func (rows *Rows) Next() bool {
	if rows.cursorIndex < rows.Iter.NumRows() {
		return true
	} else {
		_ = rows.Iter.Close()
		return false
	}
}

func (rows *Rows) Scan(item interface{}) bool {
	rows.cursorIndex++

	_type := reflect.TypeOf(item).Elem()
	val := reflect.ValueOf(item).Elem()

	m := map[string]interface{}{}

	result := rows.Iter.MapScan(m)

	for i := 0; i < _type.NumField(); i++ {
		key := _type.Field(i).Tag.Get("cass")
		if len(key) > 0 {
			val.FieldByName(_type.Field(i).Name).Set(reflect.ValueOf(m[key]))
		}
	}

	return result
}
