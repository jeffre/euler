package main

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestUsage(t *testing.T) {
	buffer := bytes.Buffer{}
	printUsage(&buffer)
	got := buffer.String()
	want := "Usage:"
	if !strings.HasPrefix(got, want) {
		t.Errorf("%q does not start with %q", got, want)
	}
}

func TestProcessArgs(t *testing.T) {

	t.Run("Test insufficient args", func(t *testing.T) {
		args := []string{"foo"}
		_, err := processArgs(args)
		if !errors.Is(err, ErrInsufficientArgs) {
			t.Errorf("got %q wanted %q", err, ErrInsufficientArgs)
		}
	})

	t.Run("Test -h flag", func(t *testing.T) {
		args := []string{"foo", "1", "2", "-h"}
		_, err := processArgs(args)
		if !errors.Is(err, ErrHelp) {
			t.Errorf("got %q wanted %q", err, ErrHelp)
		}
	})
	t.Run("Test --help flag", func(t *testing.T) {
		args := []string{"foo", "--help"}
		_, err := processArgs(args)
		if !errors.Is(err, ErrHelp) {
			t.Errorf("got %q wanted %q", err, ErrHelp)
		}
	})

	t.Run("Test MultipleSum struct", func(t *testing.T) {
		args := []string{"foo", "10", "3", "5"}
		got, err := processArgs(args)
		if err != nil {
			t.Fatalf("Did not expect error: %+v", err)
		}
		want := MultipleSum{
			Maximum: 10,
			Divisors: []int{
				3, 5,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("Test MultipleSum.Calc()", func(t *testing.T) {
		args := []string{"foo", "10", "3", "5"}
		msum, err := processArgs(args)
		if err != nil {
			t.Fatalf("Did not expect error: %+v", err)
		}
		got := msum.Calc()
		want := 23
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test MultipleSum.Calc() with overlapping divisors", func(t *testing.T) {
		args := []string{"foo", "10", "3", "5", "6"}
		msum, err := processArgs(args)
		if err != nil {
			t.Fatalf("Did not expect error: %+v", err)
		}
		got := msum.Calc()
		want := 23
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
