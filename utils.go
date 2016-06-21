package main

import (
	"fmt"
	"os"
	"time"
)

func die(status int, message string) {
	if len(message) > 0 {
		fmt.Println(message)
	}
	os.Exit(status)
}

func backoff(d time.Duration, f time.Duration, r int, fn func()) {
  repeat:
    if r == 0 { return }
    fn()
    time.Sleep(d)
    d = d*f
    r--
  goto repeat
}
