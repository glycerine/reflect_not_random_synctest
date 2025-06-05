package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

var _ = fmt.Printf

func Test802_synctest_permutations(t *testing.T) {

	bubbleOrNot(func() {

		N := 10
		first := make([]int, N+2)
		last := make([]int, N+2)

		attempts := 100_000
		for range attempts {

			t0 := time.Now()
			//vv("start test802")
			var chans []chan int
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

			lasti := N + 1
			for i := range N + 2 {

				chosen, recvVal, recvOK := reflect.Select(cases)
				if !recvOK {
					panic("why not recvOK ?")
				}
				switch i {
				case 0:
					first[chosen]++
				case lasti:
					last[chosen]++
				}

				switch x := recvVal.Interface().(type) {
				case int:
					_ = x
					//fmt.Printf("on i=%v, chan read. chosen=%v, received: %v, at %v\n", i, chosen, x, nice(time.Now().In(gtz)))
				case time.Time:
					//fmt.Printf("on i=%v, timer fired. chosen=%v, received: %v, at %v\n", i, chosen, x, nice(time.Now().In(gtz)))
				}
			}

			now := time.Now()
			if faketime {
				if !t0.Equal(now) {
					t.Fatalf("we have a problem, Houston. t0=%v, but now=%v", t0, now)
				}
			}
			//vv("end test802") // shows same time as start, good.
		} // end attempts
		vv("attempts = %v", attempts)
		vv("first = '%#v'", first)
		vv(" last = '%#v'", last)
	})
}

/*
Compilation started at Thu Jun  5 08:05:55

/usr/local/dev-go/go/bin/go test -v -run 802
faketime = false
=== RUN   Test802_synctest_permutations

recv_compare_test.go:103 [pid 95884] 2025-06-05 07:05:57.620 +0000 UTC attempts = 100000

recv_compare_test.go:104 [pid 95884] 2025-06-05 07:05:57.620 +0000 UTC

first = '[]int{8272, 8406, 8485, 8420, 8449, 8321, 8150, 8274, 8394, 8263, 8251, 8315}'

recv_compare_test.go:105 [pid 95884] 2025-06-05 07:05:57.620 +0000 UTC

last = '[]int{8357, 8400, 8263, 8585, 8422, 8220, 8257, 8261, 8266, 8236, 8387, 8346}'

--- PASS: Test802_synctest_permutations (1.64s)
PASS
ok  	github.com/glycerine/reflect_not_random_synctest	1.790s
*/

/*
Compilation started at Thu Jun  5 08:06:53

go1.24.3:

GOTRACEBACK=all GOEXPERIMENT=synctest go test -v -run 802
faketime = true
=== RUN   Test802_synctest_permutations

recv_compare_test.go:103 [pid 96171] 2000-01-01 00:00:00.000 +0000 UTC attempts = 100000

recv_compare_test.go:104 [pid 96171] 2000-01-01 00:00:00.000 +0000 UTC

first = '[]int{9947, 10021, 9999, 10057, 9915, 0, 0, 9986, 10083, 9990, 10029, 9973}'

recv_compare_test.go:105 [pid 96171] 2000-01-01 00:00:00.000 +0000 UTC

last = '[]int{0, 0, 0, 0, 0, 0, 100000, 0, 0, 0, 0, 0}'

--- PASS: Test802_synctest_permutations (1.69s)
*/

/*
go 1.25 prerelease

Compilation started at Thu Jun  5 08:10:17

GOTRACEBACK=all GOEXPERIMENT=synctest /usr/local/dev-go/go/bin/go test -v -run 802
faketime = true
=== RUN   Test802_synctest_permutations

recv_compare_test.go:103 [pid 96504] 2000-01-01 00:00:00.000 +0000 UTC attempts = 100000

recv_compare_test.go:104 [pid 96504] 2000-01-01 00:00:00.000 +0000 UTC

first = '[]int{8416, 8365, 8383, 8221, 8346, 8412, 8391, 8241, 8243, 8116, 8414, 8452}'

recv_compare_test.go:105 [pid 96504] 2000-01-01 00:00:00.000 +0000 UTC

last = '[]int{8336, 8293, 8419, 8378, 8266, 8265, 8158, 8357, 8253, 8453, 8511, 8311}'

--- PASS: Test802_synctest_permutations (1.64s)
PASS
ok  	github.com/glycerine/reflect_not_random_synctest	1.868s

*/
