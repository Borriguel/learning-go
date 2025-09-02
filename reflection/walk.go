package reflection

import "reflect"

func walk(x interface{}, fn func(entry string)) {
	value := getValue(x)
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), fn)
		}
	case reflect.String:
		fn(value.String())
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
