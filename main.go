package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	problems, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading problems")
	}

	/*for _, problem := range problems {
		fmt.Println(problem[0], problem[1])
	}*/

	var result string

	var points int

	fmt.Println(len(problems))

	fmt.Println("Let's start the quiz ...")
	for i := 0; i < len(problems); i++ {
		fmt.Println(problems[i][0])
		fmt.Scan(&result)
		if result == problems[i][1] {
			fmt.Println("true")
			points++
		} else {
			fmt.Println("false")
		}
	}
	fmt.Printf("You scored: %v/%v \n", points, len(problems))
}
