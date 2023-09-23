package verification

import (
	"log"
	"sort"
)

var ServerNumbers []int

var Ticks uint32

func IsContinuous(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	// Sort the array first
	sort.Ints(arr)

	// Start with the first element
	prev := arr[0]

	for i := 1; i < len(arr); i++ {
		// Check if the current element is consecutive to the previous one
		if arr[i] != prev+1 {
			log.Printf("%d != %d", arr[i], prev+1)
			return false
		}
		prev = arr[i]
	}

	return true
}
