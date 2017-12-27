package main

import (
	"flag"
	"time"

	"github.com/songjiayang/goxmas/tree"
)

var (
	company string
	live    bool
)

func main() {
	flag.StringVar(&company, "company", "goxmas", "what company to say Merry Christmas.")
	flag.BoolVar(&live, "live", false, "Keep printing the tree.")

	flag.Parse()

	masTree := tree.NewTree(40, company)

	for true {
		masTree.Print()
		if !live {
			return
		}
		time.Sleep(2 * time.Second)
	}
}
