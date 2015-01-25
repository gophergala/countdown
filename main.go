package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	deadline := os.Args[1]
	fmt.Printf("deadline: %q\n", deadline)
	date, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %T{%v}\n", err, err)
		fmt.Fprintf(os.Stdout, "Deadline has to be in RFC3339 format, i.e.: '%s'\n", time.RFC3339)
		os.Exit(1)
	}
	ticker := time.Tick(1 * time.Second)
	for t := range ticker {
		remaining := date.Sub(t)
		if remaining.Seconds() < 0.0 {
			fmt.Fprintf(os.Stdout, "\n>>> Deadline reached. Exiting.\n")
			os.Exit(0)
		}
		fmt.Fprintf(os.Stdout, "\r>>> T minus %d seconds", int(math.Ceil(remaining.Seconds())))
	}
}
