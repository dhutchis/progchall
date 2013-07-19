Meeting Point
=============

## Solution ##
Implemented in [Go](http://golang.org/).  Code submission on [Hackerrank](https://www.hackerrank.com/submissions/code/732308).

1. **bruteForceParallel.go** contains a somewhat naive solution.  Calculate the meeting point for each point in parallel.  Stop the calculation early if the sum exceeds the best sum so far
2. **solution.go** uses a mathematical trick: just find the center of all the points (average the x,y point coordinates) and take the point closest to the center as the best meeting point.  I'm not sure how to mathematically justify this, but it makes sense intuitively.
3. **solutionParallel.go** parallelizes the "find closest point to the center" and "calculate the Manhattan distance from the meeting point to all other points" in solution.go. 
4. **main.go** is the same as solutionParallel.go, with timing.


### Lessons Learned ###
Profile and time your code before optimizing!  I made the mistake of going straight for parallelization when all the CPU time is taken up in the "data read-in" phase.  These stats tell all the difference:

	read-in 22.4452838s
	closest 2.0001ms
	man dist 1.0001ms

After seeing the stats, I swapped out `fmt.Scan` for `bufio.NewScanner` and `strconv.Atoi` and came up with this incredible difference:

	read-in 202.0115ms
	closest 2.0001ms
	man dist 2.0001ms

Much better!  Lesson: never use `fmt.Scan`...

As some form of proof that the parallelization did accomplish something, here are the timing stats for the unparallelized version.  All these stats were taken with the 100,000 point data set in *test10.in*.

	read-in 199.0114ms
	closest 6.0004ms
	man dist 5.0003ms

## Problem ##
<https://www.hackerrank.com/challenges/meeting-point>
> There is an infinite integer grid where N people live in N different houses. They decide to create a meeting point at one person’s house. 
> 
> From any given cell, all 8 adjacent cells are reachable in 1 unit of time, e.g. (x,y) can be reached from (x-1,y+1) in one unit of time. Find a common meeting place which minimizes the combined travel time of everyone.
> 
> #####Input Format 
> N for N houses.
> 
> The following N lines will contain two integers for the x and y coordinate of the nth house.
> 
> #####Output Format 
> M for the minimum sum of all travel times when everyone meets at the best location.
> 
> #####Constraints
> N <= 105
> The absolute value of each co-ordinate in the input will be at most 109
> 
> **HINT:** Please use long long 64-bit integers;
> 
> #####Input #1
> 
> 	4 
> 	0 1
> 	2 5 
> 	3 1 
> 	4 0 
> #####Output #1
> 
> 	8
> #####Explanation
> The houses will have a travel-sum of 11, 13, 8 or 10. 8 is the minimum.
> 
> #####Input #2
> 
> 	6 
> 	12 -14 
> 	-3 3 
> 	-14 7 
> 	-14 -3 
> 	2 -12 
> 	-1 -6
> #####Output #2:
> 
> 	54

