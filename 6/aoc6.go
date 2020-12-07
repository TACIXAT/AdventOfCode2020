package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	groupCount := 0
	set := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			for _, v := range set {
				if v == groupCount {
					count += 1
				}
			}
			set = make(map[string]int)
			groupCount = 0
			continue
		}

		groupCount += 1
		for _, ch := range ln {
			set[string(ch)] += 1
		}
	}

	for _, v := range set {
		if v == groupCount {
			count += 1
		}
	}

	log.Println("Out", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
