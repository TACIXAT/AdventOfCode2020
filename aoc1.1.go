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
    nums := make([]int64, 0)

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

        nums = append(nums, i)

        diffs[2020-i] = empty
    }

    for i, n := range nums {
        for _, m := range nums[i+1:] {
            if _, ok := diffs[n+m]; ok {
                log.Println("Found!", n*m*(2020-(n+m)))
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}