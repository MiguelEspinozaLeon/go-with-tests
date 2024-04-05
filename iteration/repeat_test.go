package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {

	Repeat("a", b.N)

}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	//Output: aaaaa
}
