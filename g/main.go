package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type User struct {
	id      int
	friends []*User
}

func NewUser(id int) *User {
	return &User{id: id}
}
func MakeFriends(user_1, user_2 *User) {
	user_1.friends = append(user_1.friends, user_2)
	user_2.friends = append(user_2.friends, user_1)
}

func main() {
	var count_users, count_pairs int
	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(stdin, &count_users, &count_pairs)

	users := make(map[int]*User, count_users)

	var user_1_id, user_2_id int
	for i := 0; i < count_users; i++ {
		users[i+1] = NewUser(i + 1)
	}
	for i := 0; i < count_pairs; i++ {
		fmt.Fscan(stdin, &user_1_id, &user_2_id)
		user_1 := users[user_1_id]
		user_2 := users[user_2_id]

		MakeFriends(user_1, user_2)
	}
	var user_recomendations map[int]int
	for i := 1; i <= count_users; i++ {
		user := users[i]

		user_recomendations = make(map[int]int)
		for _, friend := range user.friends {
		cont:
			for _, friend_friend := range friend.friends {
				if friend_friend.id == user.id {
					continue
				}

				for _, f := range user.friends {
					if f.id == friend_friend.id {
						continue cont
					}
				}
				if _, ok := user_recomendations[friend_friend.id]; ok {
					user_recomendations[friend_friend.id]++
				} else {
					user_recomendations[friend_friend.id] = 1
				}
			}
		}
		if len(user_recomendations) == 0 {
			fmt.Println(0)
			continue
		}
		var max int
		var ids []int

		for id, count := range user_recomendations {
			if max == 0 {
				max = count
				ids = append(ids, id)
			}
			if max < count {
				for _, i := range ids {
					delete(user_recomendations, i)
				}
				max = count
				ids = []int{id}
			} else if max == count {
				ids = append(ids, id)
			} else if max > count {
				delete(user_recomendations, id)
			}
		}
		keys := make([]string, len(user_recomendations))
		var i int
		for key := range user_recomendations {
			keys[i] = strconv.Itoa(key)
			i++
		}
		sort.Slice(keys, func(i, j int) bool {
			x, _ := strconv.Atoi(keys[i])
			y, _ := strconv.Atoi(keys[j])
			return x < y
		})
		fmt.Println(strings.Join(keys, " "))
	}
}
