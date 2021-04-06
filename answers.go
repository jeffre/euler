package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	euler1 "github.com/jeffre/euler/001-SumMultiples"
	euler2 "github.com/jeffre/euler/002-SumEvenFibonacci"
	euler3 "github.com/jeffre/euler/003-LargePrimeFactor"
	euler4 "github.com/jeffre/euler/004-LargestPalindromeProduct"
	euler5 "github.com/jeffre/euler/005-SmallestMultiple"
	euler6 "github.com/jeffre/euler/006-SumSquareDifference"
	euler7 "github.com/jeffre/euler/007-10001stPrime"
	euler8 "github.com/jeffre/euler/008-LargestProductInSeries"
	euler9 "github.com/jeffre/euler/009-SpecialPythagoreanTriplet"
	euler10 "github.com/jeffre/euler/010-SumOfPrimes"
)

const completedProblems = 10

var (
	ErrNotIntOrFloat = errors.New("solution was not type int or float")
	ErrNotSolved     = errors.New("solution has not been solved yet")
)

func processArgs(args []string, output io.Writer) {

	// Print usage and quit
	for _, a := range args {
		if a == "-h" || a == "--help" {
			printUsage(output)
			return
		}
	}

	switch len(args) {
	case 1:
		solveAll(output)
	case 2:
		problem_number, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(output, "Error: %q", err)
			return
		}
		answer, err := solve(problem_number)
		if err != nil {
			fmt.Fprintf(output, "Error: %v\n", err)
		}
		fmt.Fprintln(output, answer)
	default:
		fmt.Fprintln(output, "Invalid arguments")
	}
}

func printUsage(output io.Writer) {
	execName, _ := os.Executable()
	execName = filepath.Base(execName)
	fmt.Fprintf(output, "Usage: %s [PROBLEM_NUMBER]\n", execName)
}

func solveAll(output io.Writer) {
	for i := 1; i <= completedProblems; i++ {
		answer, err := solve(i)
		if err != nil {
			fmt.Fprintf(output, "Error solving problem %v: %v\n", i, err)
		}
		fmt.Fprintf(output, "Problem %v: %v\n", i, answer)
	}
}

func solve(n int) (a interface{}, err error) {

	switch n {
	case 1:
		a = euler1.Solve()
	case 2:
		a = euler2.Solve()
	case 3:
		a = euler3.Solve()
	case 4:
		a = euler4.Solve()
	case 5:
		a = euler5.Solve()
	case 6:
		a = euler6.Solve()
	case 7:
		a = euler7.Solve()
	case 8:
		a = euler8.Solve()
	case 9:
		a = euler9.Solve()
	case 10:
		a = euler10.Solve()
	default:
		return nil, ErrNotSolved
	}

	// Ensure answer is either an integer or float
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int:
		return a, nil
	case reflect.Float32:
		return a, nil
	case reflect.Float64:
		return a, nil
	default:
		return nil, ErrNotIntOrFloat
	}
}

func main() {
	processArgs(os.Args, os.Stdout)
}
