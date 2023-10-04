package item

import( ."gfx2")

type data struct {
	done bool
	titel string
	deadline string
	x, y uint16
	index int
	rh1, gh1, bh1 uint8
	rh2, gh2, bh2 uint8
}

func New(titel, deadline string) *data {
	var f *data
	f = new(data)
	f.titel = titel
	f.deadline = deadline
	return f
}

func (f *data) RückgabeDone() bool{
	var d bool
	d = f.done
	return d
}

func (f *data) SwitchDone () {
			if f.done{
				f.done = false
			} else {
				f.done = true
			}
}

func (f *data) SetzeWerte(x, y uint16, r, g, b, r2, g2, b2 uint8, index int, done bool){
	f.done = done
	f.x = x
	f.y = y
	f.index = index
	f.rh1 = r
	f.gh1 = g
	f.bh1 = b
	f.rh2 = r2
	f.gh2 = g2
	f.bh2 = b2
} 

func (f *data) String () string{
	var e string
	if f.done{
		e = e + "(x)"
	} else {
		e = e + "( )"
	}
	e = e + " " + f.titel + " " + f.deadline
}

func (f *data) Draw() {
	if f.index%2==0 {
		Stiftfarbe(f.rh2, f.gh2, f.bh2)
	} else{
		Stiftfarbe(f.rh1, f.gh1, f.bh1)
	}
	x := Grafikspalten ()
	Vollrechteck(f.x, f.y, x, 80)
	if f.index%2!=0 {
		Stiftfarbe(f.rh2, f.gh2, f.bh2)
	} else{
		Stiftfarbe(f.rh1, f.gh1, f.bh1)
	}
	if f.done {
		Vollkreis (55, f.y+40, 30)
		if f.index%2==0 {
			Stiftfarbe(f.rh2, f.gh2, f.bh2)
		} else{
			Stiftfarbe(f.rh1, f.gh1, f.bh1)
		}
		Vollkreis (55, f.y+40, 25)
		if f.index%2!=0 {
			Stiftfarbe(f.rh2, f.gh2, f.bh2)
		} else{
			Stiftfarbe(f.rh1, f.gh1, f.bh1)
		}
		Vollkreis (55, f.y+40, 20)
	}else {
		Vollkreis (55, f.y+40, 30)
		if f.index%2==0 {
			Stiftfarbe(f.rh2, f.gh2, f.bh2)
		} else{
			Stiftfarbe(f.rh1, f.gh1, f.bh1)
		}
		Vollkreis (55, f.y+40, 25)
	}
	SetzeFont("../fonts/LiberationMono-Bold.ttf", 30)
	Stiftfarbe(0, 0, 0)
	SchreibeFont(110, f.y+15, f.titel)
	SetzeFont("../fonts/LiberationMono-Bold.ttf", 18)
	SchreibeFont(110, f.y+55, f.deadline)
}

	

























































