package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	MUL              = iota // MUL represents multiplication.
	ADD                     // ADD represents addition.
	SUB                     // SUB represents subtraction.
	DIV                     // DIV represents division.
	aimPercentage    = 10
	defaultThreshold = 15 // Answers won't be printed unless they are within this threshold.
)

var (
	signs   = []int{MUL, ADD, SUB, DIV}
	display = []rune{'×', '+', '-', '÷'}
)

func recurse(numbers, list, indexes []int, target, offBy, used, threshold *int) {
	for s1 := range signs {
		for n2 := range numbers {
			if has(indexes, n2) {
				// Each number can only be used once.
				continue
			}

			total, qty := sum(list, signs[s1], numbers[n2])
			if total == *target {
				switch {
				case *offBy != 0:
					log.Println("Solved answer. Attempting to simplify further...")
					*offBy = 0
					fallthrough
				case qty < *used:
					*used = qty
					printA(list, signs[s1], numbers[n2], total)
				}

				//We don't care about the rest of the iterations because we want to find a simpler answer.
				return
			}

			if v := int(math.Abs(float64(*target - total))); v < *offBy {
				*offBy = v
				if *offBy <= *threshold {
					printA(list, signs[s1], numbers[n2], total)
				}
			}

			if *offBy == 0 && qty+1 < *used || *offBy != 0 && qty < *used {
				recurse(numbers, append(list, signs[s1], numbers[n2]), append(indexes, n2), target, offBy, used, threshold)
			}
		}
	}
}

func has(haystack []int, needle int) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func printA(numbers []int, n ...int) {
	numbers = append(numbers, n...)
	for i := range numbers {
		switch {
		case i >= len(numbers)-1:
			println(" = ", numbers[i])
		case i%2 == 0:
			print(numbers[i])
		default:
			fmt.Printf(" %c ", display[numbers[i]])
		}
	}
}

func input() (input string, err error) {
	const readString = '\n'
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered.
	input, err = reader.ReadString(readString)
	if err != nil {
		return
	}

	// Remove the delimiter from the string.
	input = strings.TrimSuffix(input, string(readString))
	return
}

func parse(minQtyRequired int, message string) (numbers []int) {
	for len(numbers) < minQtyRequired {
		s, err := input()
		if err != nil {
			log.Println(err)
			continue
		}

		strs := strings.Split(s, " ")
		for i := range strs {
			strs[i] = strings.TrimSpace(strs[i])
			if strs[i] == "" {
				continue
			}

			var x int64
			x, err = strconv.ParseInt(strs[i], 10, 64)
			if err != nil || x == 0 {
				// Zeros & errors are ignored.
				continue
			}
			numbers = append(numbers, int(x))
		}

		if len(numbers) < minQtyRequired {
			log.Printf("%s\n\n", message)
			numbers = nil // Clear any previous failed inputs.
		}
	}

	return
}

func solve(numbers []int, target, threshold int) bool {
	offBy := target
	used := len(numbers)
	for i := range numbers {
		recurse(numbers, []int{numbers[i]}, []int{i}, &target, &offBy, &used, &threshold)
	}

	return offBy != 0
}

func main() {
	log.SetFlags(0)
	log.Println("The numbers are: (separate each number with a space)")
	numbers := parse(3, "At least three numbers are required.")
	sort.Ints(numbers)

	log.Println("\nThe total to aim for is:")
	target := parse(1, "At least one number is required.")[0]

	threshold := calcThreshold(target/aimPercentage, defaultThreshold)

	borderLen := len(fmt.Sprintf("%v", numbers)) + 4
	border := strings.Repeat("-", borderLen)
	log.Printf("\n%s\n%*d\n  %d  \n%[1]s\n", border, borderLen/2+1, target, numbers)

	for i := range numbers {
		if numbers[i] == target {
			log.Println(numbers[i], "=", target)
			return
		}
	}

	if solve(numbers, target, threshold) {
		log.Printf("%s\n%*sIMPOSSIBLE\n%[1]s", border, borderLen/2-5-borderLen%2, " ")
	}
}

func calcThreshold(n ...int) (max int) {
	for i := range n {
		if n[i] > max {
			max = n[i]
		}
	}
	return
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
