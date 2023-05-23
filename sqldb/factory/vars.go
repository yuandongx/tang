package factory

// go到数据库的类型映射
var (
	MysqlDataType    = make(map[string]string)
	PostgresDataType = make(map[string]string)
)

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
	PostgresDataType["time"] = "timestamp"

}

func PgType(t string) string {
	v, ok := PostgresDataType[t]
	if ok {
		return v
	}
	return "text"
}
