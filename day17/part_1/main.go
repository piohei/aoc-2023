package main

import (
    "fmt"
	"container/heap"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

const (
	debug = false

	north = 0
	east = 1
	south = 2
	west = 3

	MAX_INT = 2147483647
)

// An Item is something we manage in a priority queue.
type Item struct {
	x, y int
	shortestPathLength int
	direction, count int

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (i *Item) String() string {
	return fmt.Sprintf("{x=%v,y=%v,shortestPathLength=%v,direction=%v,count=%v}", i.x, i.y, i.shortestPathLength, i.direction, i.count)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest shortestPathLength so we use less than here.
	return pq[i].shortestPathLength < pq[j].shortestPathLength
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, shortestPathLength, direction, count int) {
	item.shortestPathLength = shortestPathLength
	item.direction = direction
	item.count = count
	heap.Fix(pq, item.index)
}

var cityMap [][]int
var shortestPaths [][][][]int
var directions [][]int
var shortestPathDirection [][]int
var maxX int
var maxY int

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for y, line := range inputAsLines {
		if cityMap == nil {
			cityMap = make([][]int, len(line))
			shortestPaths = make([][][][]int, len(line))
			directions = make([][]int, len(line))
			shortestPathDirection = make([][]int, len(line))
		}

		for x, char := range strings.Split(line, "") {
			if cityMap[x] == nil {
				cityMap[x] = make([]int, len(inputAsLines))
				shortestPaths[x] = make([][][]int, len(inputAsLines))
				directions[x] = make([]int, len(inputAsLines))
				shortestPathDirection[x] = make([]int, len(inputAsLines))
			}
			cityMap[x][y], _ = strconv.Atoi(char)
			directions[x][y] = -1
			shortestPathDirection[x][y] = -1
		}
	}

	maxX = len(cityMap) - 1
	maxY = len(cityMap[0]) - 1

	for y := range cityMap[0] {
		for x := range cityMap {
			shortestPaths[x][y] = make([][]int, 4)
			shortestPaths[x][y][0] = []int{MAX_INT, MAX_INT, MAX_INT}
			shortestPaths[x][y][1] = []int{MAX_INT, MAX_INT, MAX_INT}
			shortestPaths[x][y][2] = []int{MAX_INT, MAX_INT, MAX_INT}
			shortestPaths[x][y][3] = []int{MAX_INT, MAX_INT, MAX_INT}
		}
	}


	fmt.Printf("------------------------------------------------------------\n")
	for y := range cityMap[0] {
		for x := range cityMap {
			fmt.Printf("%v", cityMap[x][y])
		}
		fmt.Printf("\n")
	}

	calculateShortestPath()

	fmt.Printf("------------------------------------------------------------\n")
	printPaths()
	fmt.Printf("------------------------------------------------------------\n")
	// printShortestPath()
	fmt.Printf("------------------------------------------------------------\n")

	res := MAX_INT
	for _, counts := range shortestPaths[maxX][maxY] {
		for _, val := range counts {
			if val < res {
				res = val
			}
		}
	}
	fmt.Printf("res = %v\n", res)

	// sum := 0
	// for y := range contraption[0] {
	// 	for x := range contraption {
	// 		p := contraption[x][y]
	// 		if p.beamDirection[north] || p.beamDirection[east] || p.beamDirection[south] || p.beamDirection[west] {
	// 			sum++
	// 		}
	// 	}
	// }
	// fmt.Printf("sum = %v\n", sum)
}

func printPaths() {
	for y := range cityMap[0] {
		for x := range cityMap {
			switch directions[x][y] {
			case north:
				fmt.Printf("^")
			case east:
				fmt.Printf(">")
			case south:
				fmt.Printf("v")
			case west:
				fmt.Printf("<")
			default:
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

// func printShortestPath() {
// 	cx := maxX
// 	cy := maxY
// 	for cx != 0 || cy != 0 {
// 		shortestPathDirection[cx][cy] = directions[cx][cy]
// 		if debug {
// 			fmt.Printf("cx = %v, cy = %v, shortestPathDirection = %v\n", cx, cy, shortestPathDirection[cx][cy])
// 		}
// 		switch directions[cx][cy] {
// 		case north:
// 			cy = cy + 1
// 		case east:
// 			cx = cx - 1
// 		case south:
// 			cy = cy - 1
// 		case west:
// 			cx = cx + 1
// 		}
// 		if debug {
// 			fmt.Printf("cx = %v, cy = %v\n", cx, cy)
// 		}
// 	}

// 	for y := range cityMap[0] {
// 		for x := range cityMap {
// 			switch shortestPathDirection[x][y] {
// 			case north:
// 				fmt.Printf("^")
// 			case east:
// 				fmt.Printf(">")
// 			case south:
// 				fmt.Printf("v")
// 			case west:
// 				fmt.Printf("<")
// 			default:
// 				fmt.Printf(".")
// 			}
// 		}
// 		fmt.Printf("\n")
// 	}
// }

func calculateShortestPath() {
	var pq PriorityQueue
	pq = append(pq, &Item{
		x: 0, y: 0, shortestPathLength: 0, direction: east, count: -1,
	})
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		for _, m := range nextMoves(item) {
			if m.shortestPathLength < shortestPaths[m.x][m.y][m.direction][m.count] {
				if debug {
					fmt.Printf("new shortest path for x=%v,y=%v, m=%v\n", item.x, item.y, m)
				}
				shortestPaths[m.x][m.y][m.direction][m.count] = m.shortestPathLength
				// directions[m.x][m.y] = m.direction
				heap.Push(&pq, m)
			}
		}
	}
}

func nextMoves(item *Item) []*Item {
	var res []*Item
	switch item.direction {
	case north:
		if nextMove := createItem(item, west); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, north); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, east); nextMove != nil {
			res = append(res, nextMove)
		}
	case east:
		if nextMove := createItem(item, north); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, east); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, south); nextMove != nil {
			res = append(res, nextMove)
		}
	case south:
		if nextMove := createItem(item, east); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, south); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, west); nextMove != nil {
			res = append(res, nextMove)
		}
	case west:
		if nextMove := createItem(item, south); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, west); nextMove != nil {
			res = append(res, nextMove)
		}
		if nextMove := createItem(item, north); nextMove != nil {
			res = append(res, nextMove)
		}
	}
	return res
}

func createItem(origItem *Item, direction int) *Item {
	var newX, newY int
	switch direction {
	case north:
		newX = origItem.x
		newY = origItem.y - 1
	case east:
		newX = origItem.x + 1
		newY = origItem.y
	case south:
		newX = origItem.x
		newY = origItem.y + 1
	case west:
		newX = origItem.x - 1
		newY = origItem.y
	}

	if !(0 <= newX && newX <= maxX && 0 <= newY && newY <= maxY) {
		return nil
	}

	newDirection := direction
	newCount := 0
	newShortestPathLength := origItem.shortestPathLength + cityMap[newX][newY]
	if direction == origItem.direction {
		if origItem.count < 2 {
			newCount = origItem.count + 1
		} else {
			return nil
		}
	}

	return &Item{x: newX, y: newY, direction: newDirection, count: newCount, shortestPathLength: newShortestPathLength}
}