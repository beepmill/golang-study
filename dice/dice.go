package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	roll string
)

var rollPattern = regexp.MustCompile(`(\d+)?[dD](\d+)([+-]\d+)?`)

func parseRoll(str string) (rolls, size, add int, err error) {
	var (
		r, s, a                  string
		isRolls, isSize, isBonus bool
	)
	isRolls = true

	for i, c := range str {
		fmt.Printf("Parsing position %v (%s)...\n", i, string(c))
		switch {
		case c == ' ':
			isRolls = true
			isSize = false
			isBonus = false
		case c == 'd':
			isRolls = false
			isSize = true
			isBonus = false
		case c == '+' || c == '-':
			isRolls = false
			isSize = false
			isBonus = true
			a += string(c)
		case c >= '0' && c <= '9':
			if isRolls {
				r += string(c)
			} else if isSize {
				s += string(c)
			} else if isBonus {
				a += string(c)
			} else {
				r += string(c)
			}
		default:
			log.Printf("What is %v?\n", c)
			return 0, 0, 0, fmt.Errorf("What is %v?\n", c)
		}
	}
	if r != "" {
		rolls, err = strconv.Atoi(r)
		if err != nil {
			return
		}
	} else {
		rolls = 1
	}
	if s != "" {
		size, err = strconv.Atoi(s)
		if err != nil {
			return
		}
	}
	if a != "" {
		add, err = strconv.Atoi(a)
		if err != nil {
			return
		}
	}
	return rolls, size, add, nil
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
	var roll = strings.Join(os.Args[1:], " ")
	if roll == "" {
		roll = "1d6"
	}

	if rollPattern.MatchString(roll) == true {
		result, err := Roll(roll, time.Now().UTC().UnixNano())
		if err != nil {
			panic(err)
		}
		fmt.Printf("...%v\n", result)
	} else {
		log.Panicln("Invalid roll")
	}
}
