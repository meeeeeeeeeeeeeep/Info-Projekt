package todoliste

import ("fmt"
		"item")

type data struct {
	user string
	name string
	it []item.Item
	casenum int
}

func New(user, name string) *data {
	var t *data
	t = new(data)
	t.user = user
	t.name = name
	return t
}

func (t *data) DrawListe() {
	
	
	
	
	
	SetzeFont("../fonts/LiberationMono-Bold.ttf", 30)
	Stiftfarbe(0, 0, 0)
	SchreibeFont(25, 10, f.name)

func (t *data) ChangeView(nom int) {
	t.casenum = nom
}

func (t *data) AddItem(it item.Item) {
	t.it = append(t.it, it)
}

func (t *data) String () string {
	var erg string
	erg = erg + fmt.Sprintln(t.name)
	erg = erg + fmt.Sprintln(t.user)
	switch t.casenum{
		case 0://alles eif iwie
			for i:=0;i<len(t.it);i++{
				erg = erg + fmt.Sprintln(t.it[i])
			}
		case 1: //alle false(unerledigt)
			for i:=0;i<len(t.it);i++{
				if t.it[i].RückgabeDone(){
				} else {
					erg = erg + fmt.Sprintln(t.it[i])
				}
			}
		case 2: //alle erledigten true
			for i:=0;i<len(t.it);i++{
				if t.it[i].RückgabeDone(){
					erg = erg + fmt.Sprintln(t.it[i])
				}
			}
		case 3://nach done geordnet
			var ndone []item.Item
			var ydone []item.Item
			for i:=0;i<len(t.it);i++{
				if t.it[i].RückgabeDone(){
					ydone = append(ydone, t.it[i])
				} else{
					ndone = append(ndone, t.it[i])
				}
			var itnew []item.Item
			itnew = append(itnew, ndone...)
			itnew = append(itnew, ydone...)
			t.it = itnew
			erg = erg + fmt.Sprintln(t.it[i])
		case 4://nach done geordnet 2
			var ndone []item.Item
			var ydone []item.Item
			for i:=0;i<len(t.it);i++{
				if t.it[i].RückgabeDone(){
					ydone = append(ydone, t.it[i])
				} else{
					ndone = append(ndone, t.it[i])
				}
			var itnew []item.Item
			itnew = append(itnew, ydone...)
			itnew = append(itnew, ndone...)
			t.it = itnew
			erg = erg + fmt.Sprintln(t.it[i])
		}
	}
	return erg
}		

func (t *data) SwitchDone (i int) {
	t.it[i].SwitchDone()
}

func (t *data) Umschreiben (i int, newdone bool, newtitel, newdeadline string){
	t.it[i].UmschreibenAlles(newdone, newdeadline, newtitel)
}

func (t* data) Löschen (i int) {
	if len(t.it)>i && i>=0 {
		var itnew []item.Item
		itnew = append(itnew, t.it[:i]...)
		itnew = append(itnew, t.it[i+1:]...)
		t.it = itnew
	}
}

func (t*data) UmschreibenTitel (i int, newtitel string){
	t.it[i].UmschreibenTitel(newtitel)
}

func (t*data) UmschreibenDeadline (i int, newdeadline string){
	t.it[i].UmschreibenDeadline(newdeadline)
}

func (t*data) UmschreibenDone (i int, newdone bool){
	t.it[i].UmschreibenDone(newdone)
}











































