package do

import (
	"fmt"
	"testing"
	"time"
)

func TestDoSthUntilTimes(t *testing.T) {
	ct := 4

	do := func() {
		fmt.Printf("%d\n", ct)
		ct++
	}

	check := func() bool {
		return ct > 5
	}
	result := DoSthUntilTimes(do, 4, 3*time.Second, check)
	fmt.Printf("result = %v\n", result)
}
