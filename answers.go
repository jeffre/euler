package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	euler1 "github.com/jeffre/euler/001-SumMultiples"
	euler2 "github.com/jeffre/euler/002-SumEvenFibonacci"
	euler3 "github.com/jeffre/euler/003-LargePrimeFactor"
	euler4 "github.com/jeffre/euler/004-LargestPalindromeProduct"
	euler5 "github.com/jeffre/euler/005-SmallestMultiple"
	euler6 "github.com/jeffre/euler/006-SumSquareDifference"
	euler7 "github.com/jeffre/euler/007-10001stPrime"
)

const completedProblems = 7

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
		answer := solve(problem_number)
		fmt.Printf("%v\n", answer)
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
		answer := solve(i)
		fmt.Fprintf(output, "Problem %v: %v\n", i, answer)
	}
}

func solve(problemNumber int) (answer string) {

	switch problemNumber {
	case 1:
		// Answer: 233168
		return euler1.Solve()
	case 2:
		// Answer: 4613732
		return euler2.Solve()
	case 3:
		// Answer: 6857
		return euler3.Solve()
	case 4:
		// Answer: 906609
		return euler4.Solve()
	case 5:
		// Answer: 232792560
		return euler5.Solve()
	case 6:
		// Answer: 25164150
		return euler6.Solve()
	case 7:
		// Answer: 104743
		return euler7.Solve()
	default:
		return fmt.Sprintf("Problem %v has not been solved yet", problemNumber)
	}
}

func main() {
	processArgs(os.Args, os.Stdout)
}
