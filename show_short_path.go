package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	s, _ := scanner.ReadString('\n')
	s = strings.TrimSpace(s)
	n, _ := strconv.Atoi(s)
	x := []string{}
	dict := make(map[int][]int)
	for i := 0; i < n; i++ {
		s, _ = scanner.ReadString('\n')
		s = strings.TrimSpace(s)
		x = strings.Split(s, " ")
		for j := 0; j < n; j++ {
			if x[j] == "1" {
				dict[i] = append(dict[i], j)
			}
		}
	}

	v, _ := scanner.ReadString('\n')
	v = strings.TrimSpace(v)
	z := strings.Split(v, " ")
	q, _ := strconv.Atoi(z[0])
	w, _ := strconv.Atoi(z[1])
	q, w = q-1, w-1
	stack := [][2]int{}
	stack = append(stack, [2]int{q, 0})
	visited := make(map[int]int)
	for i := 0; i < n; i++ {
		visited[i] = -1
	}
	for {
		if len(stack) == 0 {
			fmt.Println(-1)
			return
		}
		i := stack[0]
		stack = stack[1:]
		if i[0] == w {
			fmt.Println(i[1])
			if i[1] == 0 {
				return
			}
			break
		}
		for _, j := range dict[i[0]] {
			if visited[j] == -1 {
				stack = append(stack, [2]int{j, i[1] + 1})
				visited[j] = i[0]
			}
		}
	}
	arr := []string{}
	arr = append(arr, strconv.Itoa(w+1))
	i := visited[w]
	for {
		arr = append(arr, strconv.Itoa(i+1))
		if i == q {
			break
		}
		i = visited[i]
	}
	slices.Reverse(arr)
	fmt.Println(strings.Join(arr, " "))
}
