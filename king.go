package main

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

func (k *King) getAllies() []*Kingdom {
	return k.Kingdom.Allies
}
