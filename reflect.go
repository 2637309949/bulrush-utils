// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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

// ISBlank defined check value is blank or not
func ISBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// MakeSlice defined MakeSlice
func MakeSlice(elemType reflect.Type) interface{} {
	if elemType.Kind() == reflect.Slice {
		elemType = elemType.Elem()
	}
	sliceType := reflect.SliceOf(elemType)
	slice := reflect.New(sliceType)
	slice.Elem().Set(reflect.MakeSlice(sliceType, 0, 0))
	return slice.Interface()
}
