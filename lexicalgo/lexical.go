package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

var content string
var p int = 0
var keywords = map[string]int {
	"begin": 1, "if": 2, "then": 3, "while": 4, "do": 5, "end": 6,
	"+": 13, "-": 14, "*": 15, "/": 16, ":": 17, 
	":=": 18, "<": 20,"<>": 21, "<=": 22, ">": 23, 
	">=": 24, "=": 25, ";": 26, "(": 27,")": 28, "#": 0}
var letter string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digit string = "0123456789"

func main() {
	var reader *bufio.Reader
	var err error
	var syn int
	var token string
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please input:")
	content, err = reader.ReadString('#')
	if err != nil {
		fmt.Println("input error")
		os.Exit(0)
	}
	for p < len(content) {
		token, syn = lexical()
		fmt.Printf("(%d, %s)\n", syn, token)
	}
}

func lexical() (string, int) {
	c := content[p]
	right, left := p, p + 1
	var token string
	var syn int
	for c == ' ' || c == '\n' || c == '\t' {
		p += 1
		right = p
		c = content[p]
	}
	if strings.ContainsAny(string(c), letter) {
		for (c >= '0' && c <= '9') || strings.ContainsAny(string(c), letter) {
			left = p + 1
			p += 1
			c = content[p]
		}
		token = content[right:left]
		if syn, ok := keywords[token]; ok {
			return token, syn
		} else {
			return token, 10
		}

	} else if c >= '0' && c <= '9' {
		for c >= '0' && c <= '9' {
			left = p + 1
			p += 1
			c = content[p]
		}
		if strings.ContainsAny(string(c), letter) {
			for strings.ContainsAny(string(c), letter) {
				left = p + 1
				p += 1
				c = content[p]
			}
			token = content[right:left]
			return token, -1
		} else {
			token = content[right:left]
			return token, 11
		}

	} else {
		switch c {
		case ':':
			if content[p+1] == '=' {
				token = ":="
				p += 1
			} else {
				token = ":"
			}
		case '<':
			if content[p+1] == '>' || content[p+1] == '=' {
				token = content[p:p+2]
				p += 1
			} else {
				token = "<"
			}
		case '>':
			if content[p+1] == '=' {
				token = ">="
				p += 1
			} else {
				token = ">"
			}
		case '#':
			token = "#"
			syn = 0
			p += 1
			return token, syn
		default:
			token = string(c)
		}
		c = content[p + 1]
		if c != ' ' && !strings.ContainsAny(string(c), letter) && !(c >= '0' && c <= '9'){
			for c != ' ' && !strings.ContainsAny(string(c), letter) && !(c >= '0' && c <= '9') {
				left = p + 1
				token = content[right:left]
				p += 1
				c = content[p]
			}
		} else {
			p += 1
		}
		if _, ok := keywords[token]; ok {
			syn = keywords[token]
		} else {
			syn = -1
		}
		return token, syn
	}
} // func lexical

