package sqldb

import (
	"fmt"
	"tang/sqldb/factory"
)

// Create 执行创建表
func (s *Session) Create(object any) (int64, error) {
	model := factory.Parse(object, s._type, s._schema)
	_sql, _ := model.Create()
	fmt.Println("---->", _sql, s.status)
	if res, err := s.Exec(_sql); err != nil {
		fmt.Println(11111111111, res)
		return -1, error_(fmt.Sprintf("create table(%s) failed.", model.GetName()))
	} else {

		fmt.Println(2222222222, res)
	}

	return 0, nil
}
