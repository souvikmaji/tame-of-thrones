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
		k.Allies = append(k.Allies, new(Kingdom))
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

func countChar(msg string) map[rune]int {
	countMap := make(map[rune]int)
	for _, char := range msg {
		countMap[char]++
	}

	return countMap
}

// TODO kingdom factory
func (k Kingdom) emblem(name string) string {
	emblems := map[string]string{
		Land:  Panda,
		Water: Octopus,
		Ice:   Mammoth,
		Air:   Owl,
		Fire:  Dragon,
	}

	return emblems[name]
}
