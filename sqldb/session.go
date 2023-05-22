package sqldb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

func CreateMySqlSession(dataSourceName string) *Session {
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

func CreatePostgresSession(dataSourceName string) *Session {
	db, err := sql.Open("postgre", dataSourceName)
	s := &Session{_type: "postgre", db: db}
	if err != nil {
		s.status = FAILURE
		logger.Fatal("postgre连接败：", err)
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
		logger.Println("sql连接败：", err)
		return PingError
	}
}
