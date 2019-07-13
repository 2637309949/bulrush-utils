package utils

import (
	"reflect"
)

// IsIteratee returns if the argument is an iteratee.
func IsIteratee(src interface{}) bool {
	srcType := reflect.TypeOf(src)

	if srcType.Kind() == reflect.Ptr && srcType.Elem().Kind() == reflect.Interface {
		srcType = srcType.Elem().Elem()
	}
	if srcType.Kind() == reflect.Ptr {
		srcType = srcType.Elem()
	}

	kind := srcType.Kind()
	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

// IndirectValue returns the value that v points to.
// If v is a nil pointer, Indirect returns a zero Value.
// If v is not a pointer, Indirect returns v.
func IndirectValue(src reflect.Value) reflect.Value {
	srcValue := src
	if srcValue.Kind() == reflect.Ptr && srcValue.Elem().Kind() == reflect.Interface {
		srcValue = srcValue.Elem().Elem()
	}
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	return srcValue
}
