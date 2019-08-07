// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sync

import (
	"sync"

	"github.com/2637309949/bulrush-utils/maps"
	"github.com/fanliao/go-promise"
)

type (
	// Lock defined rw for some mu resources
	Lock struct {
		locks *maps.SafeMap
	}
	// Async defined promise.all
	Async func(gone ...interface{}) (interface{}, error)
	// Done defined promise done
	Done func()
)

// NewLock defined NewLock
func NewLock() *Lock {
	return &Lock{
		locks: maps.NewSafeMap(),
	}
}

// Acquire defined acquire a mu resources
func (l *Lock) Acquire(name string, funk func(async Async)) {
	defer func() {
		l.locks.Get(name).(*sync.RWMutex).Unlock()
	}()

	if l.locks.Get(name) == nil {
		l.locks.Set(name, new(sync.RWMutex))
	}
	l.locks.Get(name).(*sync.RWMutex).Lock()
	funk(func(gone ...interface{}) (interface{}, error) {
		return promise.WhenAll(gone...).Get()
	})
}

// Count defined count resources
func (l *Lock) Count() int {
	return len(l.locks.ALL())
}
