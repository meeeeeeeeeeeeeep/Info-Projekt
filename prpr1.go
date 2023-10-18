package todoliste

import (//"fmt"
		"items"
		. "gfx2")

type data struct {
	user string
	name string
	it []items.Item
	casenum int
	rh1, gh1, bh1 uint8
	rh2, gh2, bh2 uint8
	rh3, gh3, bh3 uint8
	highlight int //index des gehighlighteten Items ; -1 ohne
}

func New(user, name string) *data {
	var t *data
	t = new(data)
	t.user = user
	t.name = name
	t.highlight = -1
	return t
}

func (t *data) DrawListe() {
	Stiftfarbe(t.rh1, t.gh1, t.bh1)
	x := Grafikspalten ()
	Vollrechteck(0,0, x, 100)
	SetzeFont("../fonts/LiberationMono-Bold.ttf", 30)
	Stiftfarbe(0, 0, 0)
	SchreibeFont(25, 10, t.name)
	SchreibeFont(25, 50, t.user)
	for i:=0;i<len(t.it);i++{
		t.it[i].Draw()
	}
}

func (t *data) UpdateList () {
	for i:=0; i<len(t.it); i++{
		t.it[i].SetzeWerte(0, uint16(100+i*80) , t.rh1, t.gh1, t.bh1, t.rh2, t.gh2, t.bh2 , i, t.it[i].RückgabeDone(), t.rh3, t.gh3, t.bh3 , t.highlight==i)
	}
}

func (t *data) ChangeView(nom int) {
	t.casenum = nom
}

func (t *data) AddItem(it items.Item) {
	t.it = append(t.it, it)
}

//func (t *data) Highlighter(mx, my uint16) {
	

func (t *data) SetzeFarben( r, g, b, r2, g2, b2, r3, g3, b3 uint8){
	t.rh1 = r
	t.gh1 = g
	t.bh1 = b
	t.rh2 = r2
	t.gh2 = g2
	t.bh2 = b2
	t.rh3 = r3
	t.gh3 = g3
	t.bh3 = b3
} 

//~ func (t *data) String () string {
	//~ var erg string
	//~ erg = erg + fmt.Sprintln(t.name)
	//~ erg = erg + fmt.Sprintln(t.user)
	//~ switch t.casenum{
		//~ case 0://alles eif iwie
			//~ for i:=0;i<len(t.it);i++{
				//~ erg = erg + fmt.Sprintln(t.it[i])
			//~ }
		//~ case 1: //alle false(unerledigt)
			//~ for i:=0;i<len(t.it);i++{
				//~ if t.it[i].RückgabeDone(){
				//~ } else {
					//~ erg = erg + fmt.Sprintln(t.it[i])
				//~ }
			//~ }
		//~ case 2: //alle erledigten true
			//~ for i:=0;i<len(t.it);i++{
				//~ if t.it[i].RückgabeDone(){
					//~ erg = erg + fmt.Sprintln(t.it[i])
				//~ }
			//~ }
		//~ case 3://nach done geordnet
			//~ var ndone []item.Item
			//~ var ydone []item.Item
			//~ for i:=0;i<len(t.it);i++{
				//~ if t.it[i].RückgabeDone(){
					//~ ydone = append(ydone, t.it[i])
				//~ } else{
					//~ ndone = append(ndone, t.it[i])
				//~ }
			//~ }
			//~ var itnew []item.Item
			//~ itnew = append(itnew, ndone...)
			//~ itnew = append(itnew, ydone...)
			//~ t.it = itnew
			//~ erg = erg + fmt.Sprintln(t.it[i])
		//~ case 4://nach done geordnet 2
			//~ var ndone []item.Item
			//~ var ydone []item.Item
			//~ for i:=0;i<len(t.it);i++{
				//~ if t.it[i].RückgabeDone(){
					//~ ydone = append(ydone, t.it[i])
				//~ } else{
					//~ ndone = append(ndone, t.it[i])
				//~ }
			//~ var itnew []item.Item
			//~ itnew = append(itnew, ydone...)
			//~ itnew = append(itnew, ndone...)
			//~ t.it = itnew
			//~ erg = erg + fmt.Sprintln(t.it[i])
		//~ }
	//~ }

	//~ return erg
//~ }		

func (t *data) SwitchDone (i int) {
	t.it[i].SwitchDone()
}

func (t *data) Umschreiben (i int, newdone bool, newtitel, newdeadline string){
	t.it[i].UmschreibenAlles(newdone, newdeadline, newtitel)
}

func (t* data) Löschen (i int) {
	if len(t.it)>i && i>=0 {
		var itnew []items.Item
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











































