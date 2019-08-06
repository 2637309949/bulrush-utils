// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sync

import (
	"sync"

	"github.com/2637309949/bulrush-utils/maps"
)

// Lock defined rw for some mu resources
type Lock struct {
	locks *maps.SafeMap
}

// NewLock defined NewLock
func NewLock() *Lock {
	return &Lock{
		locks: maps.NewSafeMap(),
	}
}

// AcquireForSync defined acquire a mu resources
func (l *Lock) AcquireForSync(name string, funk func()) {
	defer func() {
		l.locks.Get(name).(*sync.RWMutex).Unlock()
	}()
	lock := l.locks.Get(name)
	if lock == nil {
		l.locks.Set(name, new(sync.RWMutex))
	}
	l.locks.Get(name).(*sync.RWMutex).Lock()
	funk()
}

// AcquireForAsync defined acquire a mu resources
func (l *Lock) AcquireForAsync(name string, funk func(done func())) {
	lock := l.locks.Get(name)
	if lock == nil {
		l.locks.Set(name, new(sync.RWMutex))
	}
	l.locks.Get(name).(*sync.RWMutex).Lock()
	funk(func() {
		l.locks.Get(name).(*sync.RWMutex).Unlock()
	})
}

// Count defined count resources
func (l *Lock) Count() int {
	return len(l.locks.ALL())
}
