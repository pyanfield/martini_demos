package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type MyString interface{}

func Hello(name string, word MyString, number int) {
	fmt.Printf("%s said %s %d times \n", name, word, number)
}

type MyStruct struct {
	Name     string `inject`
	Team     string
	Location MyString `inject`
}

func main() {
	ij := inject.New()
	ij.Map(20)
	ij.MapTo("Good Morning", (*MyString)(nil))

	ij1 := inject.New()
	ij1.Map("Steven")

	ij.SetParent(ij1)
	ij.Invoke(Hello)

	fmt.Println(inject.InterfaceOf((*MyString)(nil)))

	team := MyStruct{}
	ij2 := inject.New()
	ij2.Map("Liverpool LFC")
	ij2.MapTo("England", (*MyString)(nil))
	ij2.Apply(&team)

	fmt.Printf("TEAM NAME: %s \n", team.Name)
	fmt.Printf("TEAM TEAM: %s \n", team.Team)
	fmt.Printf("TEAM LOCATION: %s \n", team.Location)
}
