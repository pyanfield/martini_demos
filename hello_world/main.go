package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "hello world!"
	})
	m.Run()
}
