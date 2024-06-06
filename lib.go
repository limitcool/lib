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
