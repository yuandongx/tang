package sqldb

import (
	"fmt"
	"tang/sqldb/factory"
)

func (s *Session) Init(args ...any) {
	// name := s.name_space
	// if len(args) > 0 {
	// 	if v, ok := args[0].(string); ok {
	// 		name = v
	// 	}
	// }
	create_sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (
		"id" serial NOT NULL,
		PRIMARY KEY ("id"),
		"name" character varying(128) NOT NULL,
		"origin_name" character varying(128) NOT NULL,
		"create_sql" text NOT NULL,
		"update_time" timestamp NOT NULL,
		"update_sql" text NULL
	  );`, s.GetRegisterName())
	res, err := s.Exec(create_sql)
	fmt.Println(res)
	fmt.Println(err)
}
func (s *Session) Register(object any) {
	model := factory.Parse(object, s._type, s._schema)
	qurey := fmt.Sprintf("SELECT * %s WHERE name=%s;", s.GetRegisterName(), model.GetName())
	data, err := s.Query(qurey)
	fmt.Println(111111111, data)
	fmt.Println(err)
	if len(data) == 0 {
		sqlstring, _ := model.Create()
		s.Exec(fmt.Sprintf("INSERT INTO %s(name, origin_name, create_sql, update_time)VALUES($1,$2,$3, $4)",
			s.GetRegisterName()),
			model.GetName(),
			model.GetOriginName(),
			sqlstring,
			"NOW()")
	}
}
