package utils

import (
	"reflect"
)

// Append defined array append
func Append(src interface{}, arr interface{}) interface{} {
	if !IsIteratee(arr) {
		panic("Second parameter must be an iteratee")
	}

	var (
		arrValue = reflect.ValueOf(arr)
		srcValue = reflect.ValueOf(src)
	)
	srcValue = IndirectValue(srcValue)
	arrValue = IndirectValue(arrValue)

	if arrValue.Type().Elem() != srcValue.Type() {
		panic("First parameter type must be eq Second parameter type")
	}

	newArr := reflect.New(reflect.SliceOf(srcValue.Type()))
	if newArr.Elem().CanSet() {
		newArr.Elem().Set(reflect.Append(arrValue, srcValue))
	}

	// return type as arr type
	if reflect.TypeOf(arr).Kind() == reflect.Slice {
		newArr = newArr.Elem()
	}
	ret := newArr.Interface()
	return ret
}
