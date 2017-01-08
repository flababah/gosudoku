gosudoku
========

Simple DFS based approach to brute-forcing sudoku puzzles. My first attempt at programming in Go.

The program groups each cell to a constraint mask shared with other cells. This makes finding a candidate value for a cell simply a matter of AND'ing the three constraint masks together. Each bit set in the result represents a valid value.

A linked list is used as a priority queue for picking the easiest of the remaining cells in each step. This does not necessarily make the program run any faster, but the assumption is that it's better to recurse on a cell which has two valid values rather than one that has three.

### Usage

The puzzle is given as an argument to the program by encoding it using the following scheme:

    .8.14.3..4.1..95...79638...9...1.4.3....8....1.7.2...5...89263...43..7.2..6.57.9.
    
Dots represent blank cells in the puzzle. The sequence formed by reading the puzzle left to right, top to bottom.

The result is output in the terminal:

    Solution found!
    6 8 2 1 4 5 3 7 9 
    4 3 1 2 7 9 5 8 6 
    5 7 9 6 3 8 2 4 1 
    9 5 8 7 1 6 4 2 3 
    2 6 3 5 8 4 9 1 7 
    1 4 7 9 2 3 8 6 5 
    7 1 5 8 9 2 6 3 4 
    8 9 4 3 6 1 7 5 2 
    3 2 6 4 5 7 1 9 8 

### Other examples
    ..57.9.....1.34.877....6.....247..1..8.....5..7..654.....6....352.39.6.....2.75..
    ..............3.85..1.2.......5.7.....4...1...9.......5......73..2.1........4...9