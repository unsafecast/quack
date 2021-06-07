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
	parser.Loop(func(node ast.Node, err error) {
		if node != nil {
			fmt.Println(ast.NodeToString(node))
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	})
}
