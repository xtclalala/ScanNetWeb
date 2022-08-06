package tools

import "reflect"

type EmptyType interface {
	[]any | chan any | map[any]any | string | int
}

func IsEmpty[T EmptyType](value T) bool {
	val := reflect.ValueOf(value)

	switch valType := val.Kind(); valType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Int() == 0
	default:
		length := val.Len()
		return length == 0
	}
}

func IsNil(value any) bool {
	return reflect.ValueOf(value).IsNil()
}

func IsAllNil(values ...any) bool {
	for _, value := range values {
		if !IsNil(value) {
			return false
		}
	}
	return true
}
