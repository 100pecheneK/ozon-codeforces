package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type InputData struct {
	data         [][]int
	sort_columns []int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count_data, _ := strconv.Atoi(scanner.Text())

	input_data := make([]*InputData, count_data)
	for i := 0; i < count_data; i++ {
		scanner.Scan()
		scanner.Scan()
		count_rows_and_columns := strings.Split(scanner.Text(), " ")
		count_rows, _ := strconv.Atoi(count_rows_and_columns[0])
		count_columns, _ := strconv.Atoi(count_rows_and_columns[1])
		rows := &InputData{
			data: make([][]int, count_rows),
		}
		for row_index := 0; row_index < count_rows; row_index++ {
			scanner.Scan()
			row_input := strings.Split(scanner.Text(), " ")
			row := make([]int, count_columns)
			for r_i, r := range row_input {
				row[r_i], _ = strconv.Atoi(r)
			}

			rows.data[row_index] = row
		}
		scanner.Scan()
		columns_count, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		columns_input := strings.Split(scanner.Text(), " ")
		rows.sort_columns = make([]int, columns_count)
		for c_i, c := range columns_input {
			rows.sort_columns[c_i], _ = strconv.Atoi(c)
		}
		input_data[i] = rows
	}

	for _, data := range input_data {
		for _, c := range data.sort_columns {
			sort.SliceStable(data.data, func(i, j int) bool {
				return data.data[i][c-1] < data.data[j][c-1]
			})
		}
		printTable(&data.data)
		fmt.Println()
	}
}

func printTable(table *[][]int) {
	for _, r := range *table {
		fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r)), " "), "[]"))
	}
}
