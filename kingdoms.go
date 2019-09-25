package main

// Kingdom struct defines informations for a particula kingdom
type Kingdom struct {
	Name   string
	Emblem string

	Allies []*Kingdom
}

func (k *Kingdom) makeAlliance(name, msg string) {
	emblem := k.emblem(name)

	if isAllie(emblem, msg) {
		k.Allies = append(k.Allies, &Kingdom{Name: name, Emblem: emblem})
	}
}

func isAllie(emblem, msg string) bool {
	msgCount := countChar(msg)
	emblemCount := countChar(emblem)

	for key, value := range emblemCount {
		if msgCount[key] < value {
			return false
		}
	}
	return true
}

	}

	}

}
