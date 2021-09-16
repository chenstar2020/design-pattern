package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	single1 := GetInstance()
	single2 := GetInstance()
	if single1 == single2{
		fmt.Println("equal")
	}
}
