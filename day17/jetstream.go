package main

type Direction uint8

const (
	Left Direction = iota
	Right
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
