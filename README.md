## Why are we poor at applying algorithms?

We humans are poor at memorizing, and tracing state transition across the code. Many messy codebases do not declare a right data structure, but advances the logic in hurry. This drops readability, bothering readers who understand the code. Here are Rob Pike's 5 rules:

Rule 1. You can't tell where a program is going to spend its time. Bottlenecks occur in surprising places, so don't try to second guess and put in a speed hack until you've proven that's where the bottleneck is.

Rule 2. Measure. Don't tune for speed until you've measured, and even then don't unless one part of the code overwhelms the rest.

Rule 3. Fancy algorithms are slow when n is small, and n is usually small. Fancy algorithms have big constants. Until you know that n is frequently going to be big, don't get fancy. (Even if n does get big, use Rule 2 first.)

Rule 4. Fancy algorithms are buggier than simple ones, and they're much harder to implement. Use simple algorithms as well as simple data structures.

Rule 5. Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming.

The preceding rules highlight a good algorithm (or code) can be mostly achieved by right data structures.

## State Is Hidden in Control Flow

Algorithm examples mostly appears only with the correct forms. If we have to implement BubbleSort, we mostly write code as:
```go
package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)
}
```

When i saw this code first, I didn't understand what i and j really are, and all the textbooks can't make me understand clearly. I still don't want to memorize `n-i-1`. Memorizing without understanding is like a solving puzzle only with an exact instruction guide.

I made a negative conclusion. It was hard to understand because the state lives in the code. we should memorize what each i and j transforms to which value, and replay entangled relationships and conditions by following code.

In other words, state transition and its semantics are encoded in for loop implictly. I really hated this.

## State Machine

So I began to re-model this algorithm. There were 2 important things to consider:

- Put the state in a data structure explicitly.
- Let data structures represent the state explicitly, and let small functions express how that state evolves.

I headed back to the state machine, which is a fundamental element for computer science and OS. 


 