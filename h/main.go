package main

import (
	"container/list"
	"fmt"
	"os"
)

type Cell struct {
	x, y int
}
type Region struct {
	color  string
	cells  []*Cell
	closed bool
}

func (r *Region) isCellExist(x, y int) bool {
	for _, cell := range r.cells {
		if cell.x == x && cell.y == y {
			return true
		}
	}
	return false
}

func (r *Region) add(cell *Cell) {
	r.cells = append(r.cells, cell)
}

// New is a new instance of a Queue
func NewQueue() *list.List {
	return list.New()
}

func main() {
	data_count := 1
	len_x := 3
	var queue *list.List
	var region *Region

	for data_index := 0; data_index < data_count; data_index++ {
		cell_map := [][]string{{"R", "R", "R", "G"}, {"Y", "G", "G"}, {"B", "Y", "V", "V"}}
		regions := make(map[string]*Region)
		fmt.Printf("\nMap:\n")
		for i, r := range cell_map {
			fmt.Printf("%d: %v %d\n", i, r, len(r))
		}
		fmt.Println()
		// traverse map
		for x, row := range cell_map {
			queue = NewQueue()
			for y, cell := range row {
				fmt.Println(cell)
				if r, exist := regions[cell]; exist {
					if exst := region.isCellExist(x, y); exst {
						continue
					}
					if r.closed {
						// WORK IS DONE :)
						fmt.Println("REGION CLOSED:", r.color)
						os.Exit(1)
					}

				} else {
					region = &Region{
						color: cell,
					}
					regions[cell] = region
				}

				queue.PushBack(&Cell{x, y})

				for queue.Len() > 0 {
					e := queue.Front()

					logic(region, e.Value.(*Cell).x, e.Value.(*Cell).y, len_x, &cell_map, cell, queue)
					// здзесь нода неверно удаляется

					queue.Remove(e)
					// queue.node = queue.node.Next
				}
				region.closed = true
			}
			fmt.Println()
		}
	}
}

func checkIfCordExistInMatrix(x, y, len_x, len_y int) bool {
	if x < 0 || x > len_x-1 {
		return false
	}
	if y < 0 || y > len_y-1 {
		return false
	}
	return true

}
func checkIfRowExist(x, len_x int) bool {
	return x < len_x && x >= 0
}
func logic(region *Region, x, y, len_x int, cell_map *[][]string, cell string, queue *list.List) {
	len_y := len((*cell_map)[x])
	var len_y_top int
	var len_y_bot int
	if exst := checkIfRowExist(x-1, len_x); exst {
		len_y_top = len((*cell_map)[x-1])
	}
	if exst := checkIfRowExist(x+1, len_x); exst {
		len_y_bot = len((*cell_map)[x+1])
	}
	// add current cell if not exist
	if !region.isCellExist(x, y) {
		fmt.Println("REGION-ADD:", x, y)
		c := &Cell{x, y}
		region.add(c)
	}

	// left exist
	exist := checkIfCordExistInMatrix(x, y-1, len_x, len_y)
	if exist {
		if (*cell_map)[x][y-1] == cell {
			// not in region
			if !region.isCellExist(x, y-1) {
				c := &Cell{x: x, y: y - 1}
				fmt.Println("REGION-ADD:", x, y-1)
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	// right exist
	exist = checkIfCordExistInMatrix(x, y+1, len_x, len_y)
	if exist {
		if (*cell_map)[x][y+1] == cell {
			// not in region
			if !region.isCellExist(x, y+1) {
				c := &Cell{x: x, y: y + 1}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	// top left exist
	exist = checkIfCordExistInMatrix(x-1, y-1, len_x, len_y_top)
	if exist {
		if (*cell_map)[x-1][y-1] == cell {
			// not in region
			if !region.isCellExist(x-1, y-1) {
				c := &Cell{x: x - 1, y: y - 1}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	// top right exist
	exist = checkIfCordExistInMatrix(x-1, y, len_x, len_y_top)
	if exist {
		if (*cell_map)[x-1][y] == cell {
			// not in region
			if !region.isCellExist(x-1, y) {
				c := &Cell{x: x - 1, y: y}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	// bot left exist
	exist = checkIfCordExistInMatrix(x+1, y-1, len_x, len_y_bot)
	if exist {
		if (*cell_map)[x+1][y-1] == cell {
			// not in region
			if !region.isCellExist(x+1, y-1) {
				c := &Cell{x: x + 1, y: y - 1}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	// bot right exist
	exist = checkIfCordExistInMatrix(x+1, y, len_x, len_y_bot)
	if exist {
		if (*cell_map)[x+1][y] == cell {
			// not in region
			if !region.isCellExist(x+1, y) {
				c := &Cell{x: x + 1, y: y}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}

}
