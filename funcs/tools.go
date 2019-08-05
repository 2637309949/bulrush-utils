// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package funcs

import (
	"reflect"

	"github.com/2637309949/bulrush-utils/reflects"
)

// Until defined flow not blank
func Until(flow ...interface{}) interface{} {
	for index := 0; index < len(flow); index++ {
		fw := flow[index]
		if reflect.TypeOf(fw).Kind() == reflect.Func {
			ret := reflect.ValueOf(fw).Call([]reflect.Value{})[0]
			if !reflects.ISBlank(ret) {
				return ret.Interface()
			}
		} else {
			if !reflects.ISBlank(reflect.ValueOf(fw)) {
				return fw
			}
		}
	}
	return nil
}

// Chain defined flow (ret, err) chain
func Chain(flow ...func(interface{}) (interface{}, error)) (interface{}, error) {
	var result interface{}
	for index := 0; index < len(flow); index++ {
		fw := flow[index]
		var resultValue reflect.Value
		if result == nil {
			resultValue = reflect.Zero(reflect.TypeOf(0))
		} else {
			resultValue = reflect.ValueOf(result)
		}
		ret := reflect.ValueOf(fw).Call([]reflect.Value{resultValue})
		result = ret[0].Interface()
		err := ret[1].Interface()
		if err != nil {
			return nil, err.(error)
		}
	}
	return result, nil
}
