package main

import "fmt"

type TokenRecord struct {
	token_val  string
	token_type string
}

type LexicalTable struct {
	token_record TokenRecord
}

type Lexer struct {
	keywords      []string
	lexical_table LexicalTable
}

func (lex *Lexer) tokenize(input string) {
}
func (lex *Lexer) __init__(keywords []string) {
	lex.keywords = keywords
}

func main() {
	fmt.Println("Generating Lexical table...")
	var lexer_obj Lexer
	lexer_obj.tokenize("module lang")
}
