/*
	Optimized Brute Force Solution
	https://www.hackerrank.com/challenges/meeting-point
	Calculate the Meeting Point distance for each point in parallel
		Stop the calculation early if the sum exceeds the best sum so far
*/
package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

type Pair struct {
	x, y int
}

// globally shared vars among threads
var N int                              // set once
var pairs []Pair                       // set once
var best uint64 = 18446744073709551615 // writes protected by mutex; free reading
var mut sync.Mutex

func main() {
	// set to use all available CPUs
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)
	//fmt.Println(NCPU)

	// read in data
	_, _ = fmt.Scanln(&N)
	pairs = make([]Pair, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scanln(&pairs[i].x, &pairs[i].y)
	}

	// assign points to CPUs
	donech := make(chan bool)
	for i := 0; i < NCPU; i++ {
		go dosome(i*N/NCPU, (i+1)*N/NCPU, donech)
	}
	// wait for everyone to finish
	for i := 0; i < NCPU; i++ {
		<-donech
	}

	fmt.Fprint(os.Stdout, best)
}

// each cpu will calculate the Manhatten distance for each point in their assignment
// calculation stops early if it exceeds the current lowest sum
func dosome(i, j int, donech chan<- bool) {
	for i < j {
		sum := calcManSum(i)
		if sum < best { // first stage check
			mut.Lock()
			if sum < best { // ensure no one else put in a better solution before locking
				best = sum
			}
			mut.Unlock()
		}
		i++
	}
	donech <- true
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calcManSum(i int) uint64 {
	sum := uint64(0)
	x, y := pairs[i].x, pairs[i].y
	for j := 0; j < N; j++ {
		sum += max(uint64(abs(pairs[j].x-x)), uint64(abs(pairs[j].y-y)))
		if sum > best { // stop if sum exceeds best distance so far
			return 18446744073709551615
		}
	}
	return sum
}
