package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pos := []int{0, 0, 0, 0, 0}
	trees := []int{0, 0, 0, 0, 0}
	right := []int{1, 3, 5, 7, 1}
	down := []int{1, 1, 1, 1, 2}

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}

		for i := range pos {
			if row%down[i] == 0 {
				if ln[pos[i]%len(ln)] == '#' {
					trees[i]++
				}

				pos[i] += right[i]
			}
		}

		row++
	}

	acc := 1
	for _, n := range trees {
		acc *= n
	}
	fmt.Println(acc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
