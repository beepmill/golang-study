package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var (
	roll string
)

var rollPattern = regexp.MustCompile(`(\d+)[dD](\d+)([+-]\d+)?`)

func parseRoll(s string) (rolls, size, add int, err error) {
	if rollPattern.MatchString(s) {
		match := rollPattern.FindStringSubmatch(s)
		fmt.Printf("Rolling %v (%s rolls of size %s), then adding %s\n", s, match[1], match[2], match[3])
		rolls, err = strconv.Atoi(match[1])
		if err != nil {
			return
		}
		size, err = strconv.Atoi(match[2])
		if err != nil {
			return
		}
		if len(match) == 4 && match[3] != "" {
			add, err = strconv.Atoi(match[3])
			if err != nil {
				return
			}
		}
	} else {
		err = fmt.Errorf("Cannot parse command '%v'", s)
	}
	return
}

// Roll TODO
func Roll(command string, seed int64) (result int, err error) {
	rand.Seed(seed)

	rolls, size, add, err := parseRoll(command)
	if err != nil {
		return
	}

	for r := 0; r < rolls; r++ {
		roll := rand.Intn(size) + 1
		fmt.Printf("Rolled %v...\n", roll)
		result += roll
	}
	if add != 0 {
		fmt.Printf("Adding %v\n", add)
		result += add
	}
	return
}

func main() {
	var roll = flag.String("roll", "1d6", "help message goes here")
	flag.Parse()

	if rollPattern.MatchString(*roll) == true {
		result, err := Roll(*roll, time.Now().UTC().UnixNano())
		if err != nil {
			panic(err)
		}
		fmt.Printf("...%v", result)
	} else {
		panic("Invalid roll")
	}
}
