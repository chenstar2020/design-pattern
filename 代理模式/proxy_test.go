package proxy

import (
	"fmt"
	"testing"
)

func TestProxy_Do(t *testing.T) {
	proxy := Proxy{}
	fmt.Println(proxy.Do())
}

