package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestUsage(t *testing.T) {
	buffer := bytes.Buffer{}
	printUsage(&buffer)
	got := buffer.String()
	want := "Usage:"
	if !strings.HasPrefix(got, want) {
		t.Fatalf("%q does not start with %q", got, want)
	}
}

func TestHelpFlags(t *testing.T) {

	usageBuffer := bytes.Buffer{}
	printUsage(&usageBuffer)
	want := usageBuffer.String()

	cases := [][]string{
		{"foo", "1", "2", "-h"},
		{"foo", "--help"},
	}
	for _, args := range cases {
		t.Run(fmt.Sprintf("Testing %v flag", args[len(args)-1]), func(t *testing.T) {
			buffer := bytes.Buffer{}
			processArgs(args, &buffer)
			got := buffer.String()

			if got != want {
				t.Errorf("got %q wanted %q", got, want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	cases := map[int]interface{}{
		1:  233168,
		2:  4613732,
		3:  6857,
		4:  906609,
		5:  232792560,
		6:  25164150,
		7:  104743,
		8:  23514624000,
		9:  31875000,
		10: 142913828922,
	}
	for id, want := range cases {
		got, _ := solve(id)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestCompletedProblems(t *testing.T) {
	t.Run("non-existent problem", func(t *testing.T) {
		id := completedProblems + 1
		_, err := solve(id)
		if err != ErrNotSolved {
			t.Errorf("Got %q want %q", err, ErrNotSolved)
		}
	})
}
