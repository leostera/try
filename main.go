package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var i string
	flag.StringVar(&i, "i", "1s", "")
	flag.StringVar(&i, "interval", "1s", "")

	var retry int
	flag.IntVar(&retry, "r", 10, "")
	flag.IntVar(&retry, "retries", 10, "")

	var factor float64
	flag.Float64Var(&factor, "f", 2, "")
	flag.Float64Var(&factor, "factor", 2, "")

	var version bool
	flag.BoolVar(&version, "v", false, "")
	flag.BoolVar(&version, "version", false, "")

	var usage bool
	flag.BoolVar(&usage, "h", false, "")
	flag.BoolVar(&usage, "help", false, "")

	flag.Parse()

	if version {
		die(0, Version)
	}

	command := flag.Args()

	if usage || len(command) == 0 {
		help()
		die(0, "")
	}

	if !commandExists(command) {
		die(0, "Executable not found in PATH")
	}

	factorDuration := time.Duration(factor)

	interval, err := time.ParseDuration(i)
	if err != nil {
		die(0, "Invalid interval: try 1s, 1ms, 2h45m2s")
	}

	backoff(interval, factorDuration, retry, func() {
		status := run(command)
		printStatus(status)
		if status == 0 {
			die(status, "")
		}
	})
}

func help() {
	s := `
   Usage: try [options] <cmd>

   Sample: try -i=10s -r=10 make get-deps

           Run the command up to 10 times, with the start interval of 10 seconds,
           doubling the interval on every iteration.

   Options:

     -i, --interval             start interval time (default to 1s)
     -r, --retries              amount of retries (default to 10)
     -f, --factor               multiply interval by this factor (default to 2)

     -h, --help                 this help page
     -v, --version              print out version

`
	fmt.Print(s)
}
