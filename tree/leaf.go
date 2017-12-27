package tree

import (
	"math/rand"
	"time"
)

const leafSeed = "  $@O⸮?        "

var (
	leafSeedSize = len(leafSeed)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func leafContent() string {
	return string(leafSeed[rand.Intn(leafSeedSize)])
}

func newLevelLeaves(size int, content string) []string {
	target := make([]string, size)

	for i := 0; i < size; i++ {
		target[i] = content
	}

	return target
}

func bottomLeaves(size int, content string) []string {
	target := newLevelLeaves(size, content)
	target[0] = " "

	center := size / 2
	target[center] = " "
	target[center+1] = " "
	target[center-1] = "|"
	target[center+2] = "|"

	return target
}
