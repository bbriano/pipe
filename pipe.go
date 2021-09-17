// Pipe alters Unix pipes.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: pipe [delay|throttle] [...]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "delay":
		delay()
	case "throttle":
		throttle()
	default:
		fmt.Fprintf(os.Stderr, "%s: not a command.\n", os.Args[1])
		os.Exit(1)
	}
}

func delay() {
	flags := flag.NewFlagSet("delay", flag.ExitOnError)
	sleepTime := flags.Duration("t", 1e9, "Sleep time before emitting lines")
	flags.Parse(os.Args[2:])

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		time.Sleep(*sleepTime)
		fmt.Println(s.Text())
	}
}

func throttle() {
	flags := flag.NewFlagSet("throttle", flag.ExitOnError)
	waitTime := flags.Duration("t", 1e9, "Wait time between emitting lines")
	flags.Parse(os.Args[2:])

	t := time.Unix(0, 0)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		if time.Since(t) < *waitTime {
			continue
		}
		fmt.Println(s.Text())
		t = time.Now()
	}
}
