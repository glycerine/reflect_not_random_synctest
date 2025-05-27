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

Confirmed on Linux too, with fresh go1.24.3 install.
~~~
(base) /usr/local/tmp/go1.24.3/reflect_not_random_synctest (master) $  GOEXPERIMENT=synctest go test -v
 GOEXPERIMENT=synctest go test -v
 faketime = true
 === RUN   Test800_synctest_reflect_select
 panic: 10000 attempts, chosen=0 always on the first reflect.Select()
 
 goroutine 19 [running, synctest group 18]:
 github.com/glycerine/reflect_not_random_synctest.Test800_synctest_reflect_select.func1()
         /usr/local/tmp/go1.24.3/reflect_not_random_synctest/synctest_test.go:72 +0x35b
         created by internal/synctest.Run in goroutine 18
                 /home/jaten/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.3.linux-amd64/src/runtime/synctest.go:178 +0x10d
                 exit status 2
                 FAIL    github.com/glycerine/reflect_not_random_synctest        0.310s
                 (base) jaten@aorus /usr/local/tmp/go1.24.3/reflect_not_random_synctest (master) $ 
~~~
