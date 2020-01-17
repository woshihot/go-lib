package do

import (
	"time"
)

/*
 * do something not stop ,except times out or check is true
 *
 * Params:
 *      - do : exe func
 *      - times : max times of func
 * 		- frequency : func exe frequency
 * 		- check : check func
 * Return:
 *     - is check func true
 */
func DoSthUntilTimes(do func(), times int, frequency time.Duration, check func() bool) bool {
	if times <= 0 {
		return false
	}
	do()
	if check() {
		return true
	} else {
		if times > 1 {
			return doSthForTimes(do, times, frequency, check)
		}
		return false
	}
}

func DoSthUntil(do func(), frequency time.Duration, check func() bool) {
	do()
	if check() {
		return
	} else {
		doSthFor(do, frequency, check)
	}
}

func doSthForTimes(do func(), times int, frequency time.Duration, check func() bool) bool {
	currentTime := 1
	t := time.NewTicker(frequency)
	for range t.C {
		if currentTime < times {
			do()
			currentTime++
			if check() {
				t.Stop()
				return true
			}
		} else {
			t.Stop()
			return false
		}
	}
	return false
}

func doSthFor(do func(), frequency time.Duration, check func() bool) {
	t := time.NewTicker(frequency)
	for range t.C {
		do()
		if check() {
			t.Stop()
		}
	}
}
