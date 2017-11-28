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
	{"TE'", "", "", "TE'", "synch", "synch"},
	{"", "+TE'", "", "", "#", "#"},
	{"FT'", "synch", "", "FT'", "synch", "synch"},
	{"", "#", "*FT'", "", "#", "#"},
	{"i", "synch", "", "(E)", "synch", "synch"},
}
var stack [] string

func main() {

	var content string
	var ip int = 0
	var x, a string
	var faultIndex []int
	var flag bool = true

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
		// fmt.Printf("stack: %v\ninput: %v\n", stack, content[ip:])
		fmt.Printf("%-16v %-16v", strings.Join(stack[:], " "), content[ip:])
		a = string(content[ip])
		if x == a {
			stack = stack[:len(stack)-1]
			// fmt.Printf("match: %v\n", x)
			out := "match " + x
			fmt.Printf("%-16s", out)
			ip++
		} else if strings.ContainsAny(x, term_sym) {
			fmt.Printf("终结符")
			stack = stack[:len(stack) - 1]
			flag = errorAction()
		} else if s, ok := checkTable(x, a); !ok {
			// fmt.Printf("报错条目")
			if s == "" {
				out := "ignore " + string(content[ip])
				fmt.Printf("%c[1;40;32m%-16s%c[0m", 0x1B, out, 0x1B)
				faultIndex = append(faultIndex, ip)
				ip++
			} else if s == "synch" {
				out := "=synch"
				fmt.Printf("%c[1;40;32m%-16s%c[0m", 0x1B, out, 0x1B)				
				stack = stack[:len(stack) - 1]
			}
			flag = errorAction()
		} else if s, ok := checkTable(x, a); ok {
			// fmt.Printf("output: %v -> %v\n", x, s)
			out := x + "->" + s
			fmt.Printf("%-16v", out)
			stack = stack[:len(stack)-1]
			if s != "#" {
				// push
				myPush(s)
			}
		}
		x = string(stack[len(stack)-1])
		fmt.Printf("%-16v\n", x)
		
		// 调试
		count++
		if count > 1000 {
			os.Exit(0)
		}
	}
	for _, item := range faultIndex {
		fmt.Printf("错误条目: %s\n", string(content[item:item+1]))
	}
	if flag {
		fmt.Printf("合法输入\n")
	} else {
		fmt.Printf("非法输入\n")
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
	if ret != "" && ret != "synch" {
		return ret, true
	}

	return ret, false
}

func errorAction() bool {
	// fmt.Printf("error")
	return false
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
