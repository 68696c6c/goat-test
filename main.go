package main

import (
	"github.com/68696c6c/goat"
	"os"
	"goat-test/foo"
)

func main() {
	goat.Init()
	p := goat.GetProjectRoot()
	println("base path from main package: ", p)
	for _, p := range foo.FooPath() {
		println(p)
	}
	os.Exit(0)
}
