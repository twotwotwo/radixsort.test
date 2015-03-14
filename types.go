// Copyright 2009 The Go Authors.
// Copyright 2015 Randall Farmer.
// All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package radixsort sorts data using numeric, string, or []byte keys.
package radixsort

import (
	"bytes"
	"math"
	"sort" // for Interface
)

// NumberInterface represents a collection that can be sorted by a numeric
// (uint64) key.
type NumberInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Key provides the numeric (uint64) key for element i. To generate
	// uint64 keys for signed integer and floating-point numbers, look
	// at IntKey, Float64Key, and other similarly-named functions.
	Key(i int) uint64
}

// StringInterface represents a collection that can be sorted by a string
// key.
type StringInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Key provides the string key for element i.
	Key(i int) string
}

// BytesInterface represents a collection that can be sorted by a []byte
// key.
type BytesInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Key provides the []byte key for element i.
	Key(i int) []byte
}

const signBit = uint64(^uint(0) ^ (^uint(0))>>1) // works f/any int size

// IntKey generates a uint64 sort key from an int
func IntKey(i int) uint64 { return uint64(uint(i)) ^ signBit }

// Int32Key generates a uint64 sort key from an int32
func Int32Key(i int32) uint64 { return uint64(uint32(i)) ^ 1<<31 }

// Int64Key generates a uint64 key from an int64
func Int64Key(i int64) uint64 { return uint64(i) ^ 1<<63 }

// Float32Key generates a uint64 key from a float32. When sorting floats,
// consider comparing keys: that will make NaN values sort to the end,
// whereas with the float less-than operator they can cause panics.
func Float32Key(f float32) uint64 {
	b := uint64(math.Float32bits(f)) << 32
	b ^= ^(b>>63 - 1) | (1 << 63)
	return b
}

// Float64Key generates a uint64 key from a float64. When sorting floats,
// consider comparing keys: that will make NaN values sort to the end,
// whereas with the float less-than operator they can cause panics.
func Float64Key(f float64) uint64 {
	b := math.Float64bits(f)
	b ^= ^(b>>63 - 1) | (1 << 63)
	return b
}

// Flip reverses the order of items in a sort.Interface.
func Flip(data sort.Interface) {
	a, b := 0, data.Len()-1
	for b > a {
		data.Swap(a, b)
		a++
		b--
	}
}

// IntSlice attaches the methods of NumberInterface to []int, sorting in increasing order.
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p IntSlice) Key(i int) uint64   { return IntKey(p[i]) }

// Sort is a convenience method.
func (p IntSlice) Sort() { ByNumber(p) }

// Int32Slice attaches the methods of NumberInterface to []int32, sorting in increasing order.
type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int32Slice) Key(i int) uint64   { return Int32Key(p[i]) }

// Sort is a convenience method.
func (p Int32Slice) Sort() { ByNumber(p) }

// Int64Slice attaches the methods of NumberInterface to []int64, sorting in increasing order.
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int64Slice) Key(i int) uint64   { return Int64Key(p[i]) }

// Sort is a convenience method.
func (p Int64Slice) Sort() { ByNumber(p) }

// UintSlice attaches the methods of NumberInterface to []uint, sorting in increasing order.
type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p UintSlice) Key(i int) uint64   { return uint64(p[i]) }

// Sort is a convenience method.
func (p UintSlice) Sort() { ByNumber(p) }

// Uint32Slice attaches the methods of NumberInterface to []int32, sorting in increasing order.
type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint32Slice) Key(i int) uint64   { return uint64(p[i]) }

// Sort is a convenience method.
func (p Uint32Slice) Sort() { ByNumber(p) }

// Uint64Slice attaches the methods of NumberInterface to []uint64, sorting in increasing order.
type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint64Slice) Key(i int) uint64   { return p[i] }

// Sort is a convenience method.
func (p Uint64Slice) Sort() { ByNumber(p) }

// Float32Slice attaches the methods of NumberInterface to []uint32, sorting in increasing order, NaNs last.
type Float32Slice []float32

func (p Float32Slice) Len() int           { return len(p) }
func (p Float32Slice) Less(i, j int) bool { return Float32Key(p[i]) < Float32Key(p[j]) }
func (p Float32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Float32Slice) Key(i int) uint64   { return Float32Key(p[i]) }

// Sort is a convenience method.
func (p Float32Slice) Sort() { ByNumber(p) }

// Float64Slice attaches the methods of NumberInterface to []float64, sorting in increasing order, NaNs last.
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return Float64Key(p[i]) < Float64Key(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Float64Slice) Key(i int) uint64   { return Float64Key(p[i]) }

// Sort is a convenience method.
func (p Float64Slice) Sort() { ByNumber(p) }

// StringSlice attaches the methods of StringInterface to []string, sorting in increasing order.
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StringSlice) Key(i int) string   { return p[i] }

// Sort is a convenience method.
func (p StringSlice) Sort() { ByString(p) }

// BytesSlice attaches the methods of BytesInterface to [][]byte, sorting in increasing order.
type BytesSlice [][]byte

func (p BytesSlice) Len() int           { return len(p) }
func (p BytesSlice) Less(i, j int) bool { return bytes.Compare(p[i], p[j]) == -1 }
func (p BytesSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p BytesSlice) Key(i int) []byte   { return p[i] }

// Sort is a convenience method.
func (p BytesSlice) Sort() { ByBytes(p) }

// Ints sorts a slice of ints in increasing order.
func Ints(a []int) { IntSlice(a).Sort() }

// Int32s sorts a slice of int32s in increasing order.
func Int32s(a []int32) { Int32Slice(a).Sort() }

// Int64s sorts a slice of int64s in increasing order.
func Int64s(a []int64) { Int64Slice(a).Sort() }

// Uints sorts a slice of ints in increasing order.
func Uints(a []uint) { UintSlice(a).Sort() }

// Uint32s sorts a slice of uint64s in increasing order.
func Uint32s(a []uint32) { Uint32Slice(a).Sort() }

// Uint64s sorts a slice of uint64s in increasing order.
func Uint64s(a []uint64) { Uint64Slice(a).Sort() }

// Float32s sorts a slice of uint64s in increasing order, NaNs last.
func Float32s(a []float32) { Float32Slice(a).Sort() }

// Float64s sorts a slice of uint64s in increasing order, NaNs last.
func Float64s(a []float64) { Float64Slice(a).Sort() }

// Strings sorts a slice of strings in increasing order.
func Strings(a []string) { StringSlice(a).Sort() }

// Bytes sorts a slice of byte slices in increasing order.
func Bytes(a [][]byte) { BytesSlice(a).Sort() }

// IntsAreSorted tests whether a slice of ints is sorted in increasing order.
func IntsAreSorted(a []int) bool { return IsSorted(IntSlice(a)) }

// Int32sAreSorted tests whether a slice of int32s is sorted in increasing order.
func Int32sAreSorted(a []int32) bool { return IsSorted(Int32Slice(a)) }

// Int64sAreSorted tests whether a slice of int64s is sorted in increasing order.
func Int64sAreSorted(a []int64) bool { return IsSorted(Int64Slice(a)) }

// UintsAreSorted tests whether a slice of ints is sorted in increasing order.
func UintsAreSorted(a []uint) bool { return IsSorted(UintSlice(a)) }

// Uint32sAreSorted tests whether a slice of uint32s is sorted in increasing order.
func Uint32sAreSorted(a []uint32) bool { return IsSorted(Uint32Slice(a)) }

// Uint64sAreSorted tests whether a slice of uint64s is sorted in increasing order.
func Uint64sAreSorted(a []uint64) bool { return IsSorted(Uint64Slice(a)) }

// Float32sAreSorted tests whether a slice of float32s is sorted in increasing order, NaNs last.
func Float32sAreSorted(a []float32) bool { return IsSorted(Float32Slice(a)) }

// Float64sAreSorted tests whether a slice of float64s is sorted in increasing order, NaNs last.
func Float64sAreSorted(a []float64) bool { return IsSorted(Float64Slice(a)) }

// StringsAreSorted tests whether a slice of strings is sorted in increasing order.
func StringsAreSorted(a []string) bool { return IsSorted(StringSlice(a)) }

// BytesAreSorted tests whether a slice of byte slices is sorted in increasing order.
func BytesAreSorted(a [][]byte) bool { return IsSorted(BytesSlice(a)) }