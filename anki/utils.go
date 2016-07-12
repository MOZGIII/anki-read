package anki

import (
	"database/sql"
	"io"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var tmpFileNames []string

func OpenSQLiteFromReader(r io.Reader) (*sql.DB, error) {
	tmpfile, err := ioutil.TempFile("", "anki-go")
	if err != nil {
		return nil, err
	}
	name := tmpfile.Name()
	tmpFileNames = append(tmpFileNames, name)
	if _, err := io.Copy(tmpfile, r); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}

	return sql.Open("sqlite3", name)
}

func CleanupSQLiteTmpFiles() {
	for _, name := range tmpFileNames {
		os.Remove(name)
	}
}
