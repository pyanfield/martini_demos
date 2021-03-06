package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "hello world!"
	})
	m.Get("/hello/:name", func(params martini.Params) string {
		return "hello " + params["name"]
	})

	m.Run()
}
