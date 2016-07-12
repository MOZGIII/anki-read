package anki

import (
	"bytes"
	"database/sql"
)

type Anki2DB struct {
	db *sql.DB
}

type Card struct {
	Jap1 string
	Jap2 string
	Desc string
}

func NewAnki2DB(db *sql.DB) *Anki2DB {
	return &Anki2DB{db: db}
}

func ParseAnkiStrings(b []byte) (s []string) {
	for {
		i := bytes.IndexByte(b, 31)
		if i == -1 {
			break
		}
		s = append(s, string(b[0:i]))
		b = b[i+1:]
	}
	return
}

func (a *Anki2DB) Cards() (cards []Card, e error) {
	rows, err := a.db.Query("SELECT  sfld, flds FROM notes")
	if err != nil {
		e = err
		return
	}
	defer rows.Close()
	for rows.Next() {
		card, err := scanCard(rows)
		if err != nil {
			return nil, err
		}
		cards = append(cards, *card)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return
}

func scanCard(rows *sql.Rows) (*Card, error) {
	var (
		sfld string
		flds []byte
	)
	if err := rows.Scan(&sfld, &flds); err != nil {
		return nil, err
	}
	fldsa := ParseAnkiStrings(flds)
	return &Card{Jap1: sfld, Jap2: fldsa[2], Desc: fldsa[1]}, nil
}
