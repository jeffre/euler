package euler007

import "strconv"

func PrimeByPos(pos int) int {
	max := 999999
	p := make([]bool, max)

	p[0] = true
	p[1] = true
	for i := 2; i <= max; i++ {

		for j := 2; i*j < max; j++ {
			p[i*j] = true
		}

	}

	primes := []int{2}
	for id, val := range p {
		if !val {
			primes = append(primes, id)
		}
	}

	return primes[pos]
}

func Solve() string {
	return strconv.Itoa(PrimeByPos(10001))
}
