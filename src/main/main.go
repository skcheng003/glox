package main

import (
	"fmt"
	"glox/src/lox"
	"os"
)

func main() {
	interpreter := lox.NewLox()
	fmt.Println(len(os.Args))
	if len(os.Args) > 2 {
		fmt.Println("glox usage: [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		interpreter.RunFile(os.Args[1])
	} else {
		interpreter.RunPrompt()
	}
}
