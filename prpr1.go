package todoliste

import "fmt"

type data struct {
	user string
	name string
	titel []string
	done []bool
	deadline []string
	casenum int
}



func New(user, name string, casenum int) *data {
	var t *data
	t = new(data)
	t.user = user
	t.name = name
	t.casenum = casenum
	return t
}

func (t *data) String () string {
	var erg string
	erg = erg + fmt.Sprintln(t.name)
	erg = erg + fmt.Sprintln(t.user)
	switch t.casenum{
		case 0://alles eif iwie
			for i:=0;i<len(t.titel);i++{
				erg = erg + "("
				if t.done[i] {
					erg = erg + "x) "
				} else {
					erg = erg + " ) "
				}
				erg = erg + t.titel[i] + " "
				erg = erg + fmt.Sprintln(t.deadline[i])
			}
		case 1: //alle false(unerledigt)
			for i:=0;i<len(t.titel);i++{
				if t.done[i]{
				} else {
					erg = erg + "( ) " + t.titel[i] + " "
					erg = erg + fmt.Sprintln(t.deadline[i])
				}
			}
		case 2: //alle erledigten true
			for i:=0;i<len(t.titel);i++{
				if t.done[i] {
					erg = erg + "(x) " + t.titel[i]
					erg = erg + fmt.Sprintln(t.deadline[i])
				}
			}
		case 3: //alle geordnet nach done
			var unerledigtdone []bool
			var erledigtdone []bool
			var unerledigttitel []string
			var erledigttitel []string
			var unerledigtdeadline []string
			var erledigtdeadline []string
			for i:=0;i<len(t.titel);i++{
				if t.done[i] {
					erledigtdone = append(erledigtdone, t.done[i])
					erledigttitel = append(erledigttitel, t.titel[i])
					erledigtdeadline = append(erledigtdeadline, t.deadline[i])
				} else{
					unerledigtdone = append(unerledigtdone, t.done[i])
					unerledigttitel = append(unerledigttitel, t.titel[i])
					unerledigtdeadline = append(unerledigtdeadline, t.deadline[i])
				}
			}
			var donenew []bool
			var titelnew []string
			var deadlinenew []string
			donenew = append(donenew, unerledigtdone...)
			donenew = append(donenew, erledigtdone...)
			titelnew = append(titelnew, unerledigttitel...)
			titelnew = append(titelnew, erledigttitel...)
			deadlinenew = append(deadlinenew, unerledigtdeadline...)
			deadlinenew = append(deadlinenew, erledigtdeadline	...)
			t.done = donenew
			t.titel = titelnew
			t.deadline = deadlinenew
			for i:=0;i<len(t.titel);i++{
				erg = erg + "("
				if t.done[i] {
					erg = erg + "x) "
				} else {
					erg = erg + " ) "
				}
				erg = erg + t.titel[i] + " "
				erg = erg + fmt.Sprintln(t.deadline[i])
			}
	}
	return erg
}		




	
func (t *data) NewItem (titel, deadline string, done bool) {
	t.done = append(t.done, done)
	t.titel = append(t.titel, titel)
	t.deadline = append(t.deadline, deadline)
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
	if len(t.done)>i && i>=0 {
		var donenew []bool
		donenew = append(donenew, t.done[:i]...)
		donenew = append(donenew, t.done[i+1:]...)
		t.done = donenew
		var titelnew []string
		titelnew = append(titelnew, t.titel[:i]...)
		titelnew = append(titelnew, t.titel[i+1:]...)
		t.titel = titelnew
		var deadlinenew []string
		deadlinenew = append(deadlinenew, t.deadline[:i]...)
		deadlinenew = append(deadlinenew, t.deadline[i+1:]...)
		t.deadline = deadlinenew
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











































