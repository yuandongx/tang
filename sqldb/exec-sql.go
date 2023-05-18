package sqldb

import "database/sql"

// Exec 执行SQL语句
func (s *Session) Exec(query string, args ...any) (sql.Result, error) {
	if tx, err := s.db.Begin(); err == nil {
		if stm, err := tx.Prepare(query); err == nil {
			defer func(stm *sql.Stmt) {
				_ = stm.Close()
			}(stm)
			return stm.Exec(args...)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
