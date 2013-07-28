// main.go
package main

import (
	"bufio"
	"fmt"
	//"log"
	"os"
	"strconv"
	//"strings"
)

type vec [64]uint8

var DEFAULT_VEC vec

func main() {
	for i := range DEFAULT_VEC {
		DEFAULT_VEC[i] = uint8(i)
	}
	//const input = "3\n2 5 9"
	scan := bufio.NewScanner(os.Stdin)
	//scan = bufio.NewScanner(strings.NewReader(input))
	scan.Split(bufio.ScanWords)
	scan.Scan()
	//N, _ := strconv.Atoi(scan.Text())
	varr := make([]vec, 0)
	for scan.Scan() {
		d, _ := strconv.ParseUint(scan.Text(), 10, 64)
		varr = append(varr, makeVec(d))
	}

	fmt.Println(recurSubset( /*make([]vec, 0, len(varr)),*/ varr, DEFAULT_VEC))
}

// make a vector out of the input integer representing a graph
func makeVec(d uint64) (v vec) {
	conList := make([]uint8, 0)
	for i := range v {
		if d&1 != 0 {
			conList = append(conList, uint8(i))
			v[i] = conList[0]
		} else {
			v[i] = uint8(i)
		}
		d >>= 1
	}
	return v
}

// call recurSubset([], varr, DEFAULT_VEC)
// New version reuses the intermediary combined vectors
// --realized I do not even need to pass the vectors already accounted for (lock)
func recurSubset( /*lock,*/ todo []vec, vmin vec) (sum uint64) {
	//fmt.Println(len(lock), len(todo), sum)
	sum = countId(vmin)
	for i, ti := range todo {
		sum += recurSubset( /*append(lock, ti),*/ todo[i+1:],
			combineVec(vmin, ti))
	}
	return sum
}

// take minimum of two vectors
func combineVec(v1, v2 vec) (vmin vec) {
	for i := range vmin {
		if v1[i] < v2[i] {
			vmin[i] = v1[i]
		} else {
			vmin[i] = v2[i]
		}
	}
	return vmin
}

// number of entries in v where the entry == its index
func countId(v vec) (sum uint64) {
	for i, n := range v {
		if uint8(i) == n {
			sum++
		}
	}
	return sum
}
