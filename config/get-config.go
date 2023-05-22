package config
import (
	"strings"
)
const (
	APP_PREFIX = "TANG_"
	PREFIX = "tang_"
)

var (
	db_name = "msyql"
	db_password = "P@ssw0rd"
	db_host = "106.75.63.248"
	db_port = 3306
	db_username = "mysql"
)

func GetEnv(key string) string {
	strings.Stsrtwit()
	return getEnv(key)
}


func getEnv(key, _default string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return _default
}