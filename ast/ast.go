package ast

import "encoding/json"

type Node interface {
	NodeOffset() int64
}

func NodeToString(node Node) string {
	jsonForm, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(jsonForm)
}

type Ident struct {
	Value      string
	nodeOffset int64
}

func (ident *Ident) NodeOffset() int64 {
	return ident.nodeOffset
}

func NewIdent(value string, offset int64) Node {
	return &Ident{
		Value:      value,
		nodeOffset: offset,
	}
}

type NumLit struct {
	Value      float64
	nodeOffset int64
}

func (num *NumLit) NodeOffset() int64 {
	return num.nodeOffset
}

func NewNumLit(value float64, offset int64) Node {
	return &NumLit{
		Value:      value,
		nodeOffset: offset,
	}
}

type Assignment struct {
	Name  Node
	Value Node

	nodeOffset int64
}

func (ass *Assignment) NodeOffset() int64 {
	return ass.nodeOffset
}

func NewAssignment(name Node, value Node, offset int64) Node {
	return &Assignment{
		Name:  name,
		Value: value,

		nodeOffset: offset,
	}
}
