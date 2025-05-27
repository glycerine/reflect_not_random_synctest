reflect_not_random_synctest
===========================

I noticed that under the go 1.24.3 
testing/synctest, there is a small
randomization failure in reflect.Select.

The first call to reflect.Select always 
returns the 0 case as chosen.

This does not happen when not under synctest.

Transcript of runs. Note that green PASS is BAD here.

It means that under synctest, reflect.Select returned
case 0 on the first call 1000 times in a row.

~~~
-*- mode: compilation; default-directory: "~/go/src/github.com/glycerine/reflect_not_random_synctest/" -*-
Compilation started at Tue May 27 06:47:51

GOTRACEBACK=all GOEXPERIMENT=synctest go test -v
faketime = true
=== RUN   Test800_synctest_all_timers_dur_0_fire_now
--- PASS: Test800_synctest_all_timers_dur_0_fire_now (0.00s)
PASS
ok  	github.com/glycerine/reflect_not_random_synctest	0.161s

Compilation finished at Tue May 27 06:47:51

~~~

~~~
-*- mode: compilation; default-directory: "~/go/src/github.com/glycerine/reflect_not_random_synctest/" -*-
Compilation started at Tue May 27 06:47:27

go test -v
faketime = false
=== RUN   Test800_synctest_all_timers_dur_0_fire_now
--- PASS: Test800_synctest_all_timers_dur_0_fire_now (0.00s)
PASS
ok  	github.com/glycerine/reflect_not_random_synctest	0.161s

Compilation finished at Tue May 27 06:47:28
~~~
