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
type _error struct {
	e string
}

func (e _error) Error() string {
	return e.e
}
func error_(msg string) *_error {
	return &_error{e: msg}
}

func New() *Session {
	return &Session{}
}

func (s *Session) Exec(query string, args []any) (sql.Result, error) {
	if tx, err := s.db.Begin(); err == nil {
		if stm, err := tx.Prepare(query); err == nil {
			defer stm.Close()
			return stm.Exec(args...)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

}

func (s *Session) Query(query string, args ...any) ([][]DV, error) {
	if tx, err := s.db.Begin(); err == nil {
		if stm, err := tx.Prepare(query); err == nil {
			defer stm.Close()
			rows, err := stm.Query(args...)
			return decodeRows(rows), err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (s *Session) Save(object any) (int64, error) {
	model := factory.Parse(object, s._type, s._schema)
	_sql, value := model.Add()
	if result, err := s.Exec(_sql, value); err == nil {
		return result.LastInsertId()
	} else {
		return 0, error_("save failed.")
	}
}

func decodeRows(rows *sql.Rows) (data [][]DV) {
	data = make([][]DV, 0)
	rows.ColumnTypes()
	for rows.Next() {
		rows.Scan()
	}
	return
}
