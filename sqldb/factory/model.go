package factory

type Filter struct {
	Key    string
	Values any
}

type Field struct {
	Key     string
	Value   any
	Options map[string]string
	Type    string
}

type Model struct {
	Fields     []Field
	Name       string
	OriginName string
}

func NewModel(name string) *Model {
	return &Model{
		Name:   name,
		Fields: make([]Field, 0),
	}
}

func (m *Model) AddField(field Field) {
	m.Fields = append(m.Fields, field)
}
