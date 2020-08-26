package main

import "fmt"

type person struct {
	name string
	age  int
}

type spy struct {
	person
	licence bool
}

func (p person) speak(someone string) string {
	return "Hello " + someone
}

func (p person) sayname() string {
	return p.name
}

func (s spy) speak(someone string) string {
	return "You are the target " + someone
}

func (s spy) sayname() string {
	return s.name
}

type speaker interface {
	speak(string) string
	sayname() string
}

// Comunicate is a function that takes an speker and a string and make the speaker talk to the string
func Comunicate(s speaker, some string) string {
	return s.sayname() + " - " + s.speak(some)
}

func main() {
	persons := []person{{"jhon", 20}, {"Juan", 33}}
	juan := spy{persons[1], true}

	speakers := []speaker{person{"laura", 21}, spy{person{"ahsley", 28}, true}}

	fmt.Println(persons[0].speak("Juan"))
	fmt.Println(juan.speak("Jhon"))
	fmt.Println(juan.name)
	fmt.Println(Comunicate(persons[0], juan.name))
	fmt.Println(Comunicate(juan, "jhon"))
	fmt.Println(Comunicate(speakers[0], speakers[1].sayname()))
	fmt.Println(Comunicate(speakers[1], speakers[0].sayname()))
}
