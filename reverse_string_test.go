package reverse_string

import "testing"

func testReverse(t *testing.T, aInput string, aResult string) {
	testValue := ReverseString(aInput)

	if testValue != aResult {
		t.Errorf("Invalid ReverseString output: '%v' expected but '%v' found", aResult, testValue)
	}
}

func TestAnsi(t *testing.T) {
	testReverse(t, "a s d", "d s a")
}

func TestMultiLine(t *testing.T) {
	testReverse(t, "a\x0d\x0as\x0dd\x0af", "f\x0ad\x0ds\x0d\x0aa")
}

func TestGraphems(t *testing.T) {
	testReverse(t, "I̪͉̜̼̼̣̟̣ ̰̟̥̞̹c͈͔͇̼a̙̹̼̦̲̞n̙̺̳̟ͅ ̤̗d̘̭̙̪̦o̬̲̜̺ ̲̬̝t̺̖̗̩̱h̟̟̱i̹s̹̱.̯̖̝̯̟̜̥", ".̯̖̝̯̟̜̥s̹̱i̹h̟̟̱t̺̖̗̩̱ ̲̬̝o̬̲̜̺d̘̭̙̪̦ ̤̗n̙̺̳̟ͅa̙̹̼̦̲̞c͈͔͇̼ ̰̟̥̞̹I̪͉̜̼̼̣̟̣")
}
