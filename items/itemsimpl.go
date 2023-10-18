package items

import( ."gfx2")

type data struct {
	done bool
	titel string
	deadline string
	x, y uint16
	index int
	rh1, gh1, bh1 uint8
	rh2, gh2, bh2 uint8
	rh3, gh3, bh3 uint8
	highlight bool
}

func New(titel, deadline string) *data {
	var f *data
	f = new(data)
	f.titel = titel
	f.deadline = deadline
	return f
}

func (f *data) RückgabeDone() bool{
	return f.done
}

func (f *data) SwitchDone () {
			if f.done{
				f.done = false
			} else {
				f.done = true
			}
}

func (f *data) UmschreibenAlles (done bool, deadline, titel string) {
	f.done = done
	f.deadline = deadline
	f.titel = titel
}

func (f *data) UmschreibenTitel (titel string) {
	f.titel = titel
}

func (f *data) UmschreibenDeadline (deadline string) {
	f.deadline = deadline
}

func (f *data) UmschreibenDone (done bool) {
	f.done = done
}

func (f *data) SetzeWerte(x, y uint16, r, g, b, r2, g2, b2 uint8, index int, done bool, r3, g3, b3 uint8, highlight bool){
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
	f.rh3 = r3
	f.gh3 = g3
	f.bh3 = b3
	f.highlight = highlight
} 

func (f *data) String () string{
	var e string
	if f.done{
		e = e + "(x)"
	} else {
		e = e + "( )"
	}
	e = e + " " + f.titel + " " + f.deadline
	return e
}

func (f *data) Draw() {
	if f.highlight {
		Stiftfarbe (f.rh3, f.gh3, f.bh3)
	} else {
		if f.index%2==0 {
			Stiftfarbe(f.rh2, f.gh2, f.bh2)
		} else{
			Stiftfarbe(f.rh1, f.gh1, f.bh1)
		}
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
		if f.highlight{
			Stiftfarbe (f.rh3, f.gh3, f.bh3)
		}else{
			if f.index%2==0 {
				Stiftfarbe(f.rh2, f.gh2, f.bh2)
			} else{
				Stiftfarbe(f.rh1, f.gh1, f.bh1)
			}
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
		if f.highlight{
			Stiftfarbe (f.rh3, f.gh3, f.bh3)
		}else{
			if f.index%2==0 {
				Stiftfarbe(f.rh2, f.gh2, f.bh2)
			} else{
				Stiftfarbe(f.rh1, f.gh1, f.bh1)
			}
		}
		Vollkreis (55, f.y+40, 25)
	}
	SetzeFont("../fonts/LiberationMono-Bold.ttf", 30)
	Stiftfarbe(0, 0, 0)
	SchreibeFont(110, f.y+15, f.titel)
	SetzeFont("../fonts/LiberationMono-Bold.ttf", 18)
	SchreibeFont(110, f.y+55, f.deadline)
}

func (f *data) MausAufDone(mx, my uint16) bool{
	return (int(mx)-55)*(int(mx)-55)+(int(my)-int(f.y+40))*(int(my)-int(f.y+40))<=30*30
}


func (f *data) MausAufItem(my uint16) bool{
	return f.y<=my&&my<=f.y+80
}

func (f *data) SetzeHighlight(highlight bool) {
	f.highlight = highlight
}
	


	

























































