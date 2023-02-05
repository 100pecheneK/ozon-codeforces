package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
)

func main() {
	var dict_len int
	var count_words int
	var dict []string

	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscanln(stdin, &dict_len)

	var d string
	for i := 0; i < dict_len; i++ {

		fmt.Fscanln(stdin, &d)
		dict = append(dict, d)
	}

	fmt.Fscanln(stdin, &count_words)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var words []string
	answers := make(map[string]string)
	var word string
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < count_words; i++ {
			fmt.Fscanln(stdin, &word)
			words = append(words, word)
			wg.Add(1)
			go func(word string, answers *map[string]string) {
				defer wg.Done()
				a := logic(word, &dict)
				mutex.Lock()
				(*answers)[word] = a
				mutex.Unlock()
			}(word, &answers)
		}
	}()

	wg.Wait()
	for i := 0; i < count_words; i++ {
		fmt.Println(answers[words[i]])
	}
}

func logic(word string, dict *[]string) string {
	rifms := make(map[string]int)
	word_len := len(word)
	for _, d := range *dict {
		if d == word {
			break
		}
		d_len := len(d)

		count := 0

		dd := reverse(d)
		w := reverse(word)

		for i := 0; i < word_len; i++ {
			if d_len-1 < i {
				break
			}
			if w[i] == dd[i] {
				count++
			}

		}
		if count != 0 {
			rifms[d] = count
		}
	}
	if len(rifms) == 0 {
		return (*dict)[0]
	}

	keys := make([]string, 0, len(rifms))
	for k := range rifms {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return rifms[keys[i]] > rifms[keys[j]]
	})

	return keys[0]
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
