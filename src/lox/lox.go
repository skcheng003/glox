package lox

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var hadError bool

type Lox struct {
	parser Parser
}

func NewLox() *Lox {
	return &Lox{}
}

func (l *Lox) run(str string) {
	scanner := NewScanner(str)
	tokens := scanner.ScanTokens()
	parser := NewParser(tokens)
	expr, _ := parser.Parse()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	astPrinter := NewAstPrinter()
	astPrinter.Print(expr)
}

func (l *Lox) RunFile(path string) {
	l.run(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	l.run(string(content))
	if hadError {
		os.Exit(65)
	}
}

func (l *Lox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _, _ := reader.ReadLine()
		hadError = false
		l.run(string(line))
	}
}

//
// func (lox *Lox) Error(line int, message string) {
// 	report(line, "", message)
// }

func report(line int, where, message string) {
	fmt.Println("[" + strconv.Itoa(line) + "] Error " + where + ": " + message)
}

func ErrorHandler(token *Token, message string) {
	if token.tokenType == EOF {
		report(token.line, "at end", message)
	} else {
		report(token.line, "at '"+token.Lexeme+"'", message)
	}
}
