package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var ErrInsufficientArgs = errors.New("Minimum of 2 arguments not met")
var ErrHelp = errors.New("Help flag found")

// MultipleSum contains all required values
type MultipleSum struct {
	Maximum  int
	Divisors []int
}

// Calc returns the sum of all multiples less than m.Maximum
func (m MultipleSum) Calc() (total int) {
	for i := 1; i < m.Maximum; i++ {
		for _, d := range m.Divisors {
			if i%d == 0 {
				total += i
				break
			}
		}
	}
	return total
}

func printUsage(output io.Writer) {
	execName, _ := os.Executable()
	execName = filepath.Base(execName)
	fmt.Fprintf(output, "Usage: %s MAX DIVISOR [DIVISOR [...]]\n", execName)
}

func processArgs(args []string) (msum MultipleSum, err error) {

	// Err on help flags
	for _, a := range args {
		if a == "-h" || a == "--help" {
			return msum, ErrHelp
		}
	}

	// Err on insufficient arguments
	if len(args) < 3 {
		return msum, ErrInsufficientArgs
	}

	// First Argument is used as maximum
	max, err := strconv.Atoi(args[1])
	if err != nil {
		return msum, err
	}
	msum.Maximum = max

	// All other arguments are divisors
	var divisors []int
	for _, a := range args[2:] {
		i, err := strconv.Atoi(a)
		if err != nil {
			return MultipleSum{}, err
		}
		divisors = append(divisors, i)
	}
	msum.Divisors = divisors

	return msum, nil
}

func main() {

	msum, err := processArgs(os.Args)
	if err != nil {
		if errors.Is(err, ErrInsufficientArgs) {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		} else if errors.Is(err, ErrHelp) {
			printUsage(os.Stdout)
			//fmt.Fprintf(os.Stderr, "%s\n", err)
		} else {
			e := err.(*strconv.NumError)
			fmt.Fprintf(os.Stderr, "Unrecognized argument: %q\n", e.Num)
		}
		os.Exit(1)
	}
	fmt.Println(msum.Calc())
}
