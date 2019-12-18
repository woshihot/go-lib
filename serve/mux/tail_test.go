package mux

import "testing"

func TestRest(t *testing.T) {

	Init("8081", nil)
	select {}
}
