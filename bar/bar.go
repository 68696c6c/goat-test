package bar

import (
	"github.com/68696c6c/goat"
)

func BarPath() string {
	return "base path from bar package: " + goat.GetProjectRoot()
}
