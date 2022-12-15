package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	raw, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer raw.Close()

	var prio_list string = "abcdefghijklmnopqrstuvwqyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var prio_sum int

	scanner := bufio.NewScanner(raw)

	// do something
	for scanner.Scan() {
		rucksack := scanner.Text()
		rucksack_end := len(rucksack)
		rucksack_middle := len(rucksack) / 2
		comp1 := rucksack[0:rucksack_middle]
		comp2 := rucksack[rucksack_middle:rucksack_end]

		// Compare every char in compartment 1 with the chars in compartment 2, break when duplicate is found
		for _, item := range comp1 {
			if strings.ContainsAny(comp2, string(item)) {
				prio_value := strings.Index(prio_list, string(item)) + 1 // Index + 1 get's the correct prio value from list
				prio_sum = prio_sum + prio_value
				break
			}
		}
	}

	fmt.Println(prio_sum)
}
