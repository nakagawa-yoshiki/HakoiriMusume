package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"hakoiri-musume/puzzle"
)

type Option struct {
	stage    int   // -s
	seed     int64 // -r
	interval int   // -i
	color    bool  // -c
}

func main() {
	option := (func() Option {
		stage := flag.Int("s", 0, "stage: 0 - 7")
		seed := flag.Int64("r", 0, "random seed")
		interval := flag.Int("i", 0, "interval msec")
		color := flag.Bool("c", true, "color output")
		flag.Parse()
		return Option{*stage, *seed, *interval, *color}
	})()

	initialState := puzzle.Stages[option.stage]
	results := puzzle.Solve(initialState, option.seed)

	n := len(results)
	for i, s := range results {
		if option.interval > 0 {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

		fmt.Printf("%d/%d\n%s\n", i, n-1, s.Output(option.color))

		if option.interval > 0 {
			ms := option.interval
			if i == 0 {
				ms = 2000
			}
			time.Sleep(time.Millisecond * time.Duration(ms))
		}
	}
}
