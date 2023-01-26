package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Summator struct {
	A int
	B int
}

func (s *Summator) Calculate() int {
	return s.A + s.B
}

func main() {
	var summators []Summator
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	length, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < length; i++ {
		scanner.Scan()
		input := scanner.Text()
		numbers := strings.Split(input, " ")

		A, _ := strconv.Atoi(numbers[0])
		B, _ := strconv.Atoi(numbers[1])

		summators = append(summators, Summator{A, B})
	}
	for _, summator := range summators {
		fmt.Println(summator.Calculate())
	}

}
