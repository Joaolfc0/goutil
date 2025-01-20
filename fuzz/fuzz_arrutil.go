//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"fmt"

	"github.com/gookit/goutil/arrutil"
)

// FuzzArrUtil testa as funções no pacote arrutil.
func FuzzArrUtil(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	// Convert data to a slice of integers for testing purposes.
	intData := make([]int, len(data))
	for i, b := range data {
		intData[i] = int(b)
	}

	// Testa Reverse
	reversed := make([]int, len(intData))
	copy(reversed, intData)
	arrutil.Reverse(reversed)
	if len(reversed) != len(intData) {
		panic(fmt.Sprintf("Reverse altered the length of the slice: expected %d, got %d", len(intData), len(reversed)))
	}

	// Testa Remove
	if len(intData) > 0 {
		removed := arrutil.Remove(intData, intData[0])
		if len(removed) > len(intData) {
			panic("Remove increased the size of the slice")
		}
	}

	// Testa Filter
	filtered := arrutil.Filter(intData, func(el int) bool {
		return el%2 == 0 // Filtra números pares
	})
	for _, v := range filtered {
		if v%2 != 0 {
			panic(fmt.Sprintf("Filter failed: found odd number %d in filtered slice", v))
		}
	}

	// Testa Map
	mapped := arrutil.Map(intData, func(input int) (int, bool) {
		return input * 2, true // Dobra os valores
	})
	for i, v := range mapped {
		if v != intData[i]*2 {
			panic(fmt.Sprintf("Map failed: expected %d, got %d", intData[i]*2, v))
		}
	}

	// Testa Unique
	unique := arrutil.Unique(intData)
	seen := make(map[int]bool)
	for _, v := range unique {
		if seen[v] {
			panic(fmt.Sprintf("Unique failed: duplicate value %d found", v))
		}
		seen[v] = true
	}

	// Testa IndexOf
	if len(intData) > 0 {
		idx := arrutil.IndexOf(intData[0], intData)
		if idx < 0 || idx >= len(intData) {
			panic(fmt.Sprintf("IndexOf failed: value %d not found", intData[0]))
		}
	}

	return 1
}
