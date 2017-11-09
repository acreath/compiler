package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
	// "reflect"
)

// E  ->  T E'
// E' ->  + T E' | #
// T  ->  F T'
// T' ->  * F T' | #
// F  ->  ( E ) | i

var term_sym string = "i+*()"
// analysis table
var AT_col [] string = [] string {"E", "E'", "T", "T'", "F"}
var AT_row [] string = [] string {"i", "+", "*", "(", ")", "$"}
// i + * ( ) $
var AT_content [][]string = [][] string {
	{"TE'", "", "", "TE'", "", ""},
	{"", "+TE'", "", "", "#", "#"},
	{"FT'", "", "", "FT'", "", ""},
	{"", "#", "*FT'", "", "#", "#"},
	{"i", "", "", "(E)", "", ""},
}
var stack [] string

func main() {

	var content string
	var ip int = 0
	var x, a string

	b, err := ioutil.ReadFile("test.zhl")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	content = string(b)
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, " ", "", -1)
	fmt.Printf("input string: %v\n\n", content)
	stack = append(stack, "$", "E")
	x = "E"
	var count int = 0;
	for x != "$" {
		fmt.Printf("stack: %v\ninput: %v\n", stack, content[ip:])
		a = string(content[ip])
		if x == a {
			stack = stack[:len(stack)-1]
			fmt.Printf("match: %v\n", x)
			ip++
		} else if strings.ContainsAny(x, term_sym) {
			fmt.Printf("终结符\n")
			errorAction()
		} else if _, ok := checkTable(x, a); !ok {
			fmt.Printf("报错条目 %v\n", ok)
			errorAction()
		} else if s, ok := checkTable(x, a); ok {
			fmt.Printf("output: %v -> %v\n", x, s)
			stack = stack[:len(stack)-1]
			if s != "#" {
				// push
				myPush(s)
			}
		}
		x = string(stack[len(stack)-1])
		fmt.Printf("x: %v\n==============\n", x)
		
		// 调试
		count++
		if count > 100 {
			os.Exit(0)
		}
	}
}

/**
 * 报错条目 返回 false, 否则返回 产生式右部, true
 */
func checkTable(x string, a string) (string, bool) {
	var idx_col, idx_row int;
	for i := 0; i < len(AT_col); i++ {
		if AT_col[i] == x {
			idx_col = i
			break
		}
	}
	for i := 0; i < len(AT_row); i++ {
		if AT_row[i] == a {
			idx_row = i
			break
		}
	}
	ret := AT_content[idx_col][idx_row]
	// fmt.Printf("%d %d %v\n", idx_col, idx_row, ret)
	if ret != "" {
		return ret, true
	}

	return "", false
}

func errorAction() {
	fmt.Printf("error\n")
}

func myPush(s string) {
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if string(s[i]) == "'" {
			stack = append(stack, s[i-1:i+1])
			i--
		} else {
			stack = append(stack, string(s[i]))
		}
	}
	// fmt.Printf("stack: %v\n", stack)
}
