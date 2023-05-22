package factory

type SqlInterface interface {
	Create() (string, []any)

	Add() (string, []any)

	Delete() (string, []any)
	DeleteMore(map[string]any) (string, []any)

	Update(map[string]any, []string) (string, []any)

	GetById(id string) (string, []any)
	GetAll() (string, []any)
	GetMore(map[string]any) (string, []any)

	GetName() string
}
