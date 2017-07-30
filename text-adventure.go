package main

import "fmt"
import "strings"

var (
	playerName string
	npcName string = "Ron"
)

// Helper Functions
func welcome () {
	fmt.Println(
		"################\n" +
		" Serious Police \n" +
		"################\n",
	)
}

func enterName() {
	fmt.Println("What is your name?")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("\nWelcome %s\n\n", input)
	playerName = input
}

func npcSays(speach string) {
	fmt.Printf("%s: \"%s\"\n", npcName, speach)
}

func npcShouts(speach string) {
	var shout = strings.ToUpper(speach)
	npcSays(shout)
}

func main ()  {
	welcome()
	enterName()
	npcShouts(playerName + " it's dangerous to go alone take this!")
}
