package config
import (
	"strings"
)
const (
	APP_PREFIX = "TANG_"
	PREFIX = "tang_"
)

func GetEnv(key string) string {
	key = strings.Trim()
	key = strings.ToUpper(key)
	if !strings.HasPrefix(key, APP_PREFIX) {
		key = APP_PREFIX+key
	}
	return getEnv(key)
}


func getEnv(key, _default string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return _default
}