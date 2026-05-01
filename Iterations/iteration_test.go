package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

//trying benchmarks
func BenchmarRepeat(b *testing.B) {
	for b.Loop() {
		//bechmarks test loop code only
		Repeat("a")
	}
}