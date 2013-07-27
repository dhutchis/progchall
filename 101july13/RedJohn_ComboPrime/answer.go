package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// pre-computed answers for all 1 <= 40 <= N
var Parr = [40]int{0, 0, 0, 1, 2, 2, 3, 4, 4, 6, 8, 9, 11, 15, 19, 24, 32, 42, 53, 68, 91, 119, 155, 204, 269, 354, 462, 615, 816, 1077, 1432, 1912, 2543, 3385, 4522, 6048, 8078, 10794, 14475, 19385}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	//T, _ := strconv.Atoi(scan.Text())
	for scan.Scan() {
		N, _ := strconv.Atoi(scan.Text())
		fmt.Println(Parr[N-1])
	}

}
