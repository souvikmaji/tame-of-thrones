package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKing(t *testing.T) {
	tests := map[string]struct {
		input  string
		output *King
	}{
		"king shan": {
			input:  "shan",
			output: &King{Name: "shan", Kingdom: &Kingdom{}},
		},
		"king wth no name": {
			input:  "",
			output: &King{Name: "", Kingdom: &Kingdom{}},
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		assert.Equal(t, test.output, newKing(test.input))
	}
}

func TestIsRuler(t *testing.T) {
	tests := map[string]struct {
		input  *King
		output bool
	}{
		"ruler with 3 allies": {
			input: &King{
				Name:    "shan",
				Kingdom: &Kingdom{Name: "shans kingdom", Allies: []*Kingdom{newKingdom(Air), newKingdom(Water), newKingdom(Fire)}},
			},
			output: true,
		},
		"ruler with more than 3 allies": {
			input: &King{
				Name:    "shan",
				Kingdom: &Kingdom{Name: "shans kingdom", Allies: []*Kingdom{newKingdom(Air), newKingdom(Water), newKingdom(Fire), newKingdom(Land)}},
			},
			output: true,
		},
		"not a ruler with less than 3 allies": {
			input: &King{
				Name:    "shan",
				Kingdom: &Kingdom{Name: "shans kingdom", Allies: []*Kingdom{newKingdom(Air), newKingdom(Water)}},
			},
			output: false,
		},
		"not a ruler with 0 allie": {
			input: &King{
				Name:    "shan",
				Kingdom: &Kingdom{Name: "shans kingdom", Allies: []*Kingdom{}},
			},
			output: false,
		},
		"not a ruler with nil allies": {
			input: &King{
				Name:    "shan",
				Kingdom: &Kingdom{Name: "shans kingdom"},
			},
			output: false,
		},
		"not a ruler without a Kingdom": {
			input: &King{
				Name: "shan",
			},
			output: false,
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		assert.Equal(t, test.output, test.input.isRuler())
	}
}
