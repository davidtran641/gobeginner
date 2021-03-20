package server

import (
	"io/ioutil"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "123456789")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	got, _ := ioutil.ReadAll(file)

	want := "abc"

	test.AssertEqual(t, want, string(got))
}
