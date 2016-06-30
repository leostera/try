package main

import (
	"flag"
	"fmt"
	omg "github.com/ostera/oh-my-gosh/lib"
	"time"
)

// Version string prefilled at build time
var (
	Version string

  i string
  retry int
  factor float64
  version bool
  usage bool
)

func main() {
	flag.StringVar(&i, "i", "1s", "")
	flag.StringVar(&i, "interval", "1s", "")

	flag.IntVar(&retry, "r", 10, "")
	flag.IntVar(&retry, "retries", 10, "")

	flag.Float64Var(&factor, "f", 2, "")
	flag.Float64Var(&factor, "factor", 2, "")

	flag.BoolVar(&version, "v", false, "")
	flag.BoolVar(&version, "version", false, "")

	flag.BoolVar(&usage, "h", false, "")
	flag.BoolVar(&usage, "help", false, "")

	flag.Parse()

	command := flag.Args()

	if version {
		omg.Die(0, Version)
	}

	if usage || len(command) == 0 {
		help()
		omg.Die(0, "")
	}

	if !omg.CommandExists(command) {
		omg.Die(0, "Executable not found in PATH")
	}

	interval, err := time.ParseDuration(i)
	if err != nil {
		omg.Die(0, "Invalid interval: try 1s, 1ms, 2h45m2s")
	}

	backoff(interval, factor, retry, func() {
		status := omg.Run(command)
		omg.PrintStatus(status)
		if status == 0 {
			omg.Die(status, "")
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

func backoff(d time.Duration, f float64, r int, fn func()) {
  for ; r > 0; r-- {
    fn()
    time.Sleep(d)
    d = time.Duration( float64(d) * f )
  }
}
