// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// type Cell struct {
// 	x, y int
// }
// type Region struct {
// 	color  string
// 	cells  []*Cell
// 	closed bool
// }

// func (r *Region) isCellExist(x, y int) bool {
// 	for _, cell := range r.cells {
// 		if cell.x == x && cell.y == y {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (r *Region) add(cell *Cell) {
// 	r.cells = append(r.cells, cell)
// }

// type Node struct {
// 	Value *Cell
// 	Next  *Node
// }
// type Queue struct {
// 	node *Node
// 	size int
// }

// func (queue *Queue) Push(v *Cell) {
// 	fmt.Println("QUEUE-PUSH:", v.x, v.y)
// 	if queue.node == nil {
// 		queue.node = &Node{v, nil}
// 		return
// 	}
// 	queue.node = &Node{v, queue.node}
// 	queue.size++
// }

// func (queue Queue) Pop() (*Cell, bool) {
// 	if queue.size == 0 {
// 		return nil, false
// 	}
// 	if queue.size == 1 {
// 		queue.size--
// 		return queue.node.Value, true
// 	}
// 	var newLastNode *Node
// 	for queue.node.Next != nil {
// 		newLastNode = queue.node
// 		queue.node = queue.node.Next
// 	}

// 	oldLastNodeValue := newLastNode.Next.Value
// 	newLastNode.Next = nil

// 	queue.size--
// 	return oldLastNodeValue, true
// }

// func (queue Queue) traverse() {
// 	if queue.node == nil {
// 		fmt.Println("empty")
// 		return
// 	}
// 	for queue.node != nil {
// 		fmt.Printf("%v -> ", queue.node.Value)
// 		queue.node = queue.node.Next
// 	}
// 	fmt.Println()
// }
// func main() {
// 	var data_count int
// 	var len_x, columns_count int
// 	var cell_map [][]string
// 	var queue *Queue
// 	var region *Region
// 	stdin := bufio.NewReader(os.Stdin)

// 	fmt.Fscanln(stdin, &data_count)

// 	for data_index := 0; data_index < data_count; data_index++ {
// 		fmt.Fscanln(stdin, &len_x, &columns_count)

// 		// write map
// 		cell_map = make([][]string, len_x)
// 		for row := 0; row < len_x; row++ {
// 			r, _ := stdin.ReadString('\n')
// 			cells := strings.Split(strings.ReplaceAll(strings.TrimSpace(r), ".", ""), "")
// 			for _, cell := range cells {
// 				cell_map[row] = append(cell_map[row], cell)
// 			}
// 		}
// 		regions := make(map[string]*Region)
// 		fmt.Printf("\nMap:\n")
// 		for i, r := range cell_map {
// 			fmt.Printf("%d: %v %d\n", i, r, len(r))
// 		}
// 		fmt.Println()
// 		// traverse map
// 		for x, row := range cell_map {
// 			len_y := len(row)
// 			queue = &Queue{}
// 			for y, cell := range row {
// 				fmt.Println("CURRENT CELL:", x, y, cell)
// 				if r, exist := regions[cell]; exist {
// 					if r.closed {
// 						// WORK IS DONE :)
// 						fmt.Println("REGION CLOSED:", r.color)
// 						os.Exit(1)
// 					}
// 				} else {
// 					region = &Region{
// 						color: cell,
// 					}
// 					regions[cell] = region
// 				}

// 				queue.Push(&Cell{x, y})

// 				for queue.node != nil {
// 					// 	fmt.Println("QN:", queue.node)
// 					// 	fmt.Printf("%d:%d -> ", queue.node.Value.x, queue.node.Value.y)
// 					fmt.Println("====")
// 					fmt.Println("cord:", x, y)
// 					queue.traverse()
// 					logic(region, queue.node.Value.x, queue.node.Value.y, len_x, len_y, &cell_map, cell, queue)

// 					queue.node = queue.node.Next
// 					time.Sleep(1 * time.Second)
// 				}
// 			}

// 			region.closed = true
// 			fmt.Println("Region:", region.color)
// 			for _, c := range region.cells {
// 				fmt.Printf("[%d:%d] ", c.x, c.y)
// 			}
// 			fmt.Println()
// 			fmt.Println()
// 		}
// 	}
// }

// func checkIfCordExistInMatrix(x, y, len_y, len_x int) bool {
// 	if x < 0 || x > len_x-1 {
// 		return false
// 	}
// 	if y < 0 || y > len_y-1 {
// 		return false
// 	}
// 	return true

// }

// func logic(region *Region, x, y, len_x, len_y int, cell_map *[][]string, cell string, queue *Queue) {
// 	// add current cell if not exist
// 	if !region.isCellExist(x, y) {
// 		fmt.Println("REGION-ADD:", x, y)
// 		c := &Cell{x, y}
// 		region.add(c)
// 	}

// 	// left exist
// 	if checkIfCordExistInMatrix(x, y-1, len_x, len_y) {
// 		if (*cell_map)[x][y-1] == cell {
// 			// not in region
// 			if !region.isCellExist(x, y-1) {
// 				c := &Cell{x: x, y: y - 1}
// 				fmt.Println("REGION-ADD:", x, y-1)
// 				region.add(c)
// 				queue.Push(c)
// 			}
// 		}
// 	}
// 	// right exist
// 	if checkIfCordExistInMatrix(x, y+1, len_x, len_y) {
// 		if (*cell_map)[x][y+1] == cell {
// 			// not in region
// 			if !region.isCellExist(x, y+1) {
// 				c := &Cell{x: x, y: y + 1}
// 				region.add(c)
// 				queue.Push(c)
// 			}
// 		}
// 	}
// 	// top left exist
// 	if checkIfCordExistInMatrix(x-1, y-1, len_x, len_y) {
// 		if (*cell_map)[x-1][y-1] == cell {
// 			// not in region
// 			if !region.isCellExist(x-1, y-1) {
// 				c := &Cell{x: x - 1, y: y - 1}
// 				region.add(c)
// 				queue.Push(c)
// 			}
// 		}
// 	}
// 	// top right exist
// 	if checkIfCordExistInMatrix(x-1, y, len_x, len_y) {
// 		if (*cell_map)[x-1][y] == cell {
// 			// not in region
// 			if !region.isCellExist(x-1, y) {
// 				c := &Cell{x: x - 1, y: y}
// 				region.add(c)
// 				queue.Push(c)
// 			}
// 		}
// 	}
// 	// bot left exist
// 	if checkIfCordExistInMatrix(x+1, y-1, len_x, len_y) {
// 		if (*cell_map)[x+1][y-1] == cell {
// 			// not in region
// 			if !region.isCellExist(x+1, y-1) {
// 				c := &Cell{x: x - 1, y: y}
// 				region.add(c)
// 				queue.Push(c)
// 			}
// 		}
// 	}
// 	// bot right exist
// 	if checkIfCordExistInMatrix(x+1, y, len_x, len_y) {
// 		if (*cell_map)[x+1][y] == cell {
// 			// not in region
// 			if !region.isCellExist(x+1, y) {
// 				c := &Cell{x: x - 1, y: y}
// 				region.add(c)
// 				queue.Push(c)
// 			}
// 		}
// 	}

// }
