package main

import (
	"fmt"
)

const MAX_SIZE int = 20

var tblptr [MAX_SIZE]*SymTable
var tbltop int = 0;
var offset [MAX_SIZE]int
var offtop int = 0;


/* static input */
var T1, T2, T3, T4 map[string]string
var id1, id2, id3 map[string]string

func main() {
	fmt.Println("*********************************")
	initss()

	// 1
	t := mktable(nil)
	pushtbl(t, &tblptr)
	pushoff(0, &offset)

	// 2
	T1["type"] = "real"
	T1["width"] = "8"

	// 3
	enter(toptbl(tblptr), id1["name"], T1["type"], topoff(offset))
	setoff(&offset, T1["width"])

	// 4
	T3["type"] = "integer"
	T3["width"] = "4"

	// 5
	T2["type"] = pointer(T3["type"])
	T2["width"] = "4"

	// 6
	enter(toptbl(tblptr), id2["name"], T2["type"], topoff(offset))
	setoff(&offset, T2["width"])

	// 7
	T4["type"] = "integer"
	T4["width"] = "4"

	// 8
	enter(toptbl(tblptr), id3["name"], T4["type"], topoff(offset))
	setoff(&offset, T4["width"])

	// 9
	addwidth(toptbl(tblptr), topoff(offset))
	table1 := poptbl(tblptr)
	popoff(offset)

	// print table
	printTbl(table1)

}