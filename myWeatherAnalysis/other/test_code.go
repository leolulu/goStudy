package other

import "fmt"

type P1 struct {
	name string
	P2
}

func (p P1) SayP1Name() {
	fmt.Println(p.name)
}

type P2 struct {
	name string
}

func (p P2) SayP2Name() {
	fmt.Println(p.name)
}

func (p P1) SayP2NameFromP1() {
	p.SayP2Name()
}
