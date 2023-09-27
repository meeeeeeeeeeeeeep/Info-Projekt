package item

import( ."gfx2")

type data struct {
	done bool
	titel string
	deadline string
	x, y uint16
	hoehe uint16
	rh, gh, bh uint8
}

func New(titel, deadline string) *data {
	var f *data
	f = new(data)
	f.titel = titel
	f.deadline = deadline
	return f
}

func (f *data) Draw() {
	.
	

























































