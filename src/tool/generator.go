package tool

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GenerateAst(args ...string) {
	if len(args) != 1 {
		fmt.Println("Usage: generate_ast <output directory>")
		os.Exit(64)
	}
	astTypes := []string{
		"Binary : Left Expr, Operator *Token, Right Expr",
		"Grouping: Expr Expr",
		"Literal : Value any",
		"Unary : Operator *Token, Right Expr",
	}

	generateAst(args[0], "expr", astTypes)
}

func generateAst(dir, filename string, astTypes []string) {
	path := dir + "/" + strings.ToLower(filename) + ".go"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file " + path + " failed. Program exits.")
		os.Exit(65)
	}
	defer file.Close()

	_ = file.Truncate(0)

	writer := bufio.NewWriter(file)

	defineAstBody(writer, astTypes)

	_ = writer.Flush()
}

func defineInterface(astTypes []string) string {
	var types []string
	for _, elem := range astTypes {
		types = append(types, strings.Trim(strings.Split(elem, ":")[0], " "))
	}
	interfaceDef := "type Visitor interface {\n"
	for _, elem := range types {
		interfaceDef += "\tVisit" + elem + "(" + strings.ToLower(elem) + " *" + elem + ") any\n"
	}
	interfaceDef += "}\n\n"

	interfaceDef += "type Expr interface {\n"
	interfaceDef += "\taccept(visitor Visitor) any\n"
	interfaceDef += "}\n\n"
	return interfaceDef
}

func defineAstBody(writer *bufio.Writer, astTypes []string) {
	_, _ = writer.WriteString("package lox\n\n")

	visitorDef := defineInterface(astTypes)

	_, _ = writer.WriteString(visitorDef)

	for _, elem := range astTypes {
		structDef, name, fields := defineStruct(elem)
		funcDef := defineFunc(name, fields)
		_, _ = writer.WriteString(structDef + funcDef)
	}
}

func defineStruct(elem string) (string, string, []string) {
	name := strings.Trim(strings.Split(elem, ":")[0], " ")
	fields := strings.Split(strings.Split(elem, ":")[1], ",")
	for idx, field := range fields {
		fields[idx] = strings.Trim(field, " ")
	}
	structDef := "type " + name + " struct {\n"
	for _, field := range fields {
		structDef += "\t" + strings.Trim(field, " ") + "\n"
	}
	structDef += "}\n\n"
	return structDef, name, fields
}

func defineFunc(name string, fields []string) string {
	var paraText string
	var paraLists []string
	for _, field := range fields {
		paraName := strings.Split(field, " ")[0]
		paraType := strings.Split(field, " ")[1]
		paraText += strings.ToLower(paraName) + " "
		paraText += paraType + ", "
		paraLists = append(paraLists, paraName)
	}
	paraText = strings.TrimRight(paraText, ", ")
	funcDef := "func New" + name + "(" + paraText + ") " + "*" + name + " {\n"
	funcDef += "\treturn &" + name + "{\n"
	for _, para := range paraLists {
		funcDef += "\t\t" + para + ": " + strings.ToLower(para) + ",\n"
	}
	funcDef += "\t}\n"
	funcDef += "}\n\n"

	funcDef += "func (" + strings.ToLower(name) + " *" + name + ") accept(visitor Visitor) any {\n"
	funcDef += "\treturn visitor.Visit" + name + "(" + strings.ToLower(name) + ")\n"
	funcDef += "}\n\n"
	return funcDef
}
