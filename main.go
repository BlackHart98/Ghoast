package main

import (
	"fmt"
)

type TokenRecord struct {
	token_val  string
	token_type string
}

type LexicalTable struct {
	token_record TokenRecord
}

type Lexer struct {
	keywords      []string
	alphabet      string
	numerals      string
	symbols       string
	lexical_table []LexicalTable
	count         int
	token_list    []string
}

// linear search ~ constant time amortized
func (lex *Lexer) isIn(char, category string) bool {
	for i := 0; i < len(category); i++ {
		if char == string(category[i]) {
			return true
		}
	}
	return false
}

// tokenize the input
func (lex *Lexer) tokenize(input string) {
	for lex.count < len(input) {
		lex.removeComment(input)
		lex.extractAlphanumerics(input)
		lex.extractWhiteSpace(input)
		lex.extractSymbol(input)
	}
}

// remove comment
func (lex *Lexer) removeComment(input string) {
	if lex.count >= len(input) || lex.count+1 >= len(input) {
		return
	}
	// fmt.Println("Try... removeComment")
	index := lex.count
	if input[index:index+2] == "//" {
		index += 3
		for index < len(input) {
			if string(input[index]) != "\n" {
				index += 1
			} else {
				break
			}
		}
	}
	lex.count = index
}

// extract tokens
func (lex *Lexer) extractAlphanumerics(input string) {
	if lex.count >= len(input) {
		return
	}
	temp_h := ""
	index := lex.count
	for index < len(input) {
		if lex.isIn(string(input[index]), lex.alphabet) ||
			lex.isIn(string(input[index]), lex.numerals) {
			temp_h += string(input[index])
			index += 1
		} else {
			break
		}
	}
	// fmt.Println(temp_h)
	lex.count = index
	// fmt.Println(lex.count)
}

// extract whitespace they are also tokens but who cares...
func (lex *Lexer) extractWhiteSpace(input string) {
	if lex.count >= len(input) {
		return
	}
	index := lex.count
	temp_h := ""
	for index < len(input) {
		if string(input[index]) == " " || string(input[index]) == "\n" || string(input[index]) == "\t" {
			temp_h += string(input[index])
			index += 1
		} else {
			break
		}
	}
	lex.count = index
}

func (lex *Lexer) extractSymbol(input string) {
	if lex.count >= len(input) {
		return
	}
	index := lex.count
	temp_h := ""
	if !lex.isIn(string(input[index]), lex.alphabet) &&
		!lex.isIn(string(input[index]), lex.alphabet) &&
		!(string(input[index]) == " ") &&
		!(string(input[index]) == "\n") &&
		!(string(input[index]) == "\t") {
		temp_h += string(input[index])
	}
	temp_h = ""
	index += 1
	lex.count = index
}

func main() {
	fmt.Println("Generating Lexical table...")
	var lexer_obj = Lexer{keywords: []string{"module", "grammar", "preference", "collaspe-w"},
		alphabet: "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-",
		numerals: "0123456789"}
	lexer_obj.tokenize("Expr.Add = [a + b]\n")
	// lexer_obj.removeComment("model lang\nimport lang2")
	// lexer_obj.extractToken("module lang\nimport lang2")
}
