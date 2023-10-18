package todoliste

// New(user, name string)
import "items"

type ToDo interface {
	//String () string
	Umschreiben (i int, newdone bool, newtitel, newdeadline string)
	SwitchDone (i int)
	DrawListe()
	Löschen (i int)
	UmschreibenTitel (i int, newtitel string)
	UmschreibenDeadline (i int, newdeadline string)
	ChangeView(nom int)
	AddItem(it items.Item)
	UmschreibenDone (i int, newdone bool)
	SetzeFarben( r, g, b, r2, g2, b2, r3, g3, b3 uint8)
	UpdateList ()
}




















