package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "value"}

	t.Run("known word", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "value"

		assertNoError(t, err)
		assertString(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("other")
		assertError(t, ErrNotFound, err)

	})
}

func TestAdd(t *testing.T) {
	dict := Dict{}
	dict.Add("test", "value")

	t.Run("known word", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "value"

		assertNoError(t, err)
		assertString(t, want, got)
	})

	t.Run("add existing word", func(t *testing.T) {
		err := dict.Add("test", "value 2")

		assertError(t, err, ErrWordExists)
	})

}

func TestUpdate(t *testing.T) {
	dict := Dict{}
	dict.Add("test", "value")

	t.Run("update existed word", func(t *testing.T) {
		dict.Update("test", "value2")
		got, err := dict.Search("test")
		want := "value2"

		assertNoError(t, err)
		assertString(t, want, got)
	})

	t.Run("update no existing word", func(t *testing.T) {
		err := dict.Update("other", "value 2")

		assertError(t, ErrWordDoesNotExists, err)
	})

}

func TestDelete(t *testing.T) {

	t.Run("update existed word", func(t *testing.T) {
		dict := Dict{}
		dict.Add("test", "value")

		dict.Delete("test")

		_, err := dict.Search("test")

		assertError(t, ErrNotFound, err)
	})

}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Should NOT return error, but got %v", err)
	}
}

func assertError(t *testing.T, want error, err error) {
	t.Helper()
	if err != want {
		t.Errorf("want %v but got %v", want, err)
	}
}

func assertString(t *testing.T, want string, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}
}
