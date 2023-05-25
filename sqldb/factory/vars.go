package factory

import (
	"strings"
	"tang/logger"
)

// go到数据库的类型映射
var (
	MysqlDataType    = make(map[string]string)
	PostgresDataType = make(map[string]string)
)
var log = logger.GetLogger("sqldb-factory", logger.DEBUG)

const (
	MYSQL    = "mysql"
	POSTGRES = "postgres"
)

func init() {
	// postgres data type
	PostgresDataType["int"] = "integer"
	PostgresDataType["int8"] = "smallint"
	PostgresDataType["int16"] = "smallint"
	PostgresDataType["int32"] = "integer"
	PostgresDataType["int64"] = "bigint"
	PostgresDataType["float32"] = "real"
	PostgresDataType["float64"] = "double precision"
	PostgresDataType["bool"] = "boolean"
	PostgresDataType["string"] = "text"
	PostgresDataType["time"] = "character varying(25)"

}

func PgType(t string) string {
	t = strings.ToLower(t)
	v, ok := PostgresDataType[t]
	if ok {
		return v
	}
	return "text"
}
