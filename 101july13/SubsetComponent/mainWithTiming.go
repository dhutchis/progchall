// main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type vec [64]uint8

func main() {
	const input = "10\n2 5 9 15651 351634883 5634164 48636 48 16315 5651 856 156 1561 8989 7853 24648 89410 486461"
	scan := bufio.NewScanner(os.Stdin)
	scan = bufio.NewScanner(strings.NewReader(input))
	scan.Split(bufio.ScanWords)
	scan.Scan()
	//N, _ := strconv.Atoi(scan.Text())
	t0 := time.Now()
	varr := make([]vec, 0)
	for scan.Scan() {
		d, _ := strconv.ParseUint(scan.Text(), 10, 64)
		varr = append(varr, makeVec(d))
	}
	t1 := time.Now()
	log.Println("Making vectors: ", t1.Sub(t0))

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
	tBase, tBuild, tMin, tCount := time.Now(), time.Now(), time.Now(), time.Now()
	// for each subset
	for mask := uint64(1); mask < uint64(1)<<uint(len(vecs)); mask++ {
		// build the subset
		tb := time.Now()
		//vsubset := make([]vec, 0, len(vecs)) // capacity can be optimized here: log(mask)
		var vmin vec
		first := true
		m := mask
		for i := range vecs {
			if m&1 == 1 {
				if first {
					first = false
					vmin = vecs[i]
				} else {
					vmin = combineVec(vmin, vecs[i])
				}
			}
			m >>= 1
		}
		t0 := time.Now()
		tBuild = tBuild.Add(t0.Sub(tb))
		//vmin := combineVec(vsubset)
		t1 := time.Now()
		tMin = tMin.Add(t1.Sub(t0))
		sum += countId(vmin)
		tCount = tCount.Add(time.Now().Sub(t1))
	}
	log.Println("Building subset: ", tBuild.Sub(tBase))
	log.Println("Combining vectors: ", tMin.Sub(tBase))
	log.Println("Counting connected components: ", tCount.Sub(tBase))
	return sum + 64 // +64 for the empty subset
}

// take minimum of all vectors combined
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
