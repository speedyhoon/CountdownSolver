package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MUL = 0
const ADD = 1
const SUB = 2
const DIV = 3
const line = "----------------"

var signs = []int{MUL, ADD, SUB, DIV}
var display = []string{"×", "+", "-", "÷"}

func rec(numbers, list, indexes []int, aim, off, used *int) {
	for _, s1 := range signs {
		for b, n2 := range numbers {
			if has(indexes, b) {
				//Each number can only be used once
				continue
			}

			total, qty := sum(list, s1, n2)
			if total == *aim && qty < *used {
				*used = qty
				*off = 0
				printa(list, s1, n2, total)
				continue
			}

			if v := int(math.Abs(float64(*aim - total))); v < *off {
				*off = v
				if *off <= 15 {
					printa(list, s1, n2, total)
				}
			}

			if qty < *used {
				rec(numbers, append(list, s1, n2), append(indexes, b), aim, off, used)
			}
		}
	}
}

func has(indexes []int, newest int) bool {
	for _, i := range indexes {
		if i == newest {
			return true
		}
	}
	return false
}

func printa(numbers []int, n ...int) {
	numbers = append(numbers, n...)
	for i, n := range numbers {
		if i >= len(numbers)-1 {
			print(" = ", n, "\n")
		} else if i%2 == 0 {
			print(n)
		} else {
			print(" ", display[n], " ")
		}
	}
}

func input() (input string, err error) {
	const readString = '\n'
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err = reader.ReadString(readString)
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, string(readString))
	return
}

func main() {
	var numbers []int
	var aim int

	fmt.Println("The numbers are: (separate each number with a space)")
	s, err := input()
	if err != nil {
		return
	}
	for _, i := range strings.Split(s, " ") {
		i = strings.TrimSpace(i)
		x, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, int(x))
	}

	fmt.Println("\nThe total to aim for is:")
	s, err = input()
	if err != nil {
		fmt.Println(err)
		return
	}

	s = strings.TrimSpace(s)
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	aim = int(a)

	fmt.Println("\n", line, "\n     ", aim, "\n", numbers, "\n", line)

	for _, n1 := range numbers {
		if n1 == aim {
			fmt.Println(n1, "=", aim)
			return
		}
	}

	off := aim

	used := len(numbers)
	for a, n1 := range numbers {
		rec(numbers, []int{n1}, []int{a}, &aim, &off, &used)
	}

	if off != 0 {
		fmt.Println(line, "\n   IMPOSSIBLE   \n", line)
	}
}

func sum(inputs []int, k ...int) (total, qty int) {
	inputs = append(inputs, k...)

	if len(inputs) < 3 {
		panic(3)
	}

	total = inputs[0]
	for n := 2; n <= len(inputs[2:])+1; n += 2 {
		switch inputs[n-1] {
		case MUL:
			total *= inputs[n]
			break
		case ADD:
			total += inputs[n]
			break
		case SUB:
			total -= inputs[n]
			break
		case DIV:
			if total != 0 && inputs[n] != 0 && total%inputs[n] == 0 {
				total /= inputs[n]
			}
		}
	}
	return total, len(inputs)/2 + 1
}
