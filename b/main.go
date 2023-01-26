package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cart struct {
	items map[int]int
}

func (c *Cart) GetTotalPrice() int {
	price := 0
	for item, count := range c.items {
		count_with_prom := count/3*2 + count%3
		price += count_with_prom * item
	}
	return price
}

func (c *Cart) AddItem(item int) {
	_, ok := c.items[item]
	if !ok {
		c.items[item] = 1
		return
	}
	c.items[item] = c.items[item] + 1
}

func main() {
	var carts []*Cart
	reader := bufio.NewReader(os.Stdin)
	count_carts_input, _ := reader.ReadString('\n')
	count_carts, _ := strconv.Atoi(strings.TrimSpace(count_carts_input))

	for i := 0; i < count_carts; i++ {
		cart := &Cart{
			items: make(map[int]int),
		}
		reader.ReadString('\n')

		items_input, _ := reader.ReadString('\n')
		splitted_items := strings.Split(strings.TrimSpace(items_input), " ")
		for _, i := range splitted_items {
			item, _ := strconv.Atoi(i)
			cart.AddItem(item)
		}
		carts = append(carts, cart)
	}

	for _, cart := range carts {
		fmt.Println(cart.GetTotalPrice())
	}
}
