package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"time"
)

const (
	time_layout = "15:04:05"
	h           = "([0-1][0-9]|2[0-3])"
	m_s         = "([0-4][0-9]|5[0-9])"
)

type TimeInterval struct {
	start time.Time
	end   time.Time
}

func main() {
	var count_data int
	var count_time int
	var time_input string
	var time_interval_group []*TimeInterval

	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscanln(stdin, &count_data)

	grammar := fmt.Sprintf("(%s:%s:%s)-(%s:%s:%s)", h, m_s, m_s, h, m_s, m_s)
	time_regexp := regexp.MustCompile(grammar)

	for i := 0; i < count_data; i++ {
		fmt.Fscanln(stdin, &count_time)

		time_interval_group = make([]*TimeInterval, 0)
		valid := true
		for j := 0; j < count_time; j++ {
			fmt.Fscanln(stdin, &time_input)

			match_time := time_regexp.FindStringSubmatch(time_input)
			if len(match_time) == 0 {
				valid = false
				skipLines(stdin, count_time-j-1)
				break
			}
			time_1, _ := time.Parse(time_layout, match_time[1])
			time_2, _ := time.Parse(time_layout, match_time[5])
			diff := time_2.Sub(time_1)
			if diff < 0 {
				valid = false
				skipLines(stdin, count_time-j-1)
				break
			}

			time_interval_group = append(time_interval_group, &TimeInterval{start: time_1, end: time_2})
		}

		if !valid {
			fmt.Println("NO")
			continue
		}

		sort.Slice(time_interval_group, func(i, j int) bool {
			return time_interval_group[i].start.Before(time_interval_group[j].start)
		})

		for j := 0; j < count_time-1; j++ {
			if (time_interval_group[j].start.Before(time_interval_group[j+1].end) || time_interval_group[j].start.Equal(time_interval_group[j+1].end)) && (time_interval_group[j+1].start.Before(time_interval_group[j].end) || time_interval_group[j+1].start.Equal(time_interval_group[j].end)) {
				valid = false
				break
			}
		}
		if !valid {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
	}
}

func skipLines(r *bufio.Reader, count int) {
	for i := 0; i < count; i++ {
		r.ReadString('\n')
	}
}
