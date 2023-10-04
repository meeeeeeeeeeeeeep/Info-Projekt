package todoliste



// New(user, name string)


type ToDo interface {
	String () string
	Umschreiben (i int, newdone bool, newtitel, newdeadline string)
	SwitchDone (i int)
	Löschen (i int)
	UmschreibenTitel (i int, newtitel string)
	UmschreibenDeadline (i int, newdeadline string)
	ChangeView(nom int)
	AddItem(it item.Item)
}
