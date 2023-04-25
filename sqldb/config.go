package sqldb

type Config struct {
	Host     string
	Port     int
	DB       string
	UserName string
	Password string
	Options  string
	Schema   string
	_type    string // mysql/postgresql/sqlite
}
