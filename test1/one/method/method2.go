package main

const(
	WITE = iota
	BLUE
	WHITE
	GREEN
	YELLOW
	)

type Color byte

type Box struct {
	width,hight,depth float64
	color             Color
}
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width*b.hight*b.depth
}
func (b*Box) setColor(c Color)  {
	b.color = c
}
func (b1 BoxList) BiggestColour() Color {
	v :=0.00
	k := Color(WHITE)
	for _,b := range b1{
		if b.Volume() > v{
			v=b.Volume()
			k=b.color
		}
		return k
	}

}

