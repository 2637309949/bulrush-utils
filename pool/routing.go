// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pool

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// RoutingPoolWithTimer defined routing pool for work
// Example
// _, done := RoutingPoolWithTimer(func(cancel context.CancelFunc) {
// 	time.Sleep(10 * time.Second)
// }, 10)
// <-done
func RoutingPoolWithTimer(worker func(context.CancelFunc), max int64) (context.CancelFunc, chan struct{}) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	limit := max
	curr := int64(0)
	timer := time.NewTicker(1 * time.Second)
	done := make(chan struct{}, 0)
	init := false
	var mutex sync.Mutex
	runWork := func(curr *int64, limit *int64) {
		if *curr < *limit {
			func() {
				defer func() {
					if *curr > 0 {
						atomic.AddInt64(curr, -1)
					}
				}()
				atomic.AddInt64(curr, 1)
				worker(cancel)
			}()
		}
	}
	go func() {
		for {
			if !init {
				init = true
				for i := int64(0); i < limit; i++ {
					go runWork(&curr, &limit)
				}
			}
			select {
			case <-ctx.Done():
				limit = 0
				done <- struct{}{}
				break
			case <-timer.C:
				mutex.Lock()
				for i := limit - curr; i > 0; i-- {
					go runWork(&curr, &limit)
				}
				mutex.Unlock()
			}
		}
	}()
	return cancel, done
}

// RoutingPoolWithAutomatic defined routing pool for work
// Example
// i := 0
// _, done := pool.RoutingPoolWithAutomatic(func(cancel context.CancelFunc) {
// 	time.Sleep(10 * time.Second)
// 	i++
// 	if i == 30 {
// 		cancel()
// 	}
// }, 10)
// <-done
func RoutingPoolWithAutomatic(worker func(context.CancelFunc), max int64) (context.CancelFunc, chan struct{}) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	limit := max
	done := make(chan struct{}, 0)
	init := false
	var runWork func(limit *int64)
	runWork = func(limit *int64) {
		func() {
			defer func() {
				if *limit > 0 {
					go runWork(limit)
				}
			}()
			worker(cancel)
		}()
	}
	go func() {
		for {
			if !init {
				init = true
				for i := int64(0); i < limit; i++ {
					go runWork(&limit)
				}
			}
			select {
			case <-ctx.Done():
				limit = 0
				done <- struct{}{}
				break
			default:
			}
		}
	}()
	return cancel, done
}