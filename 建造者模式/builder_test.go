package builder

import (
	"fmt"
	"testing"
)

func TestBuilder1(t *testing.T) {
	builder1 := &Builder1{}
	director := NewDirector(builder1)
	director.Construct()
	fmt.Println("result:", builder1.GetResult())

	builder2 := &Builder2{}
	director = NewDirector(builder2)
	director.Construct()
	fmt.Println("result:", builder2.GetResult())
}
