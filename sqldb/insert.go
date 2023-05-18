package sqldb

import (
	"fmt"
	"tang/sqldb/factory"
)

// Save 执行插入语句
func (s *Session) Save(object any) (int64, error) {
	model := factory.Parse(object, s._type, s._schema)
	_sql, value := model.Add()
	if result, err := s.Exec(_sql, value); err == nil {
		return result.LastInsertId()
	} else {
		return 0, error_(fmt.Sprintf("insert values into table(%s) failed.", model.GetName()))
	}
}
