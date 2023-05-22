package factory

import (
	"fmt"
	"strings"
)

type MysqlModel struct {
	Name   string
	Schema string
	Model
}

func (m MysqlModel) Create() (string, []any) {
	fs := make([]string, 0)
	hasKey := false
	for _, f := range m.Fields {
		s := f.Key
		if _type, ok := f.Options["type"]; ok {
			s = s + " " + _type
		} else {
			s = s + " text"
		}

		if null, ok := f.Options["null"]; ok && null == "true" {
			s = s + " NULL"
		} else {
			s = s + " NOT NULL"
		}

		if auto, ok := f.Options["auto_increment"]; ok && auto == "true" {
			s = s + " AUTO_INCREMENT PRIMARY KEY"
			hasKey = true
		}

		if _default, ok := f.Options["default"]; ok {
			s = s + " DEFAULT '" + _default + "' "
		}

		if unique, ok := f.Options["unique"]; ok && unique == "true" {
			s = s + " UNIQUE"
		}

		if comment, ok := f.Options["comment"]; ok {
			s = s + " COMMENT '" + comment + "'"
		}

		fs = append(fs, s)
	}
	sqlString := strings.Join(fs, ", ")
	if !hasKey {
		sqlString = "id INT AUTO_INCREMENT PRIMARY KEY, " + sqlString
	}
	sql := "CREATE TABLE `" + m.Name + "` (" + sqlString + ") ENGINE='InnoDB';"
	return sql, nil
}

func (m MysqlModel) Add() (string, []any) {
	fields := make([]string, len(m.Fields))
	values := make([]any, len(m.Fields))
	fills := make([]string, len(m.Fields))
	for i, v := range m.Fields {
		values[i] = v.Value
		fields[i] = v.Key
		fills[i] = "?"
	}
	sql := "INSERT INTO " + m.Name + " (" + strings.Join(fields, ",") + ")values(" + strings.Join(fills, ", ") + ")"
	return sql, values
}

func (m MysqlModel) Delete() (string, []any) {
	fields := make([]string, len(m.Fields))
	values := make([]any, len(m.Fields))
	fills := make([]string, len(m.Fields))
	for i, v := range m.Fields {
		values[i] = v.Value
		fields[i] = v.Key + "=?"
		fills[i] = "?"
	}
	sql := "DELETE FROM " + m.Name + " where " + strings.Join(fields, " AND ") + "; "
	return sql, values
}

func (m MysqlModel) DeleteMore(filter map[string]any) (string, []any) {
	if len(filter) == 0 {
		return "", nil
	}
	fields := make([]string, 0)
	values := make([]any, 0)
	for key, value := range filter {
		fields = append(fields, key+"=?")
		values = append(values, value)
	}
	sql := "DELETE FROM " + m.Name + " where " + strings.Join(fields, " AND ") + ";"
	return sql, values
}

// Update 更新对象， filters：过滤条件， keys：指定过滤哪些字段，为空的时候则全部字段
func (m MysqlModel) Update(filters map[string]any, keys []string) (string, []any) {
	fields := make([]string, 0)
	values := make([]any, 0)
	filter := make([]string, len(m.Fields))
	for _, v := range m.Fields {
		if keys != nil {
			for _, k := range keys {
				if k == v.Key {
					values = append(values, v.Value)
					fields = append(fields, v.Key+"=?")
				}
			}
		} else {
			values = append(values, v.Value)
			fields = append(fields, v.Key+"=?")
		}
	}
	for key, item := range filters {
		filter = append(filter, key+"=?")
		values = append(values, item)
	}
	sql := "UPDATE " + m.Name + " SET " + strings.Join(fields, ", ")
	sql += " WHERE " + strings.Join(filter, " AND ") + ";"
	return sql, values
}

func (m MysqlModel) GetById(id string) (string, []any) {
	values := []any{id}
	return selectId(m.Name, "id"), values
}

func (m MysqlModel) GetByKey(key string) (string, []any) {
	values := []any{key}
	return selectId(m.Name, "key"), values
}

func (m MysqlModel) GetAll() (string, []any) {
	return fmt.Sprintf("SELECT * FROM %s;", m.Name), nil
}
func (m MysqlModel) GetMore(filters map[string]any) (string, []any) {
	keys := make([]string, 0)
	values := make([]any, 0)
	for key, value := range filters {
		keys = append(keys, key+"=?")
		values = append(values, value)
	}
	sql := "SELECT * FROM " + m.Name + " WHERE " + strings.Join(keys, " AND ") + ";"
	return sql, values
}
func (m MysqlModel) GetName() string {
	return m.Name
}

func selectId(name string, tag string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE %s", name, tag)
}
