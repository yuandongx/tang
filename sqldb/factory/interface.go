package factory

type SqlInterface interface {
	Create() (string, []any)

	Add() (string, []any)
	AddMany() (string, []any)

	Delete() (string, []any)
	DeleteMany([]Filter) (string, []any)

	Update() (string, []any)
	UpdateMay() (string, []any)

	GetById(id string) (string, []any)
	GetAll() (string, []any)
	GetOne([]Filter) (string, []any)

	GetName() string
}
