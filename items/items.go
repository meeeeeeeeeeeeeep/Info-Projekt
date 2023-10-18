package items

type Item interface {

 Draw()
 
 SetzeWerte(x, y uint16, r, g, b, r2, g2, b2 uint8, index int, done bool, r3, g3, b3 uint8, highlight bool)
  
 SetzeHighlight(highlight bool)
 
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
