package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test800_synctest_reflect_select(t *testing.T) {

	if !faketime {
		//t.Skip("test only for synctest.")
		//return
	}

	bubbleOrNot(func() {

		// really, is synctest not randomizing reflect.Select? yep.
		firstCaseNonZeroSeen := false
		attempts := 10000
		for range attempts {

			t0 := time.Now()
			//vv("start test800")
			var timers []*time.Timer
			N := 10
			order := make(map[*time.Timer]int)
			for range N {
				ti := time.NewTimer(0)
				order[ti] = len(timers)
				timers = append(timers, ti)
			}

			var cases []reflect.SelectCase
			for _, ti := range timers {
				cases = append(cases,
					reflect.SelectCase{
						Dir:  reflect.SelectRecv,
						Chan: reflect.ValueOf(ti.C),
					})
			}
			got := 0
			for i, ti := range timers {
				//  <-ti.C
				chosen, recvVal, recvOK := reflect.Select(cases)
				if !recvOK {
					panic("why not recvOK ?")
				}
				fmt.Printf("on i=%v, chosen=%v, timer %v: %v\n", i, chosen, order[ti], nice(recvVal.Interface().(time.Time).In(gtz)))
				//vv("on i=%v, chosen=%v, timer %v: %v", i, chosen, order[ti], recvVal)
				_, _, _, _ = i, chosen, recvVal, ti
				got++

				if i == 0 && chosen != 0 {
					firstCaseNonZeroSeen = true
				}
			}
			now := time.Now()
			if faketime {
				if !t0.Equal(now) {
					t.Fatalf("we have a problem, Houston. t0=%v, but now=%v", t0, now)
				}
				//vv("got %v timers firing now", got)
				if got != N {
					t.Fatalf("expected all N=%v timers to fire, not %v", N, got)
				}
			}
			//vv("end test800") // shows same time as start, good.
		}

		if !firstCaseNonZeroSeen {
			panic(fmt.Sprintf("%v attempts, chosen=0 always on the first reflect.Select()", attempts))
		}

	})
}
