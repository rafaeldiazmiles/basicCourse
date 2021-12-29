package main

import (
	"fmt"
	"strconv"
)

func main() {
	println("Becoming friends with Go.")
	friend := &GoFriend{
		name:       "Rafa",
		age:        36,
		motivation: "To overcome any challenge!",
		Project: Project{
			description: "I feel like Im the guy from 300 fighting against Perse, yes, Perse is Go Proffessional",
		},
	}

	enemy := &GoEnemy{
		name:       "Party Thor",
		intentions: "To battle Hydra (and also to ruin Marvel franchise with his jokes).",
		Project: Project{
			description: "Be dumb. Support 2021's collateral sexism and racism.",
		},
	}

	Print(friend)
	println("Let's see the actual project description...")
	fmt.Println(friend.GetProjectDescription())
	friend.ChangeProject("To catch 'em all!")
	println()
	println("Now check project description...")
	fmt.Println(friend.GetProjectDescription())
	println()
	println("An enemy appeared!")
	Print(enemy)
	println("I think Thor is re-calculating...")
	println("State your real intentions Party Thor!")
	enemy.ChangeProject("To parteee!")
	fmt.Println("Thor says: 'My intentions are...'", enemy.GetProjectDescription(), "'")

}

type Gophers interface {
	GetDescription() string
	ChangeProject(s string)
	GetProjectDescription() string
}

func Print(g Gophers) {
	description := g.GetDescription()
	fmt.Println(description)
}

type GoFriend struct {
	name       string
	age        int
	motivation string
	Project    Project
}

type Project struct {
	description string
}

func (g *GoFriend) GetDescription() string {
	returnString := "Name: " + g.name + ", Age: " + strconv.Itoa(g.age) + ", Motivation: " + g.motivation
	return returnString
}

func (f *GoFriend) GetProjectDescription() string {
	return f.Project.description
}

func (f *GoFriend) ChangeProject(s string) {
	f.Project.description = s
}

type GoEnemy struct {
	name       string
	intentions string
	Project    Project
}

func (g *GoEnemy) GetDescription() string {
	returnString := "Name: " + g.name + ", Intentions: " + g.intentions
	return returnString
}

func (f *GoEnemy) GetProjectDescription() string {
	return f.Project.description
}

func (f *GoEnemy) ChangeProject(s string) {
	f.Project.description = s
}
