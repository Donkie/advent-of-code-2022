package main

type Direction bool

const (
	Left  Direction = false
	Right Direction = true
)

type JetStream struct {
	dir []Direction
	idx int
}

func (j *JetStream) PopDir() (d Direction) {
	d = j.dir[j.idx]
	j.idx = (j.idx + 1) % len(j.dir)
	return
}
