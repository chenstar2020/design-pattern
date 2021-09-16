package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	test := NewImageViewer("test")
	test.Display()
}
