package sqldb

func (s *Session) Init() {
	sql := "CREATE TABLE IF NOT EXISTS tang_register(id serial, struct_name varchar(128), create_sql text);"
	_, err := s.Exec(sql)
	if err != nil {
		log.Error("数据初始化错误[register]", err)
	} else {
		log.Debug("初始化创建表完成。")
	}
}

func (s *Session) Regist(o any) (ok bool) {
	return
}

func (s *Session) ListTables() (tables []string) {
	return
}
