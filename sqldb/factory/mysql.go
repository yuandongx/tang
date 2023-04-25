package factory

type MysqlModel struct {
	Name   string
	Schema string
	Model
}

func (m MysqlModel) Create() (string, []any) {
	return "", nil
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
