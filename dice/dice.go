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

// Roll TODO
func Roll(command string) error {
	rand.Seed(time.Now().UTC().UnixNano())
	rollValid := regexp.MustCompile(`(\d+)[dD](\d+)`)
	match := rollValid.FindStringSubmatch(command)
	fmt.Printf("Rolling %v (%s rolls of size %s)\n", command, match[1], match[2])
	result := 0
	rolls, err := strconv.Atoi(match[1])
	size, err := strconv.Atoi(match[2])

	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for r := 0; r < rolls; r++ {
		roll := rand.Intn(size) + 1
		result += roll
	}
	fmt.Printf("%v", result)
	return nil
}

func main() {
	var roll = flag.String("roll", "1d6", "help message goes here")
	flag.Parse()

	rollValid := regexp.MustCompile(`\d+[dD]\d+`)
	if rollValid.MatchString(*roll) == true {
		Roll(*roll)
	} else {
		panic("Invalid roll")
	}
}
