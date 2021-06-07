package main

import (
	"fmt"
	"os"

	"github.com/unsafecast/quack/ast"
	"github.com/unsafecast/quack/lexer"
	"github.com/unsafecast/quack/parser"
)

func main() {
	code, err := os.ReadFile("examples/something.qk")
	if err != nil {
		panic(err)
	}

	lexer := lexer.New([]rune(string(code)))
	parser := parser.New(lexer)
	node, _ := parser.Stmt()
	fmt.Println(ast.NodeToString(node))

	node, _ = parser.Assignment()
	fmt.Println(ast.NodeToString(node))
}
