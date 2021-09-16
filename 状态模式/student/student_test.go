package student

import "testing"

func TestStudent(t *testing.T){
	sc := ScoreContext{}
	sc.checkState(10)
	sc.setState()
	sc.checkState(80)
	sc.setState()
	sc.checkState(80)
	sc.setState()
}
