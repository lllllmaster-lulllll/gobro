package list

import "testing"

func Test_initRing(t *testing.T) {
	initRing()
	r := New(20)
	r.PrintRing()
}
