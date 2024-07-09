package quiz_test

import (
	"day06/practise02/quiz"
	"testing"
)

func TestAddSum(t *testing.T) {
	a, b := 5, 10
	expected := 18

	res := quiz.AddSum(a, b)

	if res != expected {
		t.Errorf("AddSum function not mach : %d+%d=%d", a, b, expected)
	}
}
