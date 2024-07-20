package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

func run(args []string) int {
	max := commandLine.Int("max", 255, "max value")
	name := commandLine.String("name", "", "my name")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse flags : %v\n", err)
	}

	if *max > 999 {
		fmt.Fprintf(os.Stderr, "invalid max value: %v\n", *max)
		return 1
	}
	if *name == "" {
		fmt.Fprintf(os.Stderr, "name must be provided")
		return 1
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func TestFlagVar(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{"max value", []string{"-name", "foo", "-max", "1000"}, 1},
		{"name", []string{"-name", "", "-max", "1000"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if got := run(tt.args); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
