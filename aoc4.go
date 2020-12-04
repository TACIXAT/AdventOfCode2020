package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func newFields() map[string]bool {
	return map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
		"cid": false,
	}
}

func check(fields map[string]bool) bool {
	for k, v := range fields {
		if k == "cid" {
			continue
		}

		if !v {
			return false
		}
	}

	return true
}

func validate(k, v string) bool {
	switch k {

    case "byr": 
    	// (Birth Year) - four digits; at least 1920 and at most 2002.
    	byr, err := strconv.ParseInt(v, 10, 64)
    	if err != nil {
    		return false
    	}

    	if byr >= 1920 && byr <= 2002 {
    		return true
    	}
    case "iyr": 
    	// (Issue Year) - four digits; at least 2010 and at most 2020.
    	iyr, err := strconv.ParseInt(v, 10, 64)
    	if err != nil {
    		return false
    	}

    	if iyr >= 2010 && iyr <= 2020 {
    		return true
    	}
    case "eyr": 
    	// (Expiration Year) - four digits; at least 2020 and at most 2030.
    	eyr, err := strconv.ParseInt(v, 10, 64)
    	if err != nil {
    		return false
    	}

    	if eyr >= 2020 && eyr <= 2030 {
    		return true
    	}
    case "hgt": 
    	// (Height) - a number followed by either cm or in:
        // If cm, the number must be at least 150 and at most 193.
        // If in, the number must be at least 59 and at most 76.
        if len(v) < 3 {
        	return false
        }
        unit := v[len(v)-2:]

        hgt, err := strconv.ParseInt(v[:len(v)-2], 10, 64)
    	if err != nil {
    		return false
    	}

    	if unit == "cm" && hgt >= 150 && hgt <= 193 {
			return true
		}

		if unit == "in" && hgt >= 59 && hgt <= 76 {
			return true
		}
    case "hcl": 
    	// (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
    	if len(v) == 7 && v[0] == '#' {
    		for _, c := range v[1:] {
    			if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
    				return false
    			}
    		}

    		return true
    	}
    case "ecl": // (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
    	colors := []string{
    		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
    	}
    	for _, c := range colors {
    		if c == v {
    			return true
    		}
    	}
    case "pid": // (Passport ID) - a nine-digit number, including leading zeroes.
    	if len(v) == 9 {
    		return true
    	}
    case "cid": // (Country ID) - ignored, missing or not.
    	return true
	}

	return false
}

func main() {
	file, err := os.Open("input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fields := newFields()
	valid := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			if check(fields) {
				valid++
			}

			fields = newFields()
			continue
		}

		for _, field := range strings.Split(ln, " ") {
			kv := strings.Split(field, ":")
			if len(kv) != 2 {
				continue
			}

			if _, ok := fields[kv[0]]; !ok {
				continue
			}

			if !validate(kv[0], kv[1]) {
				continue
			}

			fields[kv[0]] = true
		}
	}

	if check(fields) {
		valid++
	}

	fmt.Println(valid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
