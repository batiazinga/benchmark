package benchmark

import "testing"

func BenchmarkFillIn(b *testing.B) {
	// define some sizes
	hundred := 100
	tenThousand := 10000
	million := 1000000

	// sub benchmards
	// Fill in slices
	b.Run("slice default 100", func(b *testing.B) { fillInSliceDefault(b, hundred) })
	b.Run("slice default 10000", func(b *testing.B) { fillInSliceDefault(b, tenThousand) })
	b.Run("slice default 1000000", func(b *testing.B) { fillInSliceDefault(b, million) })

	b.Run("slice capacity 100", func(b *testing.B) { fillInSliceCapacity(b, hundred) })
	b.Run("slice capacity 10000", func(b *testing.B) { fillInSliceCapacity(b, tenThousand) })
	b.Run("slice capacity 1000000", func(b *testing.B) { fillInSliceCapacity(b, million) })

	b.Run("slice full 100", func(b *testing.B) { fillInSliceFull(b, hundred) })
	b.Run("slice full 10000", func(b *testing.B) { fillInSliceFull(b, tenThousand) })
	b.Run("slice full 1000000", func(b *testing.B) { fillInSliceFull(b, million) })

	b.Run("map default 100", func(b *testing.B) { fillInMapDefault(b, hundred) })
	b.Run("map default 10000", func(b *testing.B) { fillInMapDefault(b, tenThousand) })
	b.Run("map default 1000000", func(b *testing.B) { fillInMapDefault(b, million) })

	b.Run("map capacity 100", func(b *testing.B) { fillInMapCapacity(b, hundred) })
	b.Run("map capacity 10000", func(b *testing.B) { fillInMapCapacity(b, tenThousand) })
	b.Run("map capacity 1000000", func(b *testing.B) { fillInMapCapacity(b, million) })
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
