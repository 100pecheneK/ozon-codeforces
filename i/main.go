package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Proc struct {
	start  int
	time   int
	energo int
}

func main() {
	var count_procs int
	var count_tasks int
	var procs []*Proc
	var total_energo int

	stdio := bufio.NewReader(os.Stdin)
	fmt.Fscanln(stdio, &count_procs, &count_tasks)

	var energo int
	for i := 0; i < count_procs; i++ {
		fmt.Fscan(stdio, &energo)
		procs = append(procs, &Proc{energo: energo})
	}

	sort.Slice(procs, func(i, j int) bool {
		return procs[i].energo < procs[j].energo
	})

	var current_time int
	for i := 0; i < count_tasks; i++ {
		fmt.Fscan(stdio, &current_time, &energo)
		for _, proc := range procs {
			if current_time-proc.start >= proc.time {
				total_energo += energo * proc.energo
				proc.start = current_time
				proc.time = energo
				break
			}
		}
	}
	fmt.Println(total_energo)
}
