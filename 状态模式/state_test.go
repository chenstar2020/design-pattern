package state

import "testing"

func TestState(t *testing.T){
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 8;i++ {
		todayAndNext()
	}
}