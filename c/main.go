package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Developer struct {
	i   int
	lvl int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	developers_groups := make(map[int]*list.List)

	count_data_input, _ := reader.ReadString('\n')
	count_data, _ := strconv.Atoi(strings.TrimSpace(count_data_input))

	for i := 0; i < count_data; i++ {
		reader.ReadString('\n')
		developers := list.New()

		lvls_input, _ := reader.ReadString('\n')
		lvls := strings.Split(strings.TrimSpace(lvls_input), " ")
		for j, lvl_input := range lvls {
			lvl, _ := strconv.Atoi(lvl_input)
			developers.PushBack(&Developer{i: j + 1, lvl: lvl})
		}
		developers_groups[i] = developers
	}

	for i := 0; i < count_data; i++ {
		developers_group := developers_groups[i]
		for developer_node := developers_group.Front(); developer_node != nil; developer_node = developer_node.Next() {
			dev := developer_node.Value.(*Developer)

			partner_node := developer_node.Next()
			partner := partner_node.Value.(*Developer)

			minLvlDif := Abs(dev.lvl - partner.lvl)

			for dev_partner_node := partner_node.Next(); dev_partner_node != nil; dev_partner_node = dev_partner_node.Next() {
				dev_partner := dev_partner_node.Value.(*Developer)

				dif := Abs(dev.lvl - dev_partner.lvl)
				if minLvlDif > dif {
					minLvlDif = dif
					partner_node = dev_partner_node
					partner = dev_partner
				}
			}
			fmt.Println(dev.i, partner.i)
			developers_group.Remove(partner_node)
		}
		fmt.Println()
	}

}
