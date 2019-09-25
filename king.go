package main

import "fmt"

// King - everu kingdom has a king
type King struct {
	Name    string
	Kingdom *Kingdom
}

// newKing creates a new king
func newKing(name string) *King {
	k := &King{
		Name:    name,
		Kingdom: new(Kingdom),
	}

	return k
}

func (k *King) isRuler() bool {
	return len(k.Kingdom.Allies) >= 3
}

func (k *King) printIsRuler() {
	if k.isRuler() {
		fmt.Println(k.Name)
		return
	}
	fmt.Println("NONE")
}

func (k *King) getAllies() []*Kingdom {
	return k.Kingdom.Allies
}
