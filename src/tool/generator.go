package tool

import "fmt"

func GenerateAst(args ...string) {
	if len(args) != 1 {
		fmt.Println("Usage: generate_ast <output directory>")
		return
	}
	generateAst(args[0])
}

func generateAst(args string) {
	outputDir := args

}
