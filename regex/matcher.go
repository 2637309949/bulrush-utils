// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package regex

import "regexp"

// FindStringSubmatch defined regex match sub string
func FindStringSubmatch(matcher string, s string) []string {
	var rgx = regexp.MustCompile(matcher)
	rs := rgx.FindStringSubmatch(s)
	if rs != nil {
		return rs[1:]
	}
	return []string{}
}
