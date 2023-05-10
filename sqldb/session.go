package sqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"tang/sqldb/factory"
	"time"
)

const (
	OK = iota
	FAILURE
	PING_ERROR
)

type Session struct {
	db      *sql.DB
	_schema string
	_type   string // mysql/postgresql/sqlite
	status  int
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

func scanBuffer(columnTypes []*sql.ColumnType) []any {
	buffer := make([]any, len(columnTypes))
	for i, t := range columnTypes {
		fmt.Println(t.ScanType().Kind())
		fmt.Println(t.ScanType().Kind().String())
		switch t.ScanType().Kind() {
		case reflect.Int:
			buffer[i] = new(int)
		case reflect.Int8:
			buffer[i] = new(int8)
		case reflect.Int16:
			buffer[i] = new(int16)
		case reflect.Int32:
			buffer[i] = new(int32)
		case reflect.Int64:
			buffer[i] = new(int64)
		case reflect.Uint:
			buffer[i] = new(uint)
		case reflect.Uint8:
			buffer[i] = new(uint8)
		case reflect.Uint16:
			buffer[i] = new(uint16)
		case reflect.Uint32:
			buffer[i] = new(uint32)
		case reflect.Uint64:
			buffer[i] = new(uint64)
		case reflect.String:
			buffer[i] = new(string)
		case reflect.Bool:
			buffer[i] = new(bool)
		case reflect.Float32:
			buffer[i] = new(float32)
		case reflect.Float64:
			buffer[i] = new(float64)
		case reflect.Interface:
			buffer[i] = new(interface{})
		case reflect.Struct:
			switch t.ScanType().String() {
			case "sql.NullTime":
				fmt.Println(111111, t.ScanType().String(), t.ScanType().String() == "sql.NullTime")
				buffer[i] = &sql.NullTime{}
			case "sql.RawBytes":
				fmt.Println(111111, t.ScanType().String(), t.ScanType().String() == "sql.NullTime")
				buffer[i] = &sql.RawBytes{}
			case "sql.NullInt64":
				fmt.Println(111111, t.ScanType().String(), t.ScanType().String() == "sql.NullTime")
				buffer[i] = &sql.NullInt64{}
			default:
				buffer[i] = new(interface{})
			}
		default:
			buffer[i] = new(interface{})
		}
	}
	return buffer
}

func decodeValue(values []any) []any {
	realValue := make([]any, len(values))
	for i, v := range values {
		switch v.(type) {
		case *string:
			realValue[i] = *v.(*string)
		case *bool:
			realValue[i] = *v.(*bool)
		case *float32:
			realValue[i] = *v.(*float32)
		case *float64:
			realValue[i] = *v.(*float64)
		case *int:
			realValue[i] = *v.(*int)
		case *int8:
			realValue[i] = *v.(*int8)
		case *int16:
			realValue[i] = *v.(*int16)
		case *int32:
			realValue[i] = *v.(*int32)
		case *int64:
			realValue[i] = *v.(*int64)
		case *uint:
			realValue[i] = *v.(*int64)
		case *uint8:
			realValue[i] = *v.(*uint8)
		case *uint16:
			realValue[i] = *v.(*uint16)
		case *uint32:
			realValue[i] = *v.(*uint32)
		case *uint64:
			realValue[i] = *v.(*uint64)
		case *any:
			realValue[i] = *v.(*any)
		case *sql.NullTime:
			if t := *v.(*sql.NullTime); t.Valid {
				realValue[i] = t.Time
			} else {
				realValue[i] = time.Time{}
			}
		case *sql.NullInt64:
			if i64 := *v.(*sql.NullInt64); i64.Valid {
				realValue[i] = i64
			} else {
				realValue[i] = 0
			}
		case *sql.NullByte:
			realValue[i] = *v.(*sql.NullByte)
		case *sql.RawBytes:
			realValue[i] = *v.(*sql.RawBytes)
		default:
			realValue[i] = *v.(*any)
		}
	}
	return realValue
}
func decodeRows(rows *sql.Rows) (data [][]DV) {
	data = make([][]DV, 0)
	columnTypes, err := rows.ColumnTypes()

	fmt.Println(err)
	for i, v := range columnTypes {
		a, b := v.Nullable()
		fmt.Println(i, v.Name(), a, b, v.ScanType(), v.DatabaseTypeName())
	}
	for rows.Next() {
		data := scanBuffer(columnTypes)
		rows.Scan(data...)
		vs := decodeValue(data)
		fmt.Println(vs)
	}
	return
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
		logger.Fatal("mysql连接败：", err)
		s.status = PING_ERROR
		return PING_ERROR
	}
}
