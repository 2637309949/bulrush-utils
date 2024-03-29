// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package array

import (
	"reflect"

	"github.com/2637309949/bulrush-utils/reflects"
)

// Append defined slice append func
func Append(src interface{}, arr interface{}) interface{} {
	if !reflects.IsIteratee(arr) {
		panic("Second parameter must be an iteratee")
	}
	var (
		srcValue = reflects.IndirectValue(reflect.ValueOf(src))
		arrValue = reflects.IndirectValue(reflect.ValueOf(arr))
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
