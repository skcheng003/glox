package lox

import (
	"bufio"
	"fmt"
	"glox/src/internal/scanner"
	"io"
	"log"
	"os"
	"strconv"
)

type Lox struct {
	hadError bool
}

func NewLox() *Lox {
	return &Lox{
		hadError: false,
	}
}

func (lox *Lox) run(str string) {
	scanner := scanner.NewScanner(str)
	scanner.ScanTokens()
}

func (lox *Lox) RunFile(path string) {
	lox.run(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	lox.run(string(content))
	if lox.hadError {
		os.Exit(65)
	}

}

func (lox *Lox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _, _ := reader.ReadLine()
		lox.run(string(line))
		lox.hadError = false
	}
}

func (lox *Lox) Error(line int, message string) {
	lox.report(line, "", message)
}

func (lox *Lox) report(line int, where, message string) {
	fmt.Println("[" + strconv.Itoa(line) + "] Error" + where + ": " + message)
	lox.hadError = true
}