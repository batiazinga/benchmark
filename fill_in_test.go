package benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkFillIn(b *testing.B) {
	// define some sizes
	sizes := []int{100, 10000, 1000000}

	// sub benchmards
	// Fill in slices
	for _, size := range sizes {
		b.Run(fmt.Sprintf("slice default %v", size), func(b *testing.B) { fillInSliceDefault(b, size) })
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("slice capacity %v", size), func(b *testing.B) { fillInSliceCapacity(b, size) })
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("slice full %v", size), func(b *testing.B) { fillInSliceFull(b, size) })
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("map default %v", size), func(b *testing.B) { fillInMapDefault(b, size) })
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("map capacity %v", size), func(b *testing.B) { fillInMapCapacity(b, size) })

	}
}

// sub benchmarks for slices
func fillInSliceDefault(b *testing.B, size int) {
	var slice []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// fill in slice
		for i := 0; i < size; i++ {
			slice = append(slice, i)
		}
	}
}
func fillInSliceCapacity(b *testing.B, size int) {
	slice := make([]int, 0, size)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			slice = append(slice, i)
		}
	}
}
func fillInSliceFull(b *testing.B, size int) {
	slice := make([]int, size)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			slice[i] = i
		}
	}
}

// sub benchmarks for maps
func fillInMapDefault(b *testing.B, size int) {
	m := make(map[int]int)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			m[i] = i
		}
	}
}
func fillInMapCapacity(b *testing.B, size int) {
	m := make(map[int]int, size)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			m[i] = i
		}
	}
}
