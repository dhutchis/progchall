Subset Component Problem
========================

## Go Solution
### Union Find Approach
I used a sort of union find approach to storing each graph's connected nodes.  Let the node of minimum label number be the representative node for the partition of all nodes connected to that node.  Convert each int representing a graph into a vector of 8-bit unsigned integers representing using the union find structure.  Here are a few examples:

Graph#  | Vector 			|#ConnectedComponents
--------|-------------------|--------------------
0       | [0 1 2 3 ... 63]	| 64
5		| [0 1 0 3 ... 63]	| 63
9		| [0 1 2 0 ... 63]	| 63 
14		| [0 1 1 1 ... 63]	| 62

The number of connected components in the graph is equivalent to the number of entries whose value is equal to their index. To combine two graph vectors, just take the pairwise minimum of each entry in the vector.

### Outcome
Unfortunately, I ran out of time finishing my implementation.  The Go solution was also not fast enough, running out of time on the second half of the test sets.

The best optimization I can think of is vectorizing the calculations.  For example, the combining of vectors is a giant *map*.  Finding the number of connected components is a giant *reduce*.  How can we efficiently implement those in Go?  Haskell, Lisp or Matlab sound like better languages to use here.

A second optimization is to reuse the intermediary combined subset vectors.  

## Problem Description ##
<https://www.hackerrank.com/contests/101july13/challenges/subset-component>
>You are given an array with n 64-bit integers: d[0], d[1], …, d[n - 1].
>
>BIT(x, i) = (x » i) & 1. (where B(x,i) is the ith lower bit of x in binary form.)
>
>If we regard every bit as a vertex of a graph G, there exists one undirected edge between vertex i and vertex j if there exists at least one k such that BIT(d[k], i) == 1 && BIT(d[k], j) == 1.
>
>For every subset of the input array, how many connected-components are there in that graph?
>
>The number of connected-components in a graph are the sets of nodes, which are accessible to each other, but not to/from the nodes in any other set.
>
>For example if a graph has six nodes, labelled {1,2,3,4,5,6}. And contains the edges (1,2), (2,4) and (3,5). There are three connected-components: {1,2,4}, {3,5} and {6}. Because {1,2,4} can be accessed from each other through one or more edges, {3,5} can access each other and {6} is isolated from everone else.
>
>You only need to output the sum of the number of connected-component(S) in every graph.
>
>Input Format
>
> 	n
> 	d[0] d[1] ... d[n - 1]
>  
>Output Format
>
> 	S
>Constraint
>
> 	1 <= n <= 20
> 	0 <= d[i] <= 264 - 1
>
>Sample Input
>
> 	3
> 	2 5 9
>Sample Output
>
> 	504
>Explanation
>There are 8 subset of {2, 5, 9}.
>
>{} => We don’t have any number in this subset => no edge in the graph => Every node is a component by itself => Number of connected-components = 64.
>
>{2} => The Binary Representation of 2 is 00000010. There is a bit at only one position. => So there is no edge in the graph, ever node is a connected-component by itself => Number of connected-components = 64.
>
>{5} => The Binary Representation of 5 is 00000101. There is a bit at the 0th and 2nd position. => So there is an edge: (0, 2) in the graph => There is one component with a pair of nodes (0,2) in the graph. Apart from that, all remaining 62 vertices are indepenent components of one node each (1,3,4,5,6…63) => Number of connected-components = 63.
>
>{9}	 => The Binary Representation of 9 is 00001001. => There is a 1-bit at the 0th and 3rd position in this binary representation. => edge: (0, 3) in the graph => Number of components = 63
>
>{2, 5}	 => This will contain the edge (0, 2) in the graph which will form one component
>=> Other nodes are all independent components
>=> Number of connected-component = 63
>
>{2, 9}
>=> This has edge (0,3) in the graph
>=> Similar to examples above, this has 63 connected components
>
>{5, 9}	 => This has edges (0, 2) and (0, 3) in the graph
>=> Similar to examples above, this has 62 connected components
>
>{2, 5, 9}
>=> This has edges(0, 2) (0, 3) in the graph. All three vertices (0,2,3) make one component => Other 61 vertices are all independent components
>=> Number of connected-components = 62
>
>S = 64 + 64 + 63 + 63 + 63 + 63 + 62 + 62 = 504