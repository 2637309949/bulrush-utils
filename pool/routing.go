// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pool

import (
	"context"
	"time"

	addition "github.com/2637309949/bulrush-addition"
)

// rushLogger just for console log
var rushLogger = addition.RushLogger

// RoutingPoolWithTimer defined routing pool for work
// Example
// i := 0
// _, done := pool.RoutingPoolWithTimer(func(cancel context.CancelFunc) {
// 	time.Sleep(10 * time.Second)
// 	i++
// 	if i == 30 {
// 		cancel()
// 	}
// }, 10)
// <-done
func RoutingPoolWithTimer(worker func(context.CancelFunc), max int) (context.CancelFunc, chan struct{}) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	timer := time.NewTicker(1 * time.Second)
	done := make(chan struct{}, 0)
	init := false
	limiter := make(chan int, max)

	runWork := func(c chan int) {
		defer func() {
			<-c
		}()
		defer func() {
			if ret := recover(); ret != nil {
				rushLogger.Error("%v", ret)
			}
		}()
		c <- 1
		worker(cancel)
	}

	go func() {
		if !init {
			init = true
			for i := 0; i < max; i++ {
				go runWork(limiter)
			}
		}
		for {
			select {
			case <-ctx.Done():
				done <- struct{}{}
				break
			case <-timer.C:
				go runWork(limiter)
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
func RoutingPoolWithAutomatic(worker func(context.CancelFunc), max int) (context.CancelFunc, chan struct{}) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	done := make(chan struct{}, 0)
	init := false
	limiter := make(chan int, max)
	var runWork func(c chan int)
	runWork = func(c chan int) {
		defer func() {
			<-c
			select {
			case <-done:
			default:
				go runWork(c)
			}
		}()
		defer func() {
			if ret := recover(); ret != nil {
				rushLogger.Error("%v", ret)
			}
		}()
		c <- 1
		worker(cancel)
	}

	go func() {
		if !init {
			init = true
			for i := 0; i < max; i++ {
				go runWork(limiter)
			}
		}
		select {
		case <-ctx.Done():
			done <- struct{}{}
			close(done)
			break
		}
	}()
	return cancel, done
}
