package item

type Item interface {

 Draw()
 
 SetzeWerte(x, y uint16, r, g, b, r2, g2, b2 uint8, index int, done bool)
 
 New(titel, deadline string)
 
 RückgabeDone() bool
 
 SwitchDone ()
 
 String () string
 
 UmschreibenAlles (done bool, deadline, titel string)
 
 UmschreibenTitel (titel string)
 
 UmschreibenDeadline (deadline string)
 
 UmschreibenDone (done bool)
 
 MausAufDone(mx, my uint16) bool
 
 MausAufItem(my uint16) bool
 
 }
