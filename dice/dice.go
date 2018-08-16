package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var verbose = false

type rollSet struct {
	Rolls int
	Size  int
	Bonus int
}

type lexeme int

const (
	lRolls lexeme = iota + 1
	lSize
	lBonus
)

func newRollSetFromStrings(r, s, b string) (rs rollSet, err error) {
	var rolls, size, bonus int
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
	if b != "" {
		bonus, err = strconv.Atoi(b)
		if err != nil {
			return
		}
	}
	return rollSet{rolls, size, bonus}, nil
}

func parseCommand(str string) (rs rollSet, err error) {
	var (
		r, s, b    string
		lexerState lexeme
	)
	lexerState = lRolls

	for i, c := range str {
		if verbose == true {
			fmt.Printf("Parsing position %v (%s)...\n", i, string(c))
		}
		switch {
		case c == 'd':
			lexerState = lSize
		case c == '+' || c == '-':
			lexerState = lBonus
			b += string(c)
		case c >= '0' && c <= '9':
			if lexerState == lRolls {
				r += string(c)
			} else if lexerState == lSize {
				s += string(c)
			} else if lexerState == lBonus {
				b += string(c)
			} else {
				r += string(c)
			}
		default:
			message := "error parsing command "
			if i != 0 {
				message += str[:i]
			}
			message += ">>>" + string(c) + "<<<"
			if i != len(str)-1 {
				message += str[i:]
			}
			if err != nil {
				return
			}
			return rollSet{}, fmt.Errorf(message)
		}
	}
	rs, err = newRollSetFromStrings(r, s, b)
	if err != nil {
		return
	}
	return
}

func (rs *rollSet) roll(seed int64) (result int, err error) {
	rand.Seed(seed)

	for r := 0; r < rs.Rolls; r++ {
		roll := rand.Intn(rs.Size) + 1
		fmt.Printf("Rolled %v...\n", roll)
		result += roll
	}
	if rs.Bonus != 0 {
		fmt.Printf("Bonus %v\n", rs.Bonus)
		result += rs.Bonus
	}
	return
}

func printUsage() {
	log.Println("dice rollsDsize(+|-)bonus [rollsDsize(+|-)bonus]...")
	flag.PrintDefaults()
}

func main() {
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.Parse()
	if flag.NArg() == 0 {
		printUsage()
		os.Exit(1)
	}

	for _, command := range flag.Args() {
		rs, err := parseCommand(command)
		if err != nil {
			log.Panic(err)
		}
		result, err := rs.roll(time.Now().UTC().UnixNano())
		if err != nil {
			printUsage()
			panic(err)
		}
		fmt.Printf("...%v\n", result)
	}
}
