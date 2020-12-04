package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func xor(a, b bool) bool {
	if a && !b || !a && b {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid := 0
	invalid := 0

	valid2 := 0
	invalid2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}

		low := 0
		high := 0
		chr := ' '
		pwd := ""
		fmt.Sscanf(ln, "%d-%d %c: %s", &low, &high, &chr, &pwd)
		fmt.Println(low, high, string(chr), pwd)

		count := 0
		for _, c := range pwd {
			if c == chr {
				count += 1
			}
		}

		if count < low || count > high {
			invalid++
		} else {
			valid++
		}

		if xor(pwd[low-1] == byte(chr), pwd[high-1] == byte(chr)) {
			valid2++
		} else {
			invalid2++
		}
	}

	fmt.Println("Valid:", valid)
	fmt.Println("Invalid:", invalid)

	fmt.Println("Valid2:", valid2)
	fmt.Println("Invalid2:", invalid2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
