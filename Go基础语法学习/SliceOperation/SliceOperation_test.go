package SliceOperation

import "testing"

func TestAdd(t *testing.T) {
	a := []int{1, 2, 3, 4}
	got := Add(a, 2, 5)
	want := []int{1, 2, 5, 3, 4}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v gotcap:%d gotlen:%d,want %v", got, cap(got), len(got), want)
			return
		}
	}
}
func TestDelete(t *testing.T) {
	a := []int{1, 2, 3, 4}
	got := Delete(a, 0)
	want := []int{2, 3, 4}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v gotcap:%d gotlen:%d,want %v", got, cap(got), len(got), want)
			return
		}
	}
}
