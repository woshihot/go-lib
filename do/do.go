package do

import "time"

func DoSthUntilTimes(do func(), times int, frequency time.Duration, check func() bool) bool {
	if times <= 0 {
		return false
	}
	do()
	if check() {
		return true
	} else {
		if times > 1 {
			return doSthFor(do, times, frequency, check)
		}
		return false
	}
}
func doSthFor(do func(), times int, frequency time.Duration, check func() bool) bool {
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
