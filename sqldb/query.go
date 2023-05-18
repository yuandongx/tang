package sqldb

import "database/sql"

// Query 执行查询
func (s *Session) Query(query string, args ...any) ([]D, error) {
	if tx, err := s.db.Begin(); err == nil {
		if stm, err := tx.Prepare(query); err == nil {
			defer func(stm *sql.Stmt) {
				_ = stm.Close()
			}(stm)
			rows, err := stm.Query(args...)
			if err != nil {
				return nil, err
			}
			return decodeRows(rows)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
