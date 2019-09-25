// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pool

import (
	"context"
	"sync/atomic"
	"time"
)

// RoutingPoolWithCancel defined routing pool for work
func RoutingPoolWithCancel(sync func(context.CancelFunc), max int64) (context.CancelFunc, chan struct{}) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	currMax := max
	curr := int64(0)
	timer := time.NewTicker(1 * time.Second)
	done := make(chan struct{}, 0)
	go func() {
		for {
			select {
			case <-ctx.Done():
				currMax = 0
				done <- struct{}{}
			case <-timer.C:
				if curr < currMax {
					go func() {
						defer func() {
							if curr > 0 {
								atomic.AddInt64(&curr, -1)
							}
						}()
						atomic.AddInt64(&curr, 1)
						sync(cancel)
					}()
				}
			}
		}
	}()
	return cancel, done
}
