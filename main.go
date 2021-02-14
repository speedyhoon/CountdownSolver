package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MUL = iota //MUL represents multiplication
	ADD        //ADD represents addition
	SUB        //SUB represents subtraction
	DIV        //DIV represents division
)

var (
	signs   = []int{MUL, ADD, SUB, DIV}
	display = []rune{'×', '+', '-', '÷'}
)

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
				printA(list, s1, n2, total)
				continue
			}

			if v := int(math.Abs(float64(*aim - total))); v < *off {
				*off = v
				if *off <= 15 {
					printA(list, s1, n2, total)
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

func printA(numbers []int, n ...int) {
	numbers = append(numbers, n...)
	for i, n := range numbers {
		if i >= len(numbers)-1 {
			print(" = ", n, "\n")
		} else if i%2 == 0 {
			print(n)
		} else {
			fmt.Printf(" %c ", display[n])
		}
	}
}

func input() (input string, err error) {
	const readString = '\n'
	reader := bufio.NewReader(os.Stdin)
	//ReadString will block until the delimiter is entered.
	input, err = reader.ReadString(readString)
	if err != nil {
		log.Fatalln("An error occurred while reading input. Please try again", err)
	}

	//Remove the delimiter from the string.
	input = strings.TrimSuffix(input, string(readString))
	return
}

func main() {
	log.SetFlags(0)
	log.Println("The numbers are: (separate each number with a space)")

	var numbers []int
	var aim int

	s, err := input()
	if err != nil {
		log.Fatalln(err)
	}
	for _, i := range strings.Split(s, " ") {
		i = strings.TrimSpace(i)
		x, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, int(x))
	}

	if len(numbers) < 3 {
		log.Fatalln("At least three numbers are required")
	}

	log.Println("\nThe total to aim for is:")
	s, err = input()
	if err != nil {
		log.Fatalln(err)
	}

	s = strings.TrimSpace(s)
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	aim = int(a)

	borderLen := len(fmt.Sprintf("%d", numbers)) + 4
	border := strings.Repeat("-", borderLen)
	log.Printf("\n%s\n%*d\n  %d  \n%[1]s\n", border, borderLen/2+1, aim, numbers)

	for _, n1 := range numbers {
		if n1 == aim {
			log.Println(n1, "=", aim)
			return
		}
	}

	off := aim

	used := len(numbers)
	for a, n1 := range numbers {
		rec(numbers, []int{n1}, []int{a}, &aim, &off, &used)
	}

	if off != 0 {
		log.Printf("%s\n   IMPOSSIBLE   \n%[1]s", border)
	}
}

func sum(inputs []int, k ...int) (total, qty int) {
	inputs = append(inputs, k...)

	total = inputs[0]
	for n := 2; n <= len(inputs[2:])+1; n += 2 {
		switch inputs[n-1] {
		case MUL:
			total *= inputs[n]
		case ADD:
			total += inputs[n]
		case SUB:
			total -= inputs[n]
		case DIV:
			if total != 0 && inputs[n] != 0 && total%inputs[n] == 0 {
				total /= inputs[n]
			}
		}
	}
	return total, len(inputs)/2 + 1
}
