package sqldb
import (
	"os",
	"fmt"
)
type _error struct {
	e string
}

func (e _error) Error() string {
	return e.e
}
func error_(msg string) *_error {
	return &_error{e: msg}
}


func GetDbSource()string {
	host = getEnv("DB_HOST", "127.0.0.1")
	name = getEnv("DB_NAME", "mysql")
	username = getEnv("DB_USERNAME", "mysql")
	password = getEnv("DB_PASSWORD", "P@ssw0rd")
	port = getEnv("DB_PORT", "3306")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	username, password, host, port, name)
}

func getEnv(key, _default string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return _default
}