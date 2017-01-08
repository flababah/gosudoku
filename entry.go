package main

/*
 * Simple GO program to solve sudokus.
 */

import (
	"os"
	"fmt"
)

// TODO gør først på færrest sansynlige.
// TODO sorter contant væk fra working set.

// TODO hvordan er order of type const var etc... ?
// TODO kør gofmt

type mask uint16 // Enough to hold 9 bits.

type cell struct {
	groups [3]*mask
	value uint
	computed bool
}

const (
	DEFAULT_MASK mask = (1 << 9) - 1
)

var (
	// Group resolvers.
	resolvers = [3]func(int)int{

		// Row.
		func(i int) int { return i / 9 },

		// Column.
		func(i int) int { return i % 9 },

		// Box.
		func(i int) int {
			return (i / 27) * 3 +         // Row
				(i / 3) - (i / 9) * 3 // Column
		},
	}
	constraints = [3][9]mask{ }
	cells = [9*9]cell{ }
)

// Link the cell to its appropriate row, column and box constraints.
func (c *cell) initConstraints(index int) {
	for group, resolver := range resolvers {
		c.groups[group] = &constraints[group][resolver(index)]
	}
}

// Calculate possible candidates. Returns a closure that iterates
// the candidates.
func (c *cell) possibilities() mask {
	return *c.groups[0] & *c.groups[1] & *c.groups[2]
}

// Remove candidate bit from all constraint groups referenced by c.
func (c *cell) unset(candidate uint) {
	for _, group := range c.groups {
		*group &= ^(1 << candidate)
	}
}

// Reenter candidate for cell.
func (c *cell) set(candidate uint) {
	for _, group := range c.groups {
		*group |= 1 << candidate
	}
}

func exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format + "\n", args)
	os.Exit(1)
}

func printSolution() {
	for i, cell := range cells {
		if i % 9 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%c ", cell.value + '1')
	}
}

func solve(index, depth uint) bool {
	if index == 9 * 9 {
		return true
	}
	cell := &cells[index]

	if cell.computed {
		return solve(index + 1, depth + 1)
	}
	candidates := cell.possibilities()
	for value := uint(0); value != 9; value++ {
		if candidates & (1 << value) != 0 {
			cell.unset(value)
			if solve(index + 1, depth + 1) {
				cell.value = value // Setting here is OK
				return true
			}
			cell.set(value)
		}
	}
	return false
}

func main() {
	// Set default bits in constraints.
	for i, group := range constraints {
		for j, _ := range group {
			constraints[i][j] = DEFAULT_MASK
		}
	}

	// Parse input and fill out cells.
	if len(os.Args) < 2 {
		exit( "Usage: %s input", os.Args[0])
	}
	input := os.Args[1]
	if len(input) != 9 * 9 {
		exit("Puzzle is not 9 times 9")
	}
	for idx, char := range input {
		cell := &cells[idx]
		cell.initConstraints(idx)

		switch char {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value := uint(char) - '0' - 1
			cell.value = value
			cell.computed = true
			cell.unset(value)
			continue
		case '.':
			continue
		}
		exit("Invalid charater in input: %c", char)
	}
	if !solve(0, 0) {
		exit("Impossible puzzle")
	}
	fmt.Println("Solution found!")
	printSolution()
}