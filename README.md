A simple maze solver written in Go

This program will read a maze from an input file
and solve it using a follow the left wall method. 

The first line of the maze description is
ROWS COLUMNS START_ROW START_COLUMN

The rest of the file is the maze layout. A
hash mark (#) is a wall. Any other charecter is
an open floor. 

Example maze file:

6 7 5 5

####.##

####.##

#....##

##.####

##....#

#####.#

Example execution:

$ go run maze.go -file text.txt

Maze line: ####.##

Maze line: ####.##

Maze line: #....##

Maze line: ##.####

Maze line: ##....#

Maze line: #####.#

Maze size: 6 by 7

Start point: (5,5)

[[1 1 1 1 0 1 1] [1 1 1 1 0 1 1] [1 0 0 0 0 1 1] [1 1 0 1 1 1 1] [1 1 0 0 0 0 1] [1 1 1 1 1 0 1]]

Go Up

Go Left

Go Left

Go Left

Go Up

Go Up

Go Left

Go Right

Go Right

Go Right

Go Up

Go Up

Exit found at (0,4)



Limitations:

- The start location is hardcoded as being from the bottom. 

- The first exit found (other than the entrance) is used.

Unknown Limitations:

- Unknown
