package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat with iteration", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Repeat with built in", func(t *testing.T) {
		repeated := RepeatWithBuiltIn("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

// BenchmarkRepeat - let's see how good is it
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

// ExampleRepeat -  Simple example of Repeat function.
func ExampleRepeat() {
	result := Repeat("C", 3)
	fmt.Println(result)
	// Output: CCC
}
