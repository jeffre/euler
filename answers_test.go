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
