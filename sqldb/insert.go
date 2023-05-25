package sqldb

import (
	"fmt"
	"tang/sqldb/factory"
)

// Save 执行插入语句
func (s *Session) Save(object any) (int64, error) {
	model := factory.Parse(object, s._type, s._schema)
	_sql, value := model.Add()
	log.Debug(_sql, value)
	if _, err := s.Exec(_sql, value...); err == nil {
		return 1, nil
	} else {
		msg := fmt.Sprintf("insert values into table(%s) failed.", model.GetName())
		log.Error(msg, err)
		return 0, error_(msg)
	}
}
