package main

import (
	"fmt"
	"strings"
)

// Kingdom struct defines informations for a particula kingdom
type Kingdom struct {
	Name   string
	Emblem string

	Allies []*Kingdom
}

func newKingdom(name string) *Kingdom {
	if emblem := emblem(name); emblem != "" {
		kingdom := &Kingdom{
			Name:   name,
			Emblem: emblem,
		}

		return kingdom
	}
	return nil
}

func (k *Kingdom) makeAlliance(name, msg string) {
	if !k.isAllie(name) {

		emblem := emblem(name)

		if makeAllie(emblem, msg) {
			k.Allies = append(k.Allies, newKingdom(name))
		}
	}
}

func (k *Kingdom) isAllie(name string) bool {
	for _, allie := range k.Allies {
		if strings.EqualFold(allie.Name, name) {
			return true
		}
	}
	return false
}

func makeAllie(emblem, msg string) bool {
	msgCount := countChar(msg)
	emblemCount := countChar(emblem)

	for key, value := range emblemCount {
		if msgCount[key] < value {
			return false
		}
	}
	return true
}

func (k *Kingdom) printAllies() {
	if len(k.Allies) == 0 {
		fmt.Println("NONE")
		return
	}

	for _, allie := range k.Allies {
		fmt.Printf("%s ", allie.Name)
	}

	fmt.Println()
}
