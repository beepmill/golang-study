package main

import (
	"regexp"
	"flag"
	"fmt"
)

var (
	roll string
)

// Roll TODO
func Roll(command string) error {
	fmt.Printf("Rolling %v", command)
	return nil
}

func main() {
	var roll = flag.String("roll", "1d6", "help message goes here")
	flag.Parse()
	
	rollValid := regexp.MustCompile("\\d+[dD]\\d+")
	rollValid.MatchString(*roll)

	Roll(*roll)
}
