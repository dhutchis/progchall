Red John is Back
================
We are given **N**.  The problem can be decomposed into two subproblems:

1: Find **M**, the number of possible configurations of 4x1 and 1x4 blocks to fill a 4xN wall (note that C from the problem description is 4).  Writing out a few examples, the answer is (N choose 0) for no horizontal blocks + (N-3 choose 1) for 1 horizontal block + (N-6 choose 2) + ... up to floor(N/4).  Here is Mathematica code to find H for all N within the problem limits, 1 <= N <= 40:

```Mathematica
Do[Print[ Sum[Binomial[n - 3*H, H], {H, 0, Floor[n/4]}]], {n, 1, 40}]
```

2: Find **P**, the number of primes <= **H**.  This is known as the [prime-counting function](https://en.wikipedia.org/wiki/Prime-counting_function).  Because H is small (largest possible H is 217286), we can just compute all the primes up to 217286 and, for each possible H, count how many primes are <= H.  I wrote *findP.go* to do this via the Sieve of Eratosthenes, with some extra prime-counting code at the end.

The pre-computed numbers for both problems are in *data.csv*.  Creating the final program just needs a lookup table.  Success!


## Problem Description ##
<https://www.hackerrank.com/contests/101july13/challenges/red-john-is-back>
>Red John has committed another murder. But this time, he doesn’t leave a red smiley behind. What he leaves behind is a puzzle for Patrick Jane to solve. He also texts Teresa Lisbon that if Patrick is successful, he will turn himself in. The puzzle begins as follows.
>
>There is a wall of size CxN in the victim’s house where C is the 1st composite number. The victim also has an infinite supply of bricks of size Cx1 and 1xC in her house. There is a hidden safe which can only be opened by a particular configuration of bricks on the wall. In every configuration, the wall has to be completely covered using the bricks. There is a phone number written on a note in the safe which is of utmost importance in the murder case. Gale Bertram wants to know the total number of ways in which the bricks can be arranged on the wall so that a new configuration arises every time. He calls it M. Since Red John is back after a long time, he has also gained a masters degree in Mathematics from a reputed university. So, he wants Patrick to calculate the number of prime numbers (say P) up to M (i.e. <= M). If Patrick calculates P, Teresa should call Red John on the phone number from the safe and he will surrender if Patrick tells him the correct answer. Otherwise, Teresa will get another murder call after a week.
>
>You are required to help Patrick correctly solve the puzzle.
>
>Sample Input
>The first line of input will contain an integer T followed by T lines each containing an integer N.
>
>Sample Output
>Print exactly one line of output for each test case. The output should contain the number P.
>
>Constraints
> 	1<=T<=20
> 	1<=N<=40
>
>Sample Input
>
> 	2
> 	1
> 	7
>Sample Output
>
> 	0
> 	3
>Explanation
>
>For N = 1, the brick can be laid in 1 format only
>
><img src="http://hr-filepicker.s3.amazonaws.com/brick1.jpg" width="50" />
>
>The number of primes <= 1 is 0 and hence the answer.
>
>For N = 7, one of the ways in which we can lay the bricks is
>
><img src="http://hr-filepicker.s3.amazonaws.com/brick2.jpg" width="300" />
>
>There are 5 ways of arranging the bricks for N = 7 and there are 3 primes <= 5 and hence the answer 3.



-