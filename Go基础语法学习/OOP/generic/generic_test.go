package generic

import "testing"

func assertEqual[T Addable](t *testing.T, want T, got T) {
	t.Helper()
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
func TestAdd(t *testing.T) {

	t.Run("float", func(t *testing.T) {
		got := Add(2.5, 3.0)
		want := 5.5
		assertEqual(t, want, got)
	})
	t.Run("string", func(t *testing.T) {
		got := Add("Hello", "Go")
		want := "HelloGo"
		assertEqual(t, want, got)
	})
	t.Run("int", func(t *testing.T) {
		got := Add(1, 2)
		want := 3
		assertEqual(t, want, got)
	})
}
