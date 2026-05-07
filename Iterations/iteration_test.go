package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

//trying benchmarks
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		//bechmarks test loop code only
		Repeat("a")
	}
}

func ExampleRepeat() {
    repeat := Repeat("a")
    fmt.Println(repeat)
    // Output: aaaaa
}