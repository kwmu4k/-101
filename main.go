package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("main.tt")
	handle(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line int
	var stack []int
	var buffer byte

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		instruction := parts[0]

		var arg1 int
		
		if len(parts) > 1 { 
			arg1, _ = strconv.Atoi(parts[1])
		}

		switch instruction {
		case "->": 
			stack = append(stack, arg1)

		case "<-": 
			stack = stack[:len(stack)-1]

		case "~>": 
			buffer = byte(stack[len(stack)-1])
			_ = buffer
			stack = stack[:len(stack)-1]

		case "<~":
			stack = append(stack, int(buffer))
			buffer = 0

		case ":;":
			sort.Sort(sort.Reverse(sort.IntSlice(stack)))

		case "<<":
			var b [1]byte
			_, err := os.Stdin.Read(b[:])
			handle(err)
			stack = append(stack, int(b[0]))

		case ">>":
			fmt.Println(stack[len(stack)-1])
			stack = stack[:len(stack)-1]

		case "++":
			stack[len(stack)-1]++

		case "--":
			stack[len(stack)-1]--

		case "<>":
			line = arg1
		
		case "??":
			if int(buffer) != stack[len(stack)-1] {
				line += arg1
			}

		case "!!":
			fmt.Printf("\033[33m\n%d \033[32m[%d]", stack, buffer)
			os.Exit(0)
		}

		line++
	}
}