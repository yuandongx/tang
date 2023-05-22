package factory

import "strings"

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
	return "", nil
}
func (m MysqlModel) AddMany() (string, []any) {
	return "", nil
}

func (m MysqlModel) Delete() (string, []any) {
	return "", nil
}
func (m MysqlModel) DeleteMany([]Filter) (string, []any) {
	return "", nil
}

func (m MysqlModel) Update() (string, []any) {
	return "", nil
}
func (m MysqlModel) UpdateMay() (string, []any) {
	return "", nil
}

func (m MysqlModel) GetById(id string) (string, []any) {
	return "", nil
}
func (m MysqlModel) GetAll() (string, []any) {
	return "", nil
}
func (m MysqlModel) GetOne([]Filter) (string, []any) {
	return "", nil
}
func (m MysqlModel) GetName() string {
	return m.Name
}
