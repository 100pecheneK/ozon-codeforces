package main

import (
	"bufio"
	"fmt"
	"os"
)

type Report struct {
	valid string
	days  map[int]bool
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var count_reports int
	fmt.Fscan(in, &count_reports)

	var reports []*Report
	for i := 0; i < count_reports; i++ {
		var count_days int
		fmt.Fscan(in, &count_days)

		report := &Report{
			valid: "YES",
			days:  make(map[int]bool),
		}
		var day int
		fmt.Fscan(in, &day)
		report.days[day] = true
		for j := 1; j < count_days; j++ {
			prev_day := day
			fmt.Fscan(in, &day)

			_, ok := report.days[day]
			if !ok {
				report.days[day] = true
			}
			if day != prev_day {
				if !report.days[day] {
					report.valid = "NO"
					in.ReadString('\n')
					break
				}
				report.days[prev_day] = false
			}
		}
		reports = append(reports, report)
	}
	for _, report := range reports {
		fmt.Println(report.valid)
	}
}
