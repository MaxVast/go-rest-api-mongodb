package service

import (
	"database/sql"
	"reflect"
)

func ScanStruct(dest interface{}, rows *sql.Rows) error {
	v := reflect.ValueOf(dest).Elem()
	t := v.Type()

	values := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		values[i] = v.Field(i).Addr().Interface()
	}

	if err := rows.Scan(values...); err != nil {
		return err
	}

	return nil
}
