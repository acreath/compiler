package main

import (
	"fmt"
	"os"
	"strconv"
)

type SymTable struct {
	previous *SymTable
	head *SymNode
	last *SymNode
	width int
}

type SymNode struct {
	name string
	offset int
	symType string;
	next *SymNode

}

func initss() {
	T1 = make(map[string]string)
	T2 = make(map[string]string)
	T3 = make(map[string]string)
	T4 = make(map[string]string)
	id1 = make(map[string]string)
	id2 = make(map[string]string)
	id3 = make(map[string]string)
	id1["name"] = "id1"
	id2["name"] = "id2"
	id3["name"] = "id3"
}

func mktable(previous *SymTable) (*SymTable) {
	var nTable SymTable
	nTable.previous = previous
	nTable.head = nil
	nTable.last = nil
	return &nTable
}

func enter(table *SymTable, name string, tType string, noff int) {
	var nNode SymNode
	nNode.name = name
	nNode.symType = tType
	nNode.offset = noff
	nNode.next = nil
	if table.head == nil {
		table.head = &nNode
		table.last = &nNode
	} else {
		table.last.next = &nNode
		table.last = &nNode
	}
}

func addwidth(tblp *SymTable, off int) {
	tblp.width = off
}

func toptbl(tblptr [MAX_SIZE]*SymTable) (*SymTable) {
	return tblptr[tbltop - 1]
}

func topoff(offset [MAX_SIZE]int) (int) {
	return offset[offtop - 1]
}

func pushtbl(t *SymTable, tblptr *[MAX_SIZE]*SymTable) {
	tblptr[tbltop] = t;
	tbltop += 1
}
func pushoff(noff int, offset *[MAX_SIZE]int) {
	offset[offtop] = noff
	offtop += 1
}

func poptbl(tblptr [MAX_SIZE]*SymTable) (*SymTable) {
	tbl := tblptr[tbltop - 1]
	tbltop -= 1
	return tbl
}

func popoff(offset [MAX_SIZE]int) (int) {
	n := offset[offtop - 1]
	offtop -= 1
	return n
}

func setoff(offset *[MAX_SIZE]int, widstr string) {
	width, err := strconv.Atoi(widstr)
	if err != nil {
		fmt.Printf("setoff func error\n")
		os.Exit(0)
	}
	offset[offtop - 1] += width
}

func pointer(s string) (string) {
	return "*" + s
}

func printTbl(tbl *SymTable) {
	fmt.Printf("%v\n", tbl.width)
	nNode := tbl.head
	for nNode != nil {
		fmt.Printf("%-12v%-12v%-12v\n", nNode.name, nNode.symType, nNode.offset)
		nNode = nNode.next
	}
}