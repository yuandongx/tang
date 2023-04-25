package factory

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 解码struct
func getValue(value reflect.Value) any {
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Bool:
		return value.Bool()
	case reflect.Float32, reflect.Float64:
		return value.Float()
	case reflect.Int, reflect.Int32, reflect.Int8, reflect.Int16, reflect.Int64:
		return value.Int()
	default:
		return value
	}
}
func Parse(object any, _type, schema string) SqlInterface {
	m := decode(object)
	switch _type {
	case "mysql":
		return MysqlModel{Name: "mysql", Model: *m, Schema: schema}
	default:
		return MysqlModel{Name: "mysql", Model: *m, Schema: schema}
	}
}

func decode(object any) (m *Model) {
	_type := reflect.TypeOf(object)
	_value := reflect.ValueOf(object)
	if _type.Kind() != reflect.Struct {
		fmt.Println("目标解码对象不是一个struct类型的变量！")
		return
	}
	m = NewModel(_type.Name())
	for i := 0; i < _type.NumField(); i++ {
		fmt.Println(_type.Field(i))
		fmt.Println(_value.Field(i))
		tfield := _type.Field(i)
		vfield := _value.Field(i)
		f := Field{
			Key:     tfield.Name,
			Value:   getValue(vfield),
			Options: tagToMap(string(tfield.Tag)),
		}
		m.AddField(f)
	}
	return
}

func tagToMap(tag string) (result map[string]string) {
	result = make(map[string]string)
	for tag != "" {
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := tag[:i]
		tag = tag[i+1:]
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := tag[:i+1]
		tag = tag[i+1:]
		value, err := strconv.Unquote(qvalue)
		if err == nil {
			result[name] = value
		} else {
			break
		}

	}
	return
}

func Captain(a string) (b string) {
	tmp := strings.Split(a, "-")
	for _, item := range tmp {
		tmp := make([]rune, len(item))
		for i, c := range item {
			if i == 0 && c >= 'a' && c <= 'z' {
				tmp[0] = c - 32
			} else {
				tmp[i] = c
			}
		}
		b += string(tmp)
	}
	return
}
