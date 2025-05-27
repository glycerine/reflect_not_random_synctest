package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test800_synctest_all_timers_dur_0_fire_now(t *testing.T) {

	if !faketime {
		//t.Skip("test only for synctest.")
		//return
	}

	bubbleOrNot(func() {
		// "SimNet using synctest depends on all the times set to duration 0/now firing before we quiese to durable blocking. verify this assumption under synctest. yes: note the Go runtime implementation does a select with a default: so it will discard the timer alert rather than block. Update: arg. no, the runtime does a special thing where it does not execute that select until it has a goro ready to accept it, so it always suceeds, I think.

		// really? yep
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
			panic(fmt.Sprintf("in %v attempts, the first reflect.Select case was always 0", attempts))
		}

	})
}
