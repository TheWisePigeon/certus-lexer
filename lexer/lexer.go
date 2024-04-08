package lexer

import (
	"certus/tokens"
	"fmt"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}
	lexer.readChar()
	return lexer
}

func newToken(tokenKind tokens.TokenKind, literal string) tokens.Token {
	return tokens.Token{
		Kind:    tokenKind,
		Literal: literal,
	}
}

func isChar(character byte) bool {
	return character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' || character == '_'
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.currentChar = 0 //Because 0 is the ASCII representation of NUL
	} else {
		lexer.currentChar = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekAhead() {

}

func (lexer *Lexer) peekBehind() byte {
	if lexer.position == 0 {
		return 0
	}
	return lexer.input[lexer.position-1]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.currentChar == '\n' || lexer.currentChar == ' ' || lexer.currentChar == '\t' || lexer.currentChar == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readLine() {
	if lexer.currentChar == '#' {
		previousByte := lexer.peekBehind()
		if !isChar(previousByte) {
			for lexer.currentChar != '\n' {
				lexer.readChar()
			}
		}
		lexer.readChar()
		fmt.Printf("'%s'", string(lexer.currentChar))
	}
}

func (lexer *Lexer) NextToken() tokens.Token {
	var t tokens.Token
	lexer.skipWhitespace()
	lexer.readLine()
	switch lexer.currentChar {
	case 'G':
		fmt.Println("We in this shi")
	}
	lexer.readChar()
	return t
}
