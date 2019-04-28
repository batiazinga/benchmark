package benchmark

import (
	"fmt"
	"testing"
)

// BenchmarkFillInCreate benchmarks creation + filling of slices and maps
// with various initializations.
func BenchmarkFillInCreate(b *testing.B) {
	// define some sizes
	sizes := []int{10, 100, 1000, 10000}

	// sub benchmards
	for _, size := range sizes {
		b.Run(fmt.Sprintf("slice default %v", size), func(b *testing.B) { fillInCreateSliceDefault(b, size) })
	}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("slice capacity %v", size), func(b *testing.B) { fillInCreateSliceCapacity(b, size) })
	}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("slice full %v", size), func(b *testing.B) { fillInCreateSliceFull(b, size) })
	}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("map default %v", size), func(b *testing.B) { fillInCreateMapDefault(b, size) })
	}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("map capacity %v", size), func(b *testing.B) { fillInCreateMapCapacity(b, size) })

	}
}

// sub benchmarks for slices
func fillInCreateSliceDefault(b *testing.B, size int) {
	for n := 0; n < b.N; n++ {
		// create default empty slice
		var slice []int
		// fill it
		for i := 0; i < size; i++ {
			slice = append(slice, i)
		}
	}
}
func fillInCreateSliceCapacity(b *testing.B, size int) {
	for n := 0; n < b.N; n++ {
		// create empty slice with the right capacity
		slice := make([]int, 0, size)
		// fill it
		for i := 0; i < size; i++ {
			slice = append(slice, i)
		}
	}
}
func fillInCreateSliceFull(b *testing.B, size int) {
	for n := 0; n < b.N; n++ {
		// create slice full of default zero values
		slice := make([]int, size)
		// set values
		for i := 0; i < size; i++ {
			slice[i] = i
		}
	}
}

// sub benchmarks for maps
func fillInCreateMapDefault(b *testing.B, size int) {
	for n := 0; n < b.N; n++ {
		// create default empty map
		m := make(map[int]int)
		// fill it
		for i := 0; i < size; i++ {
			m[i] = i
		}
	}
}
func fillInCreateMapCapacity(b *testing.B, size int) {
	for n := 0; n < b.N; n++ {
		// create empty map with right capacity
		m := make(map[int]int, size)
		// fill it
		for i := 0; i < size; i++ {
			m[i] = i
		}
	}
}
