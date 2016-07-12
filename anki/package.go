package anki

import (
	"archive/zip"
	"io"
)

type Package struct {
	Collection Collection
	Media      Media

	Log Logger

	Loadstate Loadstate
}

type Collection struct {
	DB        *Anki2DB
	Loadstate Loadstate
}
type Media struct {
	Loadstate Loadstate
}

func ReadFile(file string) (*Package, error) {
	return ReadFile2(file, DefaultLogger())
}

func ReadFile2(file string, log Logger) (*Package, error) {
	r, err := zip.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	pkg := New(log)
	if err := pkg.Load(r.Reader); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func New(log Logger) Package {
	return Package{Log: log}
}

func (p *Package) Load(r zip.Reader) error {
	for _, f := range r.File {
		switch f.Name {
		case "collection.anki2":
			rc, err := f.Open()
			if err != nil {
				return err
			}
			err = p.Collection.Load(rc)
			rc.Close()
			if err != nil {
				return err
			}
		}
	}
	p.Loadstate.SetLoaded()
	return nil
}

func (c *Collection) Load(r io.Reader) error {
	db, err := OpenSQLiteFromReader(r)
	if err != nil {
		return err
	}
	c.DB = NewAnki2DB(db)
	c.Loadstate.SetLoaded()
	return nil
}
