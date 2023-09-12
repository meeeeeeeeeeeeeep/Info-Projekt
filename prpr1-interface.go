package todoliste



// New(user, name string)


type ToDo interface {
	String () string
	NewItem (titel, deadline string, done bool)
	Umschreiben (i int, newdone bool, newtitel, newdeadline string)
	SwitchDone (i int)
}
