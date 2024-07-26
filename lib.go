package lib

import (
	"os"
	"strings"
)

func SetDebugMode(debugFunction func()) {
	if strings.ToLower(os.Getenv("DEBUG_MODE")) == "true" {
		debugFunction()
	}
}

func Zero[T any]() T {
	var v T
	return v
}

func Ptr[T any](t T) *T { return &t }

func ToInterfaceSlice[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
