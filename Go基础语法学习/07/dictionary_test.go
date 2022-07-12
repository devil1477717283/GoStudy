package _7

import "testing"

func TestSearch(t *testing.T) {
	assertString := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got `%s` want `%s`", got, want)
		}
	}
	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == NotFound {
			t.Error(NotFound)
		}
	}
	t.Run("Find a word in map", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}
		got := Search(dictionary, "test")
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run("Use user-define type", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
		assertError(t, err)
	})
}
