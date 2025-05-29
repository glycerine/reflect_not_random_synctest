package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// look at receives and timers together.
// Are they random with respect to each other? nope. grrr.
// It appears timers are always selected last.
//
/* with one timer in the middle of the cases: always last
GOTRACEBACK=all GOEXPERIMENT=synctest go test -v -run 801
faketime = true
=== RUN   Test801_synctest_chan_receives
on i=0, chan read. chosen=1, received: 1, at 2000-01-01T00:00:00.000Z
on i=1, chan read. chosen=7, received: 6, at 2000-01-01T00:00:00.000Z
on i=2, chan read. chosen=9, received: 8, at 2000-01-01T00:00:00.000Z
on i=3, chan read. chosen=3, received: 3, at 2000-01-01T00:00:00.000Z
on i=4, chan read. chosen=10, received: 9, at 2000-01-01T00:00:00.000Z
on i=5, chan read. chosen=8, received: 7, at 2000-01-01T00:00:00.000Z
on i=6, chan read. chosen=0, received: 0, at 2000-01-01T00:00:00.000Z
on i=7, chan read. chosen=4, received: 4, at 2000-01-01T00:00:00.000Z
on i=8, chan read. chosen=2, received: 2, at 2000-01-01T00:00:00.000Z
on i=9, chan read. chosen=6, received: 5, at 2000-01-01T00:00:00.000Z
on i=10, timer fired. chosen=5, received: 2000-01-01 00:00:00 +0000 GMT m=+946307292.086710285, at 2000-01-01T00:00:00.000Z

// with two timers in the middle of the cases: timers
// always go last.

Compilation started at Thu May 29 08:38:32

GOTRACEBACK=all GOEXPERIMENT=synctest go test -v -run 801
faketime = true
=== RUN   Test801_synctest_chan_receives
on i=0, chan read. chosen=10, received: 8, at 2000-01-01T00:00:00.000Z
on i=1, chan read. chosen=9, received: 7, at 2000-01-01T00:00:00.000Z
on i=2, chan read. chosen=0, received: 0, at 2000-01-01T00:00:00.000Z
on i=3, chan read. chosen=7, received: 5, at 2000-01-01T00:00:00.000Z
on i=4, chan read. chosen=8, received: 6, at 2000-01-01T00:00:00.000Z
on i=5, chan read. chosen=3, received: 3, at 2000-01-01T00:00:00.000Z
on i=6, chan read. chosen=2, received: 2, at 2000-01-01T00:00:00.000Z
on i=7, chan read. chosen=4, received: 4, at 2000-01-01T00:00:00.000Z
on i=8, chan read. chosen=11, received: 9, at 2000-01-01T00:00:00.000Z
on i=9, chan read. chosen=1, received: 1, at 2000-01-01T00:00:00.000Z
on i=10, timer fired. chosen=5, received: 2000-01-01 00:00:00 +0000 GMT m=+946307138.581471794, at 2000-01-01T00:00:00.000Z
on i=11, timer fired. chosen=6, received: 2000-01-01 00:00:00 +0000 GMT m=+946307138.581471794, at 2000-01-01T00:00:00.000Z


Compilation started at Thu May 29 08:39:57

GOTRACEBACK=all GOEXPERIMENT=synctest go test -v -run 801
faketime = true
=== RUN   Test801_synctest_chan_receives
on i=0, chan read. chosen=1, received: 1, at 2000-01-01T00:00:00.000Z
on i=1, chan read. chosen=9, received: 7, at 2000-01-01T00:00:00.000Z
on i=2, chan read. chosen=7, received: 5, at 2000-01-01T00:00:00.000Z
on i=3, chan read. chosen=10, received: 8, at 2000-01-01T00:00:00.000Z
on i=4, chan read. chosen=2, received: 2, at 2000-01-01T00:00:00.000Z
on i=5, chan read. chosen=0, received: 0, at 2000-01-01T00:00:00.000Z
on i=6, chan read. chosen=11, received: 9, at 2000-01-01T00:00:00.000Z
on i=7, chan read. chosen=4, received: 4, at 2000-01-01T00:00:00.000Z
on i=8, chan read. chosen=8, received: 6, at 2000-01-01T00:00:00.000Z
on i=9, chan read. chosen=3, received: 3, at 2000-01-01T00:00:00.000Z
on i=10, timer fired. chosen=5, received: 2000-01-01 00:00:00 +0000 GMT m=+946307053.389413297, at 2000-01-01T00:00:00.000Z
on i=11, timer fired. chosen=6, received: 2000-01-01 00:00:00 +0000 GMT m=+946307053.389413297, at 2000-01-01T00:00:00.000Z

realtime okay? yes.

 go test -v -run 801
faketime = false
=== RUN   Test801_synctest_chan_receives
on i=0, timer fired. chosen=5, received: 2025-05-29 08:41:37.182439496 +0100 BST m=+0.000453402, at 2025-05-29T07:41:37.182Z
on i=1, chan read. chosen=1, received: 1, at 2025-05-29T07:41:37.182Z
on i=2, chan read. chosen=4, received: 4, at 2025-05-29T07:41:37.182Z
on i=3, chan read. chosen=7, received: 5, at 2025-05-29T07:41:37.182Z
on i=4, chan read. chosen=3, received: 3, at 2025-05-29T07:41:37.182Z
on i=5, chan read. chosen=11, received: 9, at 2025-05-29T07:41:37.182Z
on i=6, chan read. chosen=0, received: 0, at 2025-05-29T07:41:37.182Z
on i=7, chan read. chosen=9, received: 7, at 2025-05-29T07:41:37.182Z
on i=8, chan read. chosen=2, received: 2, at 2025-05-29T07:41:37.182Z
on i=9, chan read. chosen=8, received: 6, at 2025-05-29T07:41:37.182Z
on i=10, timer fired. chosen=6, received: 2025-05-29 08:41:37.18245407 +0100 BST m=+0.000468574, at 2025-05-29T07:41:37.182Z
on i=11, chan read. chosen=10, received: 8, at 2025-05-29T07:41:37.182Z

realtime again: good, order of timers varies.

 go test -v -run 801
faketime = false
=== RUN   Test801_synctest_chan_receives
on i=0, timer fired. chosen=6, received: 2025-05-29 08:42:05.307762552 +0100 BST m=+0.000836471, at 2025-05-29T07:42:05.307Z
on i=1, chan read. chosen=4, received: 4, at 2025-05-29T07:42:05.308Z
on i=2, chan read. chosen=10, received: 8, at 2025-05-29T07:42:05.308Z
on i=3, chan read. chosen=8, received: 6, at 2025-05-29T07:42:05.308Z
on i=4, chan read. chosen=7, received: 5, at 2025-05-29T07:42:05.308Z
on i=5, timer fired. chosen=5, received: 2025-05-29 08:42:05.307736588 +0100 BST m=+0.000811335, at 2025-05-29T07:42:05.308Z
on i=6, chan read. chosen=0, received: 0, at 2025-05-29T07:42:05.308Z
on i=7, chan read. chosen=3, received: 3, at 2025-05-29T07:42:05.308Z
on i=8, chan read. chosen=9, received: 7, at 2025-05-29T07:42:05.308Z
on i=9, chan read. chosen=1, received: 1, at 2025-05-29T07:42:05.308Z
on i=10, chan read. chosen=11, received: 9, at 2025-05-29T07:42:05.308Z
on i=11, chan read. chosen=2, received: 2, at 2025-05-29T07:42:05.308Z

realtime:

 go test -v -run 801
faketime = false
=== RUN   Test801_synctest_chan_receives
on i=0, chan read. chosen=2, received: 2, at 2025-05-29T07:42:47.935Z
on i=1, chan read. chosen=10, received: 8, at 2025-05-29T07:42:47.935Z
on i=2, chan read. chosen=11, received: 9, at 2025-05-29T07:42:47.935Z
on i=3, chan read. chosen=9, received: 7, at 2025-05-29T07:42:47.935Z
on i=4, chan read. chosen=7, received: 5, at 2025-05-29T07:42:47.935Z
on i=5, chan read. chosen=4, received: 4, at 2025-05-29T07:42:47.935Z
on i=6, chan read. chosen=3, received: 3, at 2025-05-29T07:42:47.935Z
on i=7, timer fired. chosen=6, received: 2025-05-29 08:42:47.935489343 +0100 BST m=+0.000838372, at 2025-05-29T07:42:47.935Z
on i=8, timer fired. chosen=5, received: 2025-05-29 08:42:47.93546375 +0100 BST m=+0.000812594, at 2025-05-29T07:42:47.935Z
on i=9, chan read. chosen=8, received: 6, at 2025-05-29T07:42:47.935Z
on i=10, chan read. chosen=0, received: 0, at 2025-05-29T07:42:47.935Z
on i=11, chan read. chosen=1, received: 1, at 2025-05-29T07:42:47.935Z

*/
func Test801_synctest_chan_receives(t *testing.T) {

	bubbleOrNot(func() {

		attempts := 1
		for range attempts {

			t0 := time.Now()
			//vv("start test801")
			var chans []chan int
			N := 10
			for j := range N {
				ch := make(chan int, 1)
				ch <- j
				chans = append(chans, ch)
			}

			var cases []reflect.SelectCase

			split := 4
			for k, ch := range chans {
				if k > split {
					break
				}
				cases = append(cases,
					reflect.SelectCase{
						Dir:  reflect.SelectRecv,
						Chan: reflect.ValueOf(ch),
					})
			}
			// and one timer in the middle
			ti := time.NewTimer(0)
			cases = append(cases,
				reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(ti.C),
				})

			// and now another timer in the middle
			ti2 := time.NewTimer(0)
			cases = append(cases,
				reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(ti2.C),
				})

			for k, ch := range chans {
				if k <= split {
					continue
				}
				cases = append(cases,
					reflect.SelectCase{
						Dir:  reflect.SelectRecv,
						Chan: reflect.ValueOf(ch),
					})
			}

			for i := range N + 2 {

				chosen, recvVal, recvOK := reflect.Select(cases)
				if !recvOK {
					panic("why not recvOK ?")
				}
				switch x := recvVal.Interface().(type) {
				case int:
					fmt.Printf("on i=%v, chan read. chosen=%v, received: %v, at %v\n", i, chosen, x, nice(time.Now().In(gtz)))
				case time.Time:
					fmt.Printf("on i=%v, timer fired. chosen=%v, received: %v, at %v\n", i, chosen, x, nice(time.Now().In(gtz)))
				}
			}

			now := time.Now()
			if faketime {
				if !t0.Equal(now) {
					t.Fatalf("we have a problem, Houston. t0=%v, but now=%v", t0, now)
				}
			}
			//vv("end test801") // shows same time as start, good.
		}

	})
}
