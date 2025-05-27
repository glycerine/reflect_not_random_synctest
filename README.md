reflect_not_random_synctest
===========================

for https://github.com/golang/go/issues/73876

I noticed that under the go 1.24.3 GOEXPERIMENT=synctest
version of testing/synctest, there is a small
randomization failure in reflect.Select.

The first call to reflect.Select always 
returns the 0 case as chosen.

This does not happen when not under synctest.

test logs:

~~~
-*- mode: compilation; default-directory: "~/go/src/github.com/glycerine/reflect_not_random_synctest/" -*-
Compilation started at Tue May 27 06:58:50

GOTRACEBACK=all GOEXPERIMENT=synctest go test -v
faketime = true
=== RUN   Test800_synctest_all_timers_dur_0_fire_now
panic: 10000 attempts, chosen=0 always on the first reflect.Select()
~~~

~~~
-*- mode: compilation; default-directory: "~/go/src/github.com/glycerine/reflect_not_random_synctest/" -*-
Compilation started at Tue May 27 06:56:10

 go test -v
faketime = false
=== RUN   Test800_synctest_all_timers_dur_0_fire_now
--- PASS: Test800_synctest_all_timers_dur_0_fire_now (0.20s)
PASS
ok  	github.com/glycerine/reflect_not_random_synctest	0.305s

Compilation finished at Tue May 27 06:56:10
~~~
