package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := 11
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
}
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
	//Output:15
}
