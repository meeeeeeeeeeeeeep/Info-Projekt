package item

type Item interface {

 Draw()
 
 SetzeWerte(x, y uint16, r, g, b, r2, g2, b2 uint8, index int, done bool)
 
 }
