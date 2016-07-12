package anki

type Loadstate struct {
	Loaded bool
}

func (l *Loadstate) Set(loaded bool) {
	l.Loaded = loaded
}

func (l *Loadstate) SetLoaded() {
	l.Set(true)
}
