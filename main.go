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
	alphabet      string
	numerals      string
	symbol        string
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
	// var token_list []string

	// states := [4]string{"ALPHANUMERICS*", "WHITESPACE", "SYMBOLS", "NIL"}
	// state := "NIL"
	// temp_h := ""
	for lex.count < len(input) {
		lex.removeComment(input)
	}
}

// remove comment
func (lex *Lexer) removeComment(input string) {
	if lex.count >= len(input) && lex.count+1 >= len(input) {
		return
	}
	index := lex.count
	if input[index:index+2] == "//" {
		index += 3
		lex.count += index
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
func (lex *Lexer) extractToken(input string) {

}

//extract whitespace
// func (lex *Lexer) extractWhitespace(input string) {
// 	if lex.count >= len(input) {
// 		return
// 	}
// 	index := lex.count
// 	temp_h := ""
// 	if lex.isIn(string(input[index]), lex.alphabet) {
// 		temp_h += string(input[index])
// 		index += 1
// 	}
// 	if lex.count >= len(input) {
// 		return
// 	}
// 	for index < len(input) {
// 		if string(input[index]) == "\t" {
// 			temp_h += string(input[index])
// 			index += 1
// 		} else if string(input[index]) == "\n" {
// 			temp_h += string(input[index])
// 			index += 1
// 		} else if string(input[index]) == " " {
// 			temp_h += string(input[index])
// 			index += 1
// 		} else {
// 			break
// 		}
// 	}
// 	lex.count = index
// }

// constructor
func (lex *Lexer) __init__(keywords []string, alphabet, numerals string) {
	lex.keywords = keywords
}

func main() {
	fmt.Println("Generating Lexical table...")
	var lexer_obj Lexer
	lexer_obj.__init__([]string{"module", "grammar"}, "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "0123456789")
	lexer_obj.tokenize("module\n")
	// lexer_obj.removeComment("model lang\nimport lang2")
	// lexer_obj.extractToken("module lang\nimport lang2")
}
