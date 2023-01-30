package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
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

func NewQueue() *list.List {
	return list.New()
}

func main() {
	var data_count int
	var len_x, columns_count int
	var queue *list.List
	var region *Region
	var cell_map [][]string
	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscanln(stdin, &data_count)

NO:
	for data_index := 0; data_index < data_count; data_index++ {
		fmt.Fscanln(stdin, &len_x, &columns_count)
		cell_map = make([][]string, len_x)
		for row := 0; row < len_x; row++ {
			r, _ := stdin.ReadString('\n')
			cells := strings.Split(strings.ReplaceAll(strings.TrimSpace(r), ".", ""), "")
			cell_map[row] = append(cell_map[row], cells...)
		}
		regions := make(map[string]*Region)
		for x, row := range cell_map {
			queue = NewQueue()
			for y, cell := range row {
				if r, exist := regions[cell]; exist {
					if exst := r.isCellExist(x, y); exst {
						continue
					}
					if r.closed {
						fmt.Println("NO")
						continue NO
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
					queue.Remove(e)
				}
				region.closed = true
			}
		}
		fmt.Println("YES")
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
	if !region.isCellExist(x, y) {
		c := &Cell{x, y}
		region.add(c)
	}

	exist := checkIfCordExistInMatrix(x, y-1, len_x, len_y)
	if exist {
		if (*cell_map)[x][y-1] == cell {
			if !region.isCellExist(x, y-1) {
				c := &Cell{x: x, y: y - 1}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	exist = checkIfCordExistInMatrix(x, y+1, len_x, len_y)
	if exist {
		if (*cell_map)[x][y+1] == cell {
			if !region.isCellExist(x, y+1) {
				c := &Cell{x: x, y: y + 1}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}

	top_left_x, top_left_y := getCordByX(x, x-1, y-1, x-1, y)
	exist = checkIfCordExistInMatrix(top_left_x, top_left_y, len_x, len_y_top)
	if exist {
		if (*cell_map)[top_left_x][top_left_y] == cell {
			if !region.isCellExist(x-1, y-1) {
				c := &Cell{x: top_left_x, y: top_left_y}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	top_left_x, top_left_y = getCordByX(x, x-1, y, x-1, y+1)
	exist = checkIfCordExistInMatrix(top_left_x, top_left_y, len_x, len_y_top)
	if exist {
		if (*cell_map)[top_left_x][top_left_y] == cell {
			if !region.isCellExist(x-1, y) {
				c := &Cell{x: top_left_x, y: top_left_y}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	top_left_x, top_left_y = getCordByX(x, x+1, y-1, x+1, y)
	exist = checkIfCordExistInMatrix(top_left_x, top_left_y, len_x, len_y_bot)
	if exist {
		if (*cell_map)[top_left_x][top_left_y] == cell {
			if !region.isCellExist(x+1, y-1) {
				c := &Cell{x: top_left_x, y: top_left_y}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}
	top_left_x, top_left_y = getCordByX(x, x+1, y, x+1, y+1)
	exist = checkIfCordExistInMatrix(top_left_x, top_left_y, len_x, len_y_bot)
	if exist {
		if (*cell_map)[top_left_x][top_left_y] == cell {
			if !region.isCellExist(x+1, y) {
				c := &Cell{x: top_left_x, y: top_left_y}
				region.add(c)
				queue.PushBack(c)
			}
		}
	}

}

func getCordByX(x, x_0, y_0, x_1, y_1 int) (int, int) {
	if x%2 == 0 {
		return x_0, y_0
	} else {
		return x_1, y_1
	}
}
