package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var target = 2020
	numbers := ReadFile()
	sort.Ints(numbers)

	for _, number := range numbers {
		numberToFind := target - number
		if index := sort.SearchInts(numbers, numberToFind); numberToFind == numbers[index] {
			fmt.Println(fmt.Sprintf("The two numbers adding up to %d are %d and %d", target, number, numberToFind))
			fmt.Println(fmt.Sprintf("The answer to enter is %d", number*numberToFind))
			break
		}
	}
}

func ReadFile() []int {
	var numbers []int
	file, err := os.Open("Day1/input")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}
	return numbers
}
