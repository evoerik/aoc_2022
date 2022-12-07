package main

import (
	"log"
	"fmt"
	"os"
	"bufio"
	"strconv"
	)

func main()  {
	raw, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer raw.Close()

	var inventories []int
	var inventory int = 0
	var most_calories int = 0

	scanner := bufio.NewScanner(raw)

	// Create list containing each elfs total calories
	for scanner.Scan() {
		if len(scanner.Text()) != 0  {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			inventory = inventory + i
			//fmt.Println(inventory)

		} else {
			inventories = append(inventories, inventory)
			inventory = 0 
		}
	}

	// Find the elf inventory with the most calories
	for j := 0; j < len(inventories); j++ {
		if inventories[j] > most_calories {
			most_calories = inventories[j]
		}
	}

	fmt.Println(most_calories)
	
}