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

// linear search ~ constant time amortized (I am dumb and might be wrong!)
func (lex *Lexer) isIn(char, category string) bool {
	for i := 0; i < len(category); i++ {
		if char == string(category[i]) {
			return true
		}
	}
	return false
}

// tokenize the input (very rough work)
func (lex *Lexer) tokenize(input string) {
	for lex.count < len(input) {
		if lex.removeComment(input) {
			continue
		} else if lex.extractWhiteSpace(input) {
			continue
		} else if lex.extractAlphanumerics(input) {
			continue
		} else {
			lex.extractSymbol(input)
		}

	}
}

// remove comment
func (lex *Lexer) removeComment(input string) bool {
	if lex.count >= len(input) || lex.count+1 >= len(input) {
		return false
	}
	// fmt.Println("Try... removeComment")
	index := lex.count
	if input[index:index+2] == "//" {
		index += 3
		for index < len(input) {
			if string(input[index]) != "\n" {
				index += 1
			} else {
				lex.count = index
				break
			}
		}
		return true
	}
	return false
}

// extract tokens
func (lex *Lexer) extractAlphanumerics(input string) bool {
	if lex.count >= len(input) {
		return false
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
	if temp_h != "" {
		lex.token_list = append(lex.token_list, temp_h)
		lex.count = index
		return true
	}
	return false
}

// extract whitespace they are also tokens but who cares...
func (lex *Lexer) extractWhiteSpace(input string) bool {
	if lex.count >= len(input) {
		return false
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
	if temp_h != "" {
		lex.token_list = append(lex.token_list, temp_h)
		lex.count = index
		return true
	}
	return false
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
	if temp_h != "" {
		lex.token_list = append(lex.token_list, temp_h)
	}
	index += 1
	lex.count = index
}

func main() {
	fmt.Println("Generating Lexical table...")
	var lexer_obj = Lexer{keywords: []string{"module", "grammar", "preference", "collaspe-w", "startsymbol"},
		alphabet: "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-",
		numerals: "0123456789"}
	lexer_obj.tokenize("//Expr.Add = [a + b]\nExpr.Sub = [a - b]\n")
	for i := 0; i < len(lexer_obj.token_list); i++ {
		fmt.Println(lexer_obj.token_list[i])
	}
	// lexer_obj.removeComment("model lang\nimport lang2")
	// lexer_obj.extractToken("module lang\nimport lang2")
}
