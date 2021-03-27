package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {

	cases := map[int]Primes{
		8:            Primes{2, 2, 2},
		17:           nil,
		13195:        Primes{5, 7, 13, 29},
		600851475143: Primes{71, 839, 1471, 6857},
	}

	for n, want := range cases {
		t.Run(fmt.Sprintf("Testing %v", n), func(t *testing.T) {
			got := PrimeFactors(n)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v want %+v", got, want)
			}
		})
	}
}

func TestPrimesStruct(t *testing.T) {

	t.Run("test Primes.Last() success", func(t *testing.T) {
		p := Primes{2, 3, 5}
		got, err := p.Last()
		want := 5
		if err != nil {
			t.Fatalf("received unexpected error: %+v", err)
		}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test Primes.Last() error", func(t *testing.T) {
		p := Primes{}
		_, err := p.Last()
		want := ErrNoPrimes
		if !errors.Is(err, want) {
			t.Errorf("got %q wanted %q", err, want)
		}
	})
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactors(600851475143)
	}
}
