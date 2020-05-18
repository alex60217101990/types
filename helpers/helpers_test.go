package helpers

import (
	"fmt"
	"testing"
)

func TestIP2int(t *testing.T) {
	ipnet, err := ParseIPFromString("172.17.0.1")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(IP2int(ipnet.IP))
}
