package sqldb

import (
	"fmt"
	"tang/sqldb/factory"
)

func (s *Session) Init(args ...any) {
	name := s.name_space
	if len(args) > 0 {
		if v, ok := args[0].(string); ok {
			name = v
		}
	}
	create_sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s_register" (
		"id" serial NOT NULL,
		PRIMARY KEY ("id"),
		"name" character varying(128) NOT NULL,
		"origin_name" character varying(128) NOT NULL,
		"create_sql" text NOT NULL,
		"update_time" timestamp NOT NULL,
		"update_sql" text NULL
	  );
	  COMMENT ON COLUMN "%s"."id" IS 'primary key of entry';`, name, name)
	s.Exec(create_sql)
}
func (s *Session) Register(object any) {
	model := factory.Parse(object, s._type, s._schema)
	qurey := fmt.Sprintf("SELECT * %s_register WHERE name=%s;", s.name_space, model.GetName())
	s.Query(qurey)
}
