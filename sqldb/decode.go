package sqldb

import "database/sql"

func decodeValue(values []any) []any {
	realValue := make([]any, len(values))
	for i, v := range values {
		switch v.(type) {
		case *string:
			realValue[i] = *v.(*string)
		case *bool:
			realValue[i] = *v.(*bool)
		case *float32:
			realValue[i] = *v.(*float32)
		case *float64:
			realValue[i] = *v.(*float64)
		case *int:
			realValue[i] = *v.(*int)
		case *int8:
			realValue[i] = *v.(*int8)
		case *int16:
			realValue[i] = *v.(*int16)
		case *int32:
			realValue[i] = *v.(*int32)
		case *int64:
			realValue[i] = *v.(*int64)
		case *uint:
			realValue[i] = *v.(*int64)
		case *uint8:
			realValue[i] = *v.(*uint8)
		case *uint16:
			realValue[i] = *v.(*uint16)
		case *uint32:
			realValue[i] = *v.(*uint32)
		case *uint64:
			realValue[i] = *v.(*uint64)
		case *[]uint8:
			realValue[i] = string(*v.(*[]uint8))
		case *any:
			realValue[i] = *v.(*any)
		case *sql.NullInt64:
			if i64 := *v.(*sql.NullInt64); i64.Valid {
				realValue[i] = i64
			} else {
				realValue[i] = 0
			}
		case *sql.NullByte:
			realValue[i] = *v.(*sql.NullByte)
		case *sql.RawBytes:
			realValue[i] = *v.(*sql.RawBytes)
		default:
			realValue[i] = *v.(*any)
		}
	}
	return realValue
}
func decodeRows(rows *sql.Rows) ([][]D, error) {
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Error("Can't get data rows ColumnTypes: rows.ColumnTypes()")
		return nil, err
	}
	result := make([][]D, 0)
	for rows.Next() {
		data := scanBuffer(columnTypes)
		err := rows.Scan(data...)
		if err != nil {
			log.Fatal("Scan values from sql error: ", err)
		}
		rowValues := decodeValue(data)
		tmp := make([]D, 0)
		for i, v := range columnTypes {
			d := D{
				Name:  v.Name(),
				Type:  v.DatabaseTypeName(),
				Value: rowValues[i],
			}
			tmp = append(tmp, d)
		}
		result = append(result, tmp)
	}
	return result, err
}
