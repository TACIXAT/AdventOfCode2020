package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Requirement struct {
	Name  string
	Count int
}

func countBags(parents map[string][]Requirement, name string) int {
	sum := 1

	for _, p := range parents[name] {
		sum += p.Count * countBags(parents, p.Name)
	}

	return sum
}

func main() {
	file, err := os.Open("input7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parents := make(map[string][]Requirement)
	children := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}

		toks := strings.Split(ln, " contain ")
		parent := toks[0]
		stored := strings.Split(toks[1], ", ")

		for _, child := range stored {
			child = strings.Trim(child, ". ")
			if !strings.HasSuffix(child, "s") {
				child = child + "s"
			}

			if child == "no other bags" {
				continue
			}

			toks := strings.Split(child, " ")
			name := strings.Trim(child, "123456789 ")
			count, err := strconv.ParseInt(toks[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			creq := Requirement{
				Name:  name,
				Count: int(count),
			}

			children[name] = append(children[name], parent)
			parents[parent] = append(parents[parent], creq)
		}
	}

	count := make(map[string]bool)
	toCheck := make(chan string, 1000)
	toCheck <- "shiny gold bags"

	for elem := range toCheck {
		for _, p := range children[elem] {
			count[p] = true
			toCheck <- p
		}

		if len(toCheck) == 0 {
			break
		}
	}

	log.Println("out", len(count), countBags(parents, "shiny gold bags")-1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
