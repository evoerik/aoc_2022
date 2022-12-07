package main

import (
	"log"
	"fmt"
	"os"
	"bufio"
	)

func main()  {
	raw, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer raw.Close()

	var my_total_score int = 0 
	scanner := bufio.NewScanner(raw)

	// Select my and opponents move and store away for each around
	for scanner.Scan() {
		round := scanner.Text()
		opponent_move := round[0:1]
		my_move := round[2:3]
		
		// Compary my move to opponent and add points accordingly
		switch my_move {
		case "X":
			switch opponent_move {
			case "A":
				my_total_score = my_total_score + 3
			case "B":
				my_total_score = my_total_score + 0
			case "C":
				my_total_score = my_total_score + 6
			}
			my_total_score = my_total_score + 1
		case "Y":
			switch opponent_move {
			case "A":
				my_total_score = my_total_score + 6
			case "B":
				my_total_score = my_total_score + 3
			case "C":
				my_total_score = my_total_score + 0
			}
			my_total_score = my_total_score + 2
		case "Z":
			switch opponent_move {
			case "A":
				my_total_score = my_total_score + 0
			case "B":
				my_total_score = my_total_score + 6
			case "C":
				my_total_score = my_total_score + 3
			}
			my_total_score = my_total_score + 3
		}
	}

	fmt.Println(my_total_score)

}