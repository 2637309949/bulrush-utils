// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"reflect"
)

// Some defined value with default
func Some(src interface{}, init interface{}) interface{} {
	if !ISBlank(reflect.ValueOf(src)) {
		return src
	}
	return init
}

// Until defined flow not blank
func Until(flow ...interface{}) interface{} {
	for index := 0; index < len(flow); index++ {
		fw := flow[index]
		if reflect.TypeOf(fw).Kind() == reflect.Func {
			ret := reflect.ValueOf(fw).Call([]reflect.Value{})[0]
			if !ISBlank(ret) {
				return ret.Interface()
			}
		} else {
			if !ISBlank(reflect.ValueOf(fw)) {
				return fw
			}
		}
	}
	return nil
}
