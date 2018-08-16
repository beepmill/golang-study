package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

var verbose = false

type rollSet struct {
	Command string
	Rolls   int
	Size    int
	Bonus   int
	Keep    int
}

type lexeme int

const (
	lRolls lexeme = iota + 1
	lSize
	lBonus
	lKeep
)

func newRollSet(command string) (rs rollSet, err error) {
	rs = rollSet{Command: command}
	err = rs.parseCommand()
	return rs, err
}

func (rs *rollSet) parseCommand() (err error) {
	var (
		str, r, s, b, k string
		lexerState      lexeme
	)
	str = rs.Command
	lexerState = lRolls

	for i, c := range str {
		if verbose == true {
			fmt.Printf("[V] Parsing position %v (%s)...\n", i, string(c))
		}
		switch {
		case c == 'd':
			lexerState = lSize
		case c == 'k':
			lexerState = lKeep
		case c == '+' || c == '-':
			lexerState = lBonus
			b += string(c)
		case c >= '0' && c <= '9':
			switch lexerState {
			case lRolls:
				r += string(c)
			case lSize:
				s += string(c)
			case lBonus:
				b += string(c)
			case lKeep:
				k += string(c)
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
			return fmt.Errorf(message)
		}
	}
	rs.Rolls, rs.Size, rs.Bonus, rs.Keep, err = validateParsedCommand(r, s, b, k)
	if err != nil {
		return
	}
	return
}

func validateParsedCommand(r, s, b, k string) (rolls, size, bonus, keep int, err error) {
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
	if k != "" {
		keep, err = strconv.Atoi(k)
		if err != nil {
			return
		}
		if keep > rolls || keep < 0 {
			err = fmt.Errorf("cannot keep %v rolls out of %v", keep, rolls)
		}
	}
	return
}

func (rs *rollSet) roll(seed int64) (result int, err error) {
	rand.Seed(seed)

	fmt.Printf("Rolling %v...\n", rs.Command)
	var results []int
	fmt.Printf("Rolled ")
	for r := 0; r < rs.Rolls; r++ {
		roll := rand.Intn(rs.Size) + 1
		fmt.Printf("%v...", roll)
		results = append(results, roll)
	}
	fmt.Println("")
	if rs.Keep > 0 && rs.Keep < rs.Rolls {
		sort.Ints(results)
		results = results[rs.Rolls-rs.Keep:]
		fmt.Printf("Keeping the highest %v rolls...%v...\n", rs.Keep, results)
	}
	if rs.Bonus != 0 {
		fmt.Printf("Modifying by %v...\n", rs.Bonus)
		results = append(results, rs.Bonus)
	}
	result = 0
	for _, r := range results {
		result += r
	}
	return
}

func printUsage() {
	fmt.Println("dice [ROLLS]d<SIZE>[(+|-)MODIFIER][k<KEEP>] [[ROLLS]d<SIZE>[(+|-)MODIFIER][k<KEEP>]]...")
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
		rs, err := newRollSet(command)
		if err != nil {
			panic(err)
		}
		result, err := rs.roll(time.Now().UTC().UnixNano())
		if err != nil {
			printUsage()
			panic(err)
		}
		fmt.Printf("...%v!\n\n", result)
	}
}
