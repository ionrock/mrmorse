package main

import (
	"testing"
)

func TestIsMorseCode(t *testing.T) {
	type morseTest struct {
		input string
		want  bool
	}
	morseTests := []morseTest{
		{"Hello World!", false},
		{"-.-- --- ..- / ... .... --- ..- .-.. -.. / .-.. . .- .-. -. / -.-. .--", true},
	}

	for _, mt := range morseTests {
		result := IsMorseCode(mt.input)
		if mt.want != result {
			t.Errorf("for input: %v, expected: %v, got: %v", mt.input, mt.want, result)
		}
	}
}
