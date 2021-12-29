package main

import (
	"github.com/stephensli/advent-of-code-2021/helpers/aoc"
)

func partOne(sections [][]string) {
	aoc.PrintAnswer(1, "99298993199873")
}

func partTwo(sections [][]string) {
	aoc.PrintAnswer(2, "73181221197111")
}

// solved by hand.
//
// Z = stack
// Z div 1 = push
// Z div 26 = pop
// Want empty stack
//
// A-N (14 input numbers)
//
// z.push(A+6)
// z.push(B+2)
// z.push(C+13)
// if D!=z.pop()-6: z.push(D+8)
// z.push(E+13)
// if F!=z.pop()-12: z.push(F+8)
// z.push(G+3)
// z.push(H+11)
// z.push(I+10)
// if J!=z.pop()-2: z.push(J+8)
// if K!=z.pop()-5: z.push(K+14)
// if L!=z.pop()-4: z.push(L+6)
// if M!=z.pop()-4: z.push(L+8)
// if N!=z.pop()-12: z.push(N+2)
//
// equal push and pop, we want to always pop. The pushes are enforced since the if statements are
// always something along the lines of INPUT == VALUE_GREATER_THAN_9. Meaning its always a push.
//
// simplified version:
//
// C+7=D
// E+1=F
// I+8=J
// H=6=K
// G=1=L
// B-2=M
// A-6=N
//
// largest: 99298993199873
// smallest: 73181221197111
func main() {
	defer aoc.Setup(2021, 24)()

	input := parseInput("./input.txt")

	partOne(input)
	partTwo(input)
}
