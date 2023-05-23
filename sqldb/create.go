package sqldb

import (
	"fmt"
	"tang/sqldb/factory"
)

// Create 执行创建表
func (s *Session) Create(object any) (int64, error) {
	model := factory.Parse(object, s._type, s._schema)
	_sql, _ := model.Create()
	if _, err := s.Exec(_sql); err != nil {
		return -1, error_(fmt.Sprintf("create table(%s) failed.", model.GetName()))
	}
	return 0, nil
}
