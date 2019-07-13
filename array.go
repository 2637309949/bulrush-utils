// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"reflect"
)

// Append defined slice append func
func Append(src interface{}, arr interface{}) interface{} {
	if !IsIteratee(arr) {
		panic("Second parameter must be an iteratee")
	}
	var (
		srcValue = IndirectValue(reflect.ValueOf(src))
		arrValue = IndirectValue(reflect.ValueOf(arr))
	)

	if srcValue.Type().Kind() != reflect.Func && arrValue.Type().Elem() != srcValue.Type() {
		panic("First parameter type must be eq Second parameter type")
	}

	array := reflect.New(arrValue.Type())
	if array.Elem().CanSet() {
		array.Elem().Set(reflect.Append(arrValue, srcValue))
	}

	if reflect.TypeOf(arr).Kind() == reflect.Slice {
		array = array.Elem()
	}
	return array.Interface()
}
