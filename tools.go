// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"reflect"

	"github.com/2637309949/bulrush-utils/reflects"
)

// Some defined value with default
func Some(src interface{}, init interface{}) interface{} {
	if !reflects.ISBlank(reflect.ValueOf(src)) {
		return src
	}
	return init
}
