package app

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Game struct {
	Name      string
	Creator   string
	Version   string
	Desc      string
	Directory string
}

func LoadGame(directory string) Game {
	var game Game = Game{}
	// read content of game.toml in directory to string
	content, err := os.ReadFile(directory + "/game.toml")
	if err != nil {
		panic(err)
	}

	toml.Unmarshal(content, &game)
	if err != nil {
		panic(err)
	}

	fmt.Println("Game Name: ", game.Name)
	return game
}
