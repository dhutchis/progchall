/*
	Dylan Hutchison
	18 July 2013
	https://www.hackerrank.com/challenges/meeting-point

	Parallel Version
	The meeting point is the point closest to the average center of all the points.
	I wonder if there is a mathematical proof for this out there...

*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Pair struct {
	x, y int
}
type IndexedFloat struct {
	i    int
	dist float64
}

// globals
var N int
var pairs []Pair

func main() {
	// set to use all available CPUs
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)
	//fmt.Println(NCPU)

	// read in data & calc average
	t0 := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	var avgx, avgy = float64(0), float64(0)
	pairs = make([]Pair, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		pairs[i].x, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		pairs[i].y, _ = strconv.Atoi(scanner.Text())
		avgx += float64(pairs[i].x) / float64(N)
		avgy += float64(pairs[i].y) / float64(N)
	}
	t1 := time.Now()
	fmt.Printf("read-in %v\n", t1.Sub(t0))

	// find point closest to average center
	closech := make(chan IndexedFloat)
	for i := 0; i < NCPU; i++ {
		go calcClosest(i*N/NCPU, (i+1)*N/NCPU, avgx, avgy, closech)
	}
	// wait for everyone to finish, take best result
	idxf := <-closech
	closei, closed := idxf.i, idxf.dist
	for i := 1; i < NCPU; i++ {
		idxf = <-closech
		if idxf.dist < closed {
			closed = idxf.dist
			closei = idxf.i
		}
	}
	t2 := time.Now()
	fmt.Printf("closest %v\n", t2.Sub(t1))

	// get answer: the Manhatten distance
	manch := make(chan uint64)
	for i := 0; i < NCPU; i++ {
		go calcManSum(closei, i*N/NCPU, (i+1)*N/NCPU, manch)
	}
	// wait for everyone to finish, take best result
	mansum := uint64(0)
	for i := 0; i < NCPU; i++ {
		mansum += <-manch
	}
	t3 := time.Now()
	fmt.Printf("man dist %v\n", t3.Sub(t2))
	fmt.Print(mansum)
}

// calculate the closest
func calcClosest(i, j int, avgx, avgy float64, closech chan<- IndexedFloat) {
	closei := i
	closed := math.Sqrt((float64(pairs[i].x)-avgx)*(float64(pairs[i].x)-avgx) +
		(float64(pairs[i].y)-avgy)*(float64(pairs[i].y)-avgy))
	for i++; i < j; i++ {
		// use Manhatten sum here too instead?
		d := math.Sqrt(float64((float64(pairs[i].x)-avgx)*(float64(pairs[i].x)-avgx) +
			(float64(pairs[i].y)-avgy)*(float64(pairs[i].y)-avgy)))
		if d < closed {
			closed = d
			closei = i
		}
	}
	closech <- IndexedFloat{closei, closed}
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

// Manhatten distance is the shorter of the x distance and y distance from a to the target
func calcManSum(a, i, j int, manch chan<- uint64) {
	sum := uint64(0)
	x, y := pairs[a].x, pairs[a].y
	for ; i < j; i++ {
		sum += max(uint64(abs(pairs[i].x-x)), uint64(abs(pairs[i].y-y)))
	}
	manch <- sum
}
