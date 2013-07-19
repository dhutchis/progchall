/*
	Dylan Hutchison
	18 July 2013
	https://www.hackerrank.com/challenges/meeting-point

	The meeting point is the point closest to the average center of all the points.
	I wonder if there is a mathematical proof for this out there...
*/
package main

import (
	"fmt"
	"math"
)

type Pair struct {
	x, y int
}

// globals
var N int
var pairs []Pair

func main() {
	// read in data & calc average
	var avgx, avgy = float64(0), float64(0)
	_, _ = fmt.Scanln(&N)
	pairs = make([]Pair, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scanln(&pairs[i].x, &pairs[i].y)
		avgx += float64(pairs[i].x) / float64(N)
		avgy += float64(pairs[i].y) / float64(N)
	}

	// find point closest to average center
	var closei = 0
	var closed = math.Sqrt((float64(pairs[0].x)-avgx)*(float64(pairs[0].x)-avgx) +
		(float64(pairs[0].y)-avgy)*(float64(pairs[0].y)-avgy))
	for i := 1; i < N; i++ {
		d := math.Sqrt(float64((float64(pairs[i].x)-avgx)*(float64(pairs[i].x)-avgx) +
			(float64(pairs[i].y)-avgy)*(float64(pairs[i].y)-avgy)))
		if d < closed {
			closed = d
			closei = i
		}
	}

	// get answer: the Manhatten distance
	fmt.Print(calcManSum(closei))
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

// Manhatten distance is the shorter of the x distance and y distance
func calcManSum(i int) uint64 {
	sum := uint64(0)
	x, y := pairs[i].x, pairs[i].y
	for j := 0; j < N; j++ {
		sum += max(uint64(abs(pairs[j].x-x)), uint64(abs(pairs[j].y-y)))
	}
	return sum
}
