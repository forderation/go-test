package main

import (
	"testing"
)

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExist        = DictionaryErr("word already exist")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(keyword string) (string, error) {
	definition, ok := dictionary[keyword]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(keyword, definition string) error {
	_, err := d.Search(keyword)
	switch err {
	case ErrNotFound:
		d[keyword] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(keyword, definition string) error {
	_, err := d.Search(keyword)
	switch err {
	case nil:
		d[keyword] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(keyword string) {
	delete(d, keyword)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertErrorMap(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAddWord(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		keyword := "test"
		definition := "this is just a test"
		err := dictionary.Add(keyword, definition)
		assertErrorMap(t, err, nil)
		assertDefinition(t, dictionary, keyword, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		keyword := "test"
		definition := "this is just a test"
		dictionary := Dictionary{keyword: definition}
		err := dictionary.Add(keyword, definition)
		assertErrorMap(t, err, ErrWordExist)
		assertDefinition(t, dictionary, keyword, definition)
	})

}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertErrorMap(t, got, ErrNotFound)
	})
}

func TestUpdateWord(t *testing.T) {
	t.Run("update exist word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertErrorMap(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update unknown word", func(t *testing.T) {
		word := "test"
		definition := "new definition"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertErrorMap(t, err, ErrWordDoesNotExist)
	})
}

func TestDeleteWord(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}
	dictionary.Delete(word)
	assertWordIsDeleted(t, dictionary, word)
}

func assertWordIsDeleted(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
