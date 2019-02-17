package utils

import (
	"testing"
)

var greatTests = []struct {
	input  string
	output string
}{
	{"Sumo", "Hello Sumo! I'm not in the $GOPATH!!!\n"},
	{"Sumo1", "Hello Sumo1! I'm not in the $GOPATH!!!\n"},
	//{"", "no name provided"},
}

func TestGreet(t *testing.T) {
	for _, tt := range greatTests {
		if s := Greet(tt.input); s != nil {
			t.Errorf("Invalid input (%s)", tt.input)
		}
	}
}
