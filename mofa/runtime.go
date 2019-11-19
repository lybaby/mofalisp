package mofa

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func LoadFile(filepath string) string {
	code, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ""
	}
	return string(code)
}

var globalEnv *Env = NewEnv()

func car(node *AstNode) *AstNode {
	if node.car != nil {
		return node.car
	}
	panic("cannot car on empty node")
}

func cdr(node *AstNode) *AstNode {
	if node.cdr != nil {
		return node.cdr
	}
	panic("cannot cdr on empty node")
}

func cadr(node *AstNode) *AstNode {
	return car(cdr(node))
}

func caddr(node *AstNode) *AstNode {
	return car(cdr(cdr(node)))
}

func isEmpty(node *AstNode) bool {
	return node.car == nil && node.cdr == nil
}

func isSym(node *AstNode, sym string) bool {
	return node.nodeType == NTSYM && node.sval == sym
}

func progn(node *AstNode) *AstNode {
	cdrNode := cdr(node)
	var lastVal *AstNode = nil
	for !isEmpty(cdrNode) {
		carNode := car(cdrNode)
		lastVal = Eval(carNode)
		cdrNode = cdr(cdrNode)
	}
	if lastVal == nil {
		panic("progn require at lease one form")
	}
	return lastVal
}

func define(node *AstNode) *AstNode {
	varName := cadr(node)
	varVal := Eval(caddr(node))
	globalEnv.Setq(varName.sval, varVal)
	return &AstNode{nodeType: NTSEXPR}
}

func print(node *AstNode) *AstNode {
	cdrNode := cdr(node)
	strs:=make([]string, 0)
	for !isEmpty(cdrNode) {
		carNode := car(cdrNode)
		strs = append(strs, stringify(Eval(carNode)))
		cdrNode = cdr(cdrNode)
	}
	ret:= strings.Join(strs, " ")
	fmt.Println(ret)
	return &AstNode{nodeType:NTLITERAL, sval:ret}
}

func stringify(node *AstNode) string {
	switch node.nodeType {
	case NTNUM:
		return fmt.Sprintf("%g", node.nval)
	case NTSYM:
		return node.sval
	case NTSEXPR:
		return "<form>"
	}
	//todo
	return "<unsupport>"
}

func Eval(node *AstNode) *AstNode {
	switch node.nodeType {
	case NTNUM:
		return node
	case NTSYM:
		val, ok := globalEnv.Find(node.sval)
		if ok {
			return val
		} else {
			panic("undefined var:" + node.sval)
		}
	case NTSEXPR:
		if isEmpty(node) {
			panic("cannot eval empty node")
		}
		carNode := car(node)
		// temporary implementation
		if isSym(carNode, "progn") {
			return progn(node)
		} else if isSym(carNode, "define") {
			return define(node)
		} else if isSym(carNode, "print") {
			return print(node)
		} else {
			panic("not support yet")
		}
	default:
		panic("not support yet")
	}
}
