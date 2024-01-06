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
