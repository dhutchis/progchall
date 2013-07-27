package main

import (
	"fmt"
)

// Mathematica code to find Marr:
// Do[Print[ Sum[Binomial[n - 3*H, H], {H, 0, Floor[n/4]}]], {n, 1, 40}]
var Marr = [40]int{
	1,
	1,
	1,
	2,
	3,
	4,
	5,
	7,
	10,
	14,
	19,
	26,
	36,
	50,
	69,
	95,
	131,
	181,
	250,
	345,
	476,
	657,
	907,
	1252,
	1728,
	2385,
	3292,
	4544,
	6272,
	8657,
	11949,
	16493,
	22765,
	31422,
	43371,
	59864,
	82629,
	114051,
	157422,
	217286}
var Parr [40]int

func main() {
	sieve()
	for i := range Marr {
		//fmt.Println(i+1, "\t", Marr[i], "\t", Parr[i])
		fmt.Print(Parr[i], ",")
	}
}

// computes Parr
func sieve() {
	const limit = 217287 // means sieve numbers < 201

	// Sieve of Eratosthenes -- mark composites as true
	// Based on: http://rosettacode.org/mw/index.php?title=Sieve_of_Eratosthenes&oldid=165120#Go
	c := make([]bool, limit) // c for composite.  false means prime candidate
	c[1] = true              // 1 not considered prime
	p := 2
	for {
		// first allowed optimization:  outer loop only goes to sqrt(limit)
		p2 := p * p
		if p2 >= limit {
			break
		}
		// second allowed optimization:  inner loop starts at sqr(p)
		for i := p2; i < limit; i += p {
			c[i] = true // it's a composite

		}
		// scan to get next prime for outer loop
		for {
			p++
			if !c[p] {
				break
			}
		}
	}

	// finish writing primes
	np := 0  // number of primes so far
	num := 2 // current number
	for midx := 3; midx < 40; midx++ {
		for ; num <= Marr[midx]; num++ {
			if !c[num] { // if not prime
				np++ // increment number of primes
			}
		}
		Parr[midx] = np
	}

	/*
		// sieve complete.  now print a representation.
		for n := 1; n < limit; n++ {
			if c[n] {
				fmt.Print("  .")
			} else {
				fmt.Printf("%3d", n)
			}
			if n%20 == 0 {
				fmt.Println("")
			}
		}
	*/
}
