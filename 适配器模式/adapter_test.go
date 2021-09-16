package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	adapter1 := NewAdaptee()
	target := NewAdapter(adapter1)
	fmt.Println(target.Request())
}