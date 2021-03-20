package server

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `
	[
		{"Name":"Julia", "Wins": 10},
		{"Name":"Bean", "Wins": 20}
	]
	`)
	defer cleanDatabase()

	store := NewFileSystemPlayerStore(database)

	got := store.GetLeague()

	want := League{
		{"Julia", 10},
		{"Bean", 20},
	}

	test.AssertEqual(t, want, got)

	// read again
	got = store.GetLeague()
	test.AssertEqual(t, want, got)
}

func TestGetPlayerScore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `
	[
		{"Name":"Julia", "Wins": 10},
		{"Name":"Bean", "Wins": 20}
	]
	`)
	defer cleanDatabase()

	store := NewFileSystemPlayerStore(database)
	got := store.GetPlayerScore("Julia")

	want := 10
	test.AssertEqual(t, want, got)

}

func TestRecordPlayerScore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `
	[
		{"Name":"Julia", "Wins": 10},
		{"Name":"Bean", "Wins": 20}
	]
	`)
	defer cleanDatabase()

	store := NewFileSystemPlayerStore(database)
	store.RecordScore("Julia")

	got := store.GetPlayerScore("Julia")
	want := 11
	test.AssertEqual(t, want, got)

}

func TestRecordScoreForNewPlayer(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `
	[
		{"Name":"Julia", "Wins": 10},
		{"Name":"Bean", "Wins": 20}
	]
	`)
	defer cleanDatabase()

	store := NewFileSystemPlayerStore(database)
	store.RecordScore("David")

	got := store.GetPlayerScore("David")
	want := 1
	test.AssertEqual(t, want, got)

}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
