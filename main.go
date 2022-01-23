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
	lexical_table []LexicalTable
	count         int
}

// tokenize the input
func (lex *Lexer) tokenize(input string) {
}

// remove comment
func (lex *Lexer) removeComment(input string) {
	if lex.count >= len(input) && lex.count+1 >= len(input) {
		return
	}
	index := lex.count
	// fmt.Println(input[index : index+2])
	if input[index:index+2] == "//" {
		lex.count = index + 3
		for lex.count < len(input) {
			if string(input[index]) != "\n" {

				index += 1
			} else {
				fmt.Println("Debugging")
				break
			}
		}
	}
	lex.count = index
	fmt.Println(lex.count)
}

// constructor
func (lex *Lexer) __init__(keywords []string) {
	lex.keywords = keywords
}

func main() {
	fmt.Println("Generating Lexical table...")
	var lexer_obj Lexer
	lexer_obj.__init__([]string{"module", "grammar"})
	// lexer_obj.tokenize("//module lang")
	lexer_obj.removeComment("//model lang\nimport lang2")
}
