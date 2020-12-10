package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	size := 25
	nums := make([]int, size)
	idx := 0
	invalid := -1

	running := 0
	base := 0
	target := 1930745883
	// target = 127
	all := make([]int, 0)
	sum := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}

		n, err := strconv.ParseInt(ln, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		running += int(n)
		all = append(all, int(n))

		for sum < 0 && running > target {
			running -= all[base]
			base++
		}

		if sum < 0 && running == target {
			min := 1 << 62
			max := 0
			for _, c := range all[base:] {
				if c < min {
					min = c
				}

				if c > max {
					max = c
				}
			}

			log.Println("FOUND!", min+max)
			sum = min + max
		}

		if idx >= size {
			diffs := make(map[int]struct{})
			valid := false
			for _, cmp := range nums {
				diff := int(n) - cmp
				// log.Println(diff, cmp)

				if _, ok := diffs[cmp]; ok {
					valid = true
					break
				}

				diffs[diff] = struct{}{}
			}

			if !valid {
				// log.Println("INVALID", n)
				invalid = int(n)
			} else {
				// log.Println("VALID", n)
			}
		} else {
			// log.Println("PRE", n)
		}

		nums[idx%size] = int(n)
		idx++

		if invalid > -1 {
			break
		}
	}

	log.Println("out", invalid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
