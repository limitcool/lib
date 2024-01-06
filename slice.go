package lib

import "golang.org/x/exp/constraints"

func InSlice[T comparable](element T, Slice []T) bool {
	for _, v := range Slice {
		if element == v {
			return true
		}
	}
	return false
}

func Unique[T comparable](slice []T) []T {
	// 创建一个临时map用来存储数组元素
	temp := make(map[T]struct{})
	index := 0
	// 将元素放入map中
	for _, v := range slice {
		temp[v] = struct{}{}
	}
	tempSlice := make([]T, len(temp))
	for key := range temp {
		tempSlice[index] = key
		index++
	}
	return tempSlice
}

func Push[T any](slice []T, element T) []T {
	slice = append(slice, element)
	return slice
}

func Pop[T any](slice []T) []T {
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}
	return slice
}

func PopFront[T any](slice []T) []T {
	if len(slice) > 0 {
		slice = slice[1:]
	}
	return slice
}

func Exclude[T comparable](slice []T, element T) []T {
	result := slice[:0]
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}

func Reverse[T any](s []T) []T {
	h := make([]T, len(s))
	for i := 0; i < len(s); i++ {
		h[i] = s[len(s)-i-1]
	}
	return h
}

// Keys returns a slice of the keys of the map. based with go 1.18 generics
func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func IntToBool[T constraints.Signed](a T) bool {
	switch a {
	case 0, -1:
		return false
	}
	return true
}
