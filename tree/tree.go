package tree

import "fmt"

type Tree struct {
	company string

	leaves [][]string
}

func NewTree(width int, company string) *Tree {
	height := width / 2

	leaves := make([][]string, height)

	for i := 0; i < height; i++ {
		leaves[i] = newLevelLeaves(width, " ")
		if i == 0 {
			leaves[i][width/2] = "★"
			continue
		}

		leaves[i][height-i] = "/"
		leaves[i][height+i] = "\\"
		for j := (height - i + 1); j < height+i; j++ {
			leaves[i][j] = leafContent()
		}
	}

	leaves = append(leaves, bottomLeaves(width, "^"))
	leaves = append(leaves, bottomLeaves(width, " "))

	return &Tree{
		leaves:  leaves,
		company: company,
	}
}

func (tree *Tree) Print() {
	fmt.Println()

	for _, leaves := range tree.leaves {
		for _, leaf := range leaves {
			switch leaf {
			case "★":
				yellowColor().Print(leaf)
			default:
				randomColor().Print(leaf)
			}
		}

		fmt.Print("\n")
	}

	fmt.Println()

	yellowColor().Print(tree.company)
	redColor().Print(" loves")
	yellowColor().Print(" you,")
	redColor().Print(" Merry Christmas!")

	fmt.Println()
	fmt.Println()
}
