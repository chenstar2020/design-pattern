package facade

import "testing"

var expect = "A module running\nB module running"

func TestFacade(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != expect{
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}