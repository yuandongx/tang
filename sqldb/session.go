package sqldb

import (
	"database/sql"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
		log.Fatal("mysql连接败：", err)
	} else {
		s.Ping()
	}
	return s
}

func CreatePostgresSession(dataSourceName string) *Session {
	db, err := sql.Open("postgres", dataSourceName)
	s := &Session{_type: "postgres", db: db}
	if err != nil {
		s.status = FAILURE
		log.Fatal("postgre连接败：", err)
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
		log.Error("sql连接败：", err)
		return PingError
	}
}
