package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "sort"
)

//go:embed input
var input string

const (
	debug = true

	MAX_INT = 2147483647
)

type Point struct {
	x, y, z int
}

func NewPoint(s string) Point {
	splited := strings.Split(s, ",")
	x, _ := strconv.Atoi(splited[0])
	y, _ := strconv.Atoi(splited[1])
	z, _ := strconv.Atoi(splited[2])
	return Point{x, y , z}
}

type Block struct {
	id int
	p1, p2 Point
	supports []int
	lays []int
}

func NewBlock(id int, s string) *Block {
	splited := strings.Split(s, "~")
	return &Block{id, NewPoint(splited[0]), NewPoint(splited[1]), []int{}, []int{}}
}

func (b *Block) String() string {
	return fmt.Sprintf("%v", *b)
}

func (b *Block) AddSupports(id int) {
	b.supports = append(b.supports, id)
}

func (b *Block) AddLays(id int) {
	b.lays = append(b.lays, id)
}

func (b *Block) LeftTopXY() (x, y int) {
	if b.p1.x < b.p2.x {
		if b.p1.y > b.p2.y {
			return b.p1.x, b.p1.y
		} else {
			return b.p1.x, b.p2.y
		}
	} else {
		if b.p1.y > b.p2.y {
			return b.p2.x, b.p1.y
		} else {
			return b.p2.x, b.p2.y
		}
	}
}

func (b *Block) RightBottomXY() (x, y int) {
	if b.p1.x < b.p2.x {
		if b.p1.y > b.p2.y {
			return b.p2.x, b.p2.y
		} else {
			return b.p2.x, b.p1.y
		}
	} else {
		if b.p1.y > b.p2.y {
			return b.p1.x, b.p2.y
		} else {
			return b.p1.x, b.p1.y
		}
	}
}

func (b *Block) MinZ() int {
	if b.p1.z < b.p2.z {
		return b.p1.z
	} else {
		return b.p2.z
	}
}

func (b *Block) MaxZ() int {
	if b.p1.z > b.p2.z {
		return b.p1.z
	} else {
		return b.p2.z
	}
}

func (b *Block) MoveZ(delta int) {
	if debug {
		fmt.Printf("moving z: %v\n", delta)
	}
	b.p1.z += delta
	b.p2.z += delta
}

type ByLowestPoint []*Block

func (a ByLowestPoint) Len() int           { return len(a) }
func (a ByLowestPoint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLowestPoint) Less(i, j int) bool {
	switch {
	case a[i].p1.z < a[j].p2.z:
		return true
	case a[i].p1.z > a[j].p2.z:
		return false
	case a[i].p1.x < a[j].p2.x:
		return true
	case a[i].p1.x > a[j].p2.x:
		return false
	case a[i].p1.y < a[j].p2.y:
		return true
	case a[i].p1.y > a[j].p2.y:
		return false
	}
	return false
}

var space [][][]int
var blocks []*Block
var blocksMap map[int]*Block

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	blocksMap = make(map[int]*Block)
	for i, line := range inputAsLines {
		b := NewBlock(i, line)
		blocks = append(blocks, b)
		blocksMap[i] = b
	}
	sort.Sort(ByLowestPoint(blocks))
	fmt.Printf("blocks = %v\n", blocks)

	fallDownBlocks()

	fmt.Printf("blocks = %v\n", blocks)

	fillSupportsLays()

	fmt.Printf("blocks = %v\n", blocks)

	sum := 0
	for _, b := range blocks {
		res := calculateHowManyWillFall(b)
		fmt.Printf("%v: falling %v\n", b.id, res)
		sum += res

	}
	fmt.Printf("sum = %v\n", sum)
}

func calculateHowManyWillFall(b *Block) int {
	// fmt.Printf("calculateHowManyWillFall -> %v\n", b.id)
	var removed []int
	var queue []int

	queue = append(queue, b.id)
	removed = append(removed, b.id)
	for len(queue) > 0 {
		currentId := queue[0]
		current := blocksMap[currentId]
		queue = queue[1:]
		// fmt.Printf("removed=%v\n", removed)

		supports := filterSlice(current.supports, removed)
		// fmt.Printf("supports -> %v\n", supports)
		for _, supportedId := range supports {
			supportedBlock := blocksMap[supportedId]
			lays := filterSlice(supportedBlock.lays, removed)
			// fmt.Printf("suppId=%v, lays -> %v\n", supportedId, lays)
			if len(lays) > 0 {
				continue
			}

			queue = append(queue, supportedBlock.id)
			removed = append(removed, supportedBlock.id)
		}
	}

	return len(removed) - 1
}

func fallDownBlocks() {
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		minZ := 1
		for j := i - 1; j >= 0; j-- {
			candidate := blocks[j]
			if !isOverlapping(block, candidate) {
				continue
			}

			newZ := candidate.MaxZ() + 1
			if minZ < newZ {
				minZ = newZ
			}
		}
		if minZ < block.MinZ() {
			block.MoveZ(minZ - block.MinZ())
		}
	}
}

func fillSupportsLays() {
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		for j := i - 1; j >= 0; j-- {
			candidate := blocks[j]
			if !isOverlapping(block, candidate) {
				continue
			}

			if block.MinZ() == candidate.MaxZ() + 1 {
				block.AddLays(candidate.id)
				candidate.AddSupports(block.id)
			}
		}
	}
}

func canBeDisintegrated(b *Block) bool {
	block := blocksMap[b.id]

	isOnlySupporter := false
	for _, supportedId := range block.supports {
		supportedBlock := blocksMap[supportedId]
		if len(supportedBlock.lays) == 1 {
			isOnlySupporter = true
			break
		}
	}

	return !isOnlySupporter
}

func isOverlapping(b1, b2 *Block) bool {
	l1x, l1y := b1.LeftTopXY()
	r1x, r1y := b1.RightBottomXY()
	l2x, l2y := b2.LeftTopXY()
	r2x, r2y := b2.RightBottomXY()

    // If one rectangle is on left side of other
    if l1x > r2x || l2x > r1x {
        return false
    }
 
    // If one rectangle is above other
    if r1y > l2y || r2y > l1y {
        return false;
    }
 
    return true;
}

func withingRange(p, q, v int) bool {
	if p <= q {
		return p <= v && v <= q
	} else {
		return q <= v && v <= p
	}
}

func filterSlice(orig []int, removed []int) []int {
	var res []int
	for _, x := range orig {
		found := false
		for _, y := range removed {
			if x == y {
				found = true
				break
			}
		}
		if !found {
			res = append(res, x)
		}
	}
	return res
}