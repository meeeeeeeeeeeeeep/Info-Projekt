package todoliste

import ("fmt"
		"item")

type data struct {
	user string
	name string
	//titel []string
	//done []bool
	//deadline []string
	it []item.Item
	casenum int
}

func New(user, name string) *data {
	var t *data
	t = new(data)
	t.user = user
	t.name = name
	//t.casenum = 0
	return t
}



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
		if len(t.done)>i && i>=0 {
			if t.done[i]{
				t.done[i] = false
			} else {
				t.done[i] = true
			}
		}
}

func (t *data) Umschreiben (i int, newdone bool, newtitel, newdeadline string){
	if len(t.done)>i && i>=0 {
			t.titel[i]=newtitel
			t.done[i] = newdone
			t.deadline[i] = newdeadline
		}
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
	if len(t.done)>i && i>=0 {
		t.titel[i] = newtitel
	}
}

func (t*data) UmschreibenDeadline (i int, newdeadline string){
	if len(t.done)>i && i>=0 {
		t.deadline[i] = newdeadline
	}
}











































