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
