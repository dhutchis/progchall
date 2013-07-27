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

func main() {
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

	fmt.Println(doSumSubsets(varr))
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

// generate all subsets and get the big sum
func doSumSubsets(vecs []vec) (sum uint64) {
	// for each subset
	for mask := uint64(1); mask < uint64(1)<<uint(len(vecs)); mask++ {
		// build the subset
		vsubset := make([]vec, 0, len(vecs)) // capacity can be optimized here: log(mask)
		m := mask
		for i := range vecs {
			if m&1 == 1 {
				vsubset = append(vsubset, vecs[i])
			}
			m >>= 1
		}
		vmin := combineVec(vsubset)
		sum += countId(vmin)
	}
	return sum + 64 // +64 for the empty subset
}

// take minimum of all vectors combined
func combineVec(varr []vec) (vmin vec) {
	for i := range vmin {
		min := uint8(i)
		for j := range varr {
			if varr[j][i] < min {
				min = varr[j][i]
			}
		}
		vmin[i] = min
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
