package factory

import (
	"fmt"
	"strings"
)

type PostgresModel struct {
	Name string
	Model
}

func (m PostgresModel) Create() (string, []any) {
	fs := make([]string, 0)
	hasKey := ""
	for _, f := range m.Fields {
		s := lower(f.Key)
		if _type, ok := f.Options["type"]; ok {
			s = s + " " + _type
			fmt.Println(s)
		} else {
			fmt.Println(s, f.Type)
			s = s + " " + PgType(f.Type)
		}

		if null, ok := f.Options["null"]; ok && null == "true" {
			s = s + " NULL"
		} else {
			s = s + " NOT NULL"
		}

		if auto, ok := f.Options["auto_increment"]; ok && auto == "true" {
			s = s + " serial"
			hasKey = f.Key
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
	if hasKey == "" {
		sqlString = sqlString + ", id serial, PRIMARY KEY(id)"
	} else {
		sqlString += ", PRIMARY KEY(" + hasKey + ")"
	}
	sql := "CREATE TABLE " + m.Name + "(" + sqlString + ");"
	return sql, nil
}

func (m PostgresModel) Add() (string, []any) {
	fields := make([]string, len(m.Fields))
	values := make([]any, len(m.Fields))
	fills := make([]string, len(m.Fields))
	for i, v := range m.Fields {
		values[i] = v.Value
		fields[i] = v.Key
		fills[i] = fmt.Sprintf("$%d", i+1)
	}
	sql := "INSERT INTO " + m.Name + "(" + strings.Join(fields, ",") + ")values(" + strings.Join(fills, ", ") + ")"
	return sql, values
}

func (m PostgresModel) Delete() (string, []any) {
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

func (m PostgresModel) DeleteMore(filter map[string]any) (string, []any) {
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
func (m PostgresModel) Update(filters map[string]any, keys []string) (string, []any) {
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

func (m PostgresModel) GetById(id string) (string, []any) {
	values := []any{id}
	return selectId(m.Name, "id"), values
}

func (m PostgresModel) GetByKey(key string) (string, []any) {
	values := []any{key}
	return selectId(m.Name, "key"), values
}

func (m PostgresModel) GetAll() (string, []any) {
	return fmt.Sprintf("SELECT * FROM %s;", m.Name), nil
}
func (m PostgresModel) GetMore(filters map[string]any) (string, []any) {
	keys := make([]string, 0)
	values := make([]any, 0)
	for key, value := range filters {
		keys = append(keys, key+"=?")
		values = append(values, value)
	}
	sql := "SELECT * FROM " + m.Name + " WHERE " + strings.Join(keys, " AND ") + ";"
	return sql, values
}
func (m PostgresModel) GetName() string {
	return m.Name
}
