package genericstests

import "testing"

func TestAssertFunction(t *testing.T) {
	t.Run("Assert with interger value", func(t *testing.T) {
		AssertFunction(t, 50, sum(30, 30))
	})

	t.Run("Assert with string value", func(t *testing.T) {
		result := "Generics Tests"
		AssertFunction(t, "Generics Tests", result)
	})
}
func AssertFunction[T comparable](t *testing.T, expected, result T) {
	t.Helper()
	if expected != result {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func sum(val1, val2 int) int {
	return val1 + val2
}
