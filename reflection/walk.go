package reflection

import "reflect"

func walk(x interface{}, fn func(entry string)) {
	value := getValue(x)
	quantityOfValues := 0
	var getField func(int) reflect.Value
	switch value.Kind() {
	case reflect.Struct:
		quantityOfValues = value.NumField()
		getField = value.Field
	case reflect.String:
		fn(value.String())
	case reflect.Slice, reflect.Array:
		quantityOfValues = value.Len()
		getField = value.Index
	}
	for i := 0; i < quantityOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
