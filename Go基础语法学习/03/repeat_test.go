package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected `%s` repeated `%s`", expected, repeated)
	}
}
func BenchmarkRepeated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	//Output:aaaaa
}
