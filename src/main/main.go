package main

import (
	"fmt"
	"glox/src/lox"
	"os"
)

func main() {
	lox := lox.NewLox()
	fmt.Println(len(os.Args))
	if len(os.Args) > 2 {
		fmt.Println("glox usage: [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		lox.RunFile(os.Args[1])
	} else {
		lox.RunPrompt()
	}
}
