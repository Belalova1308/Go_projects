
<h3 id="ex00">Exercise 00: Toys on a Tree</h3>

After some time, you two put together a structure for a [Binary tree](https://en.wikipedia.org/wiki/Binary_tree) node:

```go
type TreeNode struct {
    HasToy bool
    Left *TreeNode
    Right *TreeNode
}
```
&mdash; Looks like you are supposed to... "hang toys" on trees? - Lily looked a bit confused. - Okay, anyway, let's hope just a boolean value will suffice. But they say it's also wrong to put more toys on one side, should it be uniform?

&mdash; Okay, I get it - you said. - Let's write a function `areToysBalanced` which will receive a pointer to a tree root as an argument. The point is to spit out a true/false boolean value depending on if left subtree has the same amount of toys as the right one. The value on the root itself can be ignored.

So, your function should return `true` for such trees (0/1 represent false/true, equal amount of 1's on both subtrees):

```
    0
   / \
  0   1
 / \
0   1
```

```
    1
   /  \
  1     0
 / \   / \
1   0 1   1
```

and `false` for such trees (non-equal amount of 1's on both subtrees):

```
  1
 / \
1   0
```

```
  0
 / \
1   0
 \   \
  1   1
```

<h3 id="ex01">Exercise 01: Decorating</h3>

So, now you have to write another function called `unrollGarland()`, which also receives a pointer to a root node. The idea is to go top down, layer by layer, going right on even horisontal layers and going left on every odd. The returned value of this function should be a slice of bools. So, for this tree:

```
    1
   /  \
  1     0
 / \   / \
1   0 1   1
```

The answer will be [true, true, false, true, true, false, true] (root is true, then on second level we go from left to right, and then on third from right to left, like a zig-zag).

<h3 id="ex02">Exercise 02: Heap of Presents</h3>

```go
type Present struct {
    Value int
    Size int
}
```
You need to implement a PresentHeap data structure (using built-in library "container/heap" is recommended, but is not strictly required). Presents are compared by Value first (most valuable present goes on top of the heap). *Only* in case two presents have an equal Value, the smaller one is considered to be "cooler" than the other one (wins in comparison).

Apart from the structure itself, you should implement a function `getNCoolestPresents()`, that, given an unsorted slice of Presents and an integer `n`, will return a sorted slice (desc) of the "coolest" ones from the list. It should use the PresentHeap data structure inside and return an error if `n` is larger than the size of the slice or is negative.

So, if we represent each Present by a tuple of two numbers (Value, Size), then for this input:

```
(5, 1)
(4, 5)
(3, 1)
(5, 2)
```

the two "coolest" Presents would be [(5, 1), (5, 2)], because the first one has the smaller size of those two with Value = 5.

<h3 id="ex03">Exercise 03: Knapsack</h3>

As a last task, you have to implement a classic dynamic programming algorithm, also known as "Knapsack Problem". Input is almost the same, as in the last task - you have a slice of Presents, each with Value and Size, but this time you also have a hard drive with a limited capacity. So, you have to pick only those presents, that fit into that capacity and maximize the resulting value.

Please write a function `grabPresents()`, that receives a slice of Present instances and a capacity of your hard drive. As an output, this function should give out another slice of Presents, which should have a maximum cumulative Value that you can get with such capacity.


