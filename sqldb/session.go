package sqldb

import (
	"database/sql"
	"tang/sqldb/factory"
)

type Session struct {
	db      *sql.DB
	_schema string
	_type   string // mysql/postgresql/sqlite
}

func New() *Session {
	return &Session{}
}

func (s *Session) Exec(query string, args []any) {
	tx, err := s.db.Begin()
}

func (s *Session) Query(query string, args []any) {}

func (s Session) Save(object any) {
	model := factory.Parse(object, s._type, s.Schema)
	_sql, value := model.Add()
	s.Exec(_sql, value)
}
