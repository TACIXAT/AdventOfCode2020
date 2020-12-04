package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    diffs := make(map[int64]struct{})
    empty := struct{}{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        ln := scanner.Text()
        if len(ln) == 0 {
        	continue
        }

        i, err := strconv.ParseInt(ln, 10, 64)
        if err != nil {
        	log.Fatal(err)
        }

        if _, ok := diffs[i]; ok {
        	log.Println("Found!", (2020-i)*i)
        }

        diffs[2020-i] = empty
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}