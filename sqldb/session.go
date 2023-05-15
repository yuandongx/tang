package sqldb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"tang/sqldb/factory"
)

const (
	OK = iota
	FAILURE
	PingError
)

type Session struct {
	db      *sql.DB
	_schema string
	_type   string // mysql/postgresql/sqlite
	status  int
}

func (s *Session) Exec(query string, args []any) (sql.Result, error) {
	if tx, err := s.db.Begin(); err == nil {
		if stm, err := tx.Prepare(query); err == nil {
			defer func(stm *sql.Stmt) {
				_ = stm.Close()
			}(stm)
			return stm.Exec(args...)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

}

func (s *Session) Query(query string, args ...any) ([]D, error) {
	if tx, err := s.db.Begin(); err == nil {
		if stm, err := tx.Prepare(query); err == nil {
			defer func(stm *sql.Stmt) {
				_ = stm.Close()
			}(stm)
			rows, err := stm.Query(args...)
			if err != nil {
				return nil, err
			}
			return decodeRows(rows)
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

func CreateNySqlSession(dataSourceName string) *Session {
	db, err := sql.Open("mysql", dataSourceName)
	s := &Session{_type: "mysql", db: db}
	if err != nil {
		s.status = FAILURE
		logger.Fatal("mysql连接败：", err)
	} else {
		s.Ping()
	}
	return s
}

func (s *Session) Ping() int {
	err := s.db.Ping()
	if err == nil {
		s.status = OK
		return OK
	} else {
		s.status = PingError
		logger.Println("mysql连接败：", err)
		return PingError
	}
}
