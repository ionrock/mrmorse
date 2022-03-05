package main

import (
	"testing"
)

type morseTest struct {
	input   string
	output  string
	isMorse bool
}

var morseTests = []morseTest{
	{"HELLO WORLD!", ".... . .-.. .-.. --- / .-- --- .-. .-.. -.. -.-.--", false},
	{".... . .-.. .-.. --- / .-- --- .-. .-.. -.. -.-.--", "HELLO WORLD!", true},
}

func TestIsMorseCode(t *testing.T) {
	for _, mt := range morseTests {
		result := IsMorseCode(mt.input)
		if mt.isMorse != result {
			t.Errorf("for input: %v, expected: %v, got: %v", mt.input, mt.isMorse, result)
		}
	}
}

func TestTranslate(t *testing.T) {
	for _, mt := range morseTests {
		result, err := Translate(mt.input)
		if err != nil {
			t.Errorf("Translate error: %v", err)
			return
		}

		if mt.output != result {
			t.Errorf("for input: %v, expected: %v, got: %v", mt.input, mt.output, result)
		}
	}
}
