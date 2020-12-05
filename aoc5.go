package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	max := 0
	ids := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}

		minRow := 0
		maxRow := 128
		for _, ch := range ln[:7] {
			rng := maxRow - minRow
			if ch == 'F' {
				maxRow = minRow + rng / 2
			} else if ch == 'B' {
				minRow = minRow + rng / 2
			}
		}

		minCol := 0
		maxCol := 8
		for _, ch := range ln[7:] {
			rng := maxCol - minCol
			if ch == 'L' {
				maxCol = minCol + rng / 2
			} else if ch == 'R' {
				minCol = minCol + rng / 2
			}
		}

		id := minRow * 8 + minCol
		if id > max {
			max = id
		}

		ids = append(ids, id)
	}

	sort.Ints(ids)
	last := 0
	for _, id := range ids {
		if id - last == 2 {
			log.Println("Found", id-1)
		}
		last = id
	}

	log.Println("Max", max)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
