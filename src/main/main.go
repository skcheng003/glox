package main

import (
	"fmt"
	"os"
)

func main() {
	lox := NewLox()
	if len(os.Args) > 2 {
		fmt.Println("glox usage: [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		lox.runFile(os.Args[1])
	} else {
		lox.runPrompt()
	}
}
