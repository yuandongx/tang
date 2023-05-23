package sqldb

import (
	"database/sql"
	"reflect"
)

func scanBuffer(columnTypes []*sql.ColumnType) []any {
	buffer := make([]any, len(columnTypes))
	for i, t := range columnTypes {
		log.Debug(t.Name(), t.ScanType().Kind(), t.DatabaseTypeName())
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
		case reflect.Slice:
			buffer[i] = &[]uint8{}
		case reflect.Struct:
			switch t.ScanType().String() {
			case "sql.NullTime":
				buffer[i] = &[]uint8{}
			case "sql.RawBytes":
				buffer[i] = &[]uint8{}
			case "sql.NullInt64":
				buffer[i] = new(int64)
			default:
				buffer[i] = new(interface{})
			}
		default:
			buffer[i] = new(interface{})
		}
	}
	return buffer
}
