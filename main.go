package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/elvodqa/wright/app"
	"github.com/elvodqa/wright/app/ast"
)

var (
	GameDirectory string = "."
	ReplMode      string = "lexer"
)

func init() {
	fmt.Println("Wright v0.0.1")
	fmt.Println("------------------")
	flag.StringVar(&GameDirectory, "game", "", "Game directory")
	flag.StringVar(&ReplMode, "repl", "", "Lexer REPL mode")
	flag.Parse()
}

func main() {
	switch ReplMode {
	case "lexer":
		fmt.Println("Lexer REPL mode")
		whitespace_visible := false
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')
			if text == "?exit\n" || text == "?quit\n" {
				// break
				return
			}
			if text == "##clear\n" || text == "clear\n" {
				fmt.Print("\033[H\033[2J")
				continue
			}
			if text == "##ws\n" {
				whitespace_visible = !whitespace_visible
				fmt.Printf("Whitespace visible: %t\n", whitespace_visible)
				continue
			}
			l := ast.NewLexer(text)
			for tok := l.NextToken(); tok.Type != ast.EOF; tok = l.NextToken() {
				if whitespace_visible {
					fmt.Printf("\033[36m%+v\033[0m\n", tok)
					//fmt.Printf("%+v\n", tok)
				} else {
					if tok.Type != "WHITESPACE" {
						fmt.Printf("\033[36m%+v\033[0m\n", tok)
					}
				}
			}
		}
	case "parser":
		break
	}

	if GameDirectory == "" {
		flag.PrintDefaults()
		return
	}

	game := app.LoadGame(GameDirectory)
	app.Run(game)

}
