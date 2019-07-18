// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import "sync"

// SafeMap defined SafeMap
type SafeMap struct {
	m map[string]string
	l *sync.RWMutex
}

// Set defined Set
func (s *SafeMap) Set(key string, value string) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[key] = value
}

// Get defined Get
func (s *SafeMap) Get(key string) string {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.m[key]
}

// NewSafeMap defined SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{l: new(sync.RWMutex), m: make(map[string]string)}
}
