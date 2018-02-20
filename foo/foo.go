package foo

import (
	"github.com/68696c6c/goat"
	"goat-test/bar"
)

func FooPath() []string {
	return []string{
		"base path from foo package: " + goat.GetProjectRoot(),
		bar.BarPath(),
	}
}
