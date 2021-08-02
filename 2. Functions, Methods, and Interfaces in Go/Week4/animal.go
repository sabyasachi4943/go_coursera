package main

import (
	"fmt"
)

type animal struct {
	food, locomotion, noise string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (ani animal) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani animal) Move() {
	fmt.Println(ani.locomotion)
	return
}

func (ani animal) Speak() {
	fmt.Println(ani.noise)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}
	var generalAni animalInterface
	for {
		var command, requestAni, requestType string
		fmt.Println(">")
		fmt.Scan(&command, &requestAni, &requestType)
		if command == "query" {
			generalAni = animalMap[requestAni]
			switch requestType {
			case "eat":
				generalAni.Eat()
			case "move":
				generalAni.Move()
			case "speak":
				generalAni.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[requestAni] = animalMap[requestType]
			fmt.Println("Created it!")
		}
	}
}