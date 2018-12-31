package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

var Maze struct {
	columns int
	rows int
	startRow int
	startColumn int
	maze [][]int
}

type Direction int

const (
	Up Direction    = 0
	Left Direction  = 1
	Right Direction = 2
	Down Direction  = 3
)

func usage() {
	fmt.Println( "Script Usage:" )
	fmt.Printf( "$ %s -file FILENAME\n", path.Base(os.Args[0]))
	fmt.Println( "Where FILENAME is the name of the file to parse.")
}

func exitCheck( curRow int, curCol int) (finished bool) {
	//fmt.Printf("Testing for exit at (%d, %d)\n", curRow, curCol)
	//fmt.Printf("Maze.rows: %d Maze.columns:%d\n", Maze.rows-1, Maze.columns-1)
	if curRow == Maze.startRow && curCol == Maze.startColumn {
		finished = false
	} else if curRow == 0 || curCol == 0 || curRow == Maze.rows-1 || curCol == Maze.columns - 1 {
		finished = true
	} else {
		finished = false
	}
	return
}

func solveMaze( curRow int, curCol int, forward Direction) {
	//fmt.Printf("Current spot (%d,%d)\n", curRow, curCol)

	if exitCheck( curRow, curCol ) {
		fmt.Printf("Exit found at (%d,%d)\n", curRow, curCol )
		return
	}

	if forward == Up {
		//look left
		if Maze.maze[curRow][curCol-1] == 0 {
			fmt.Println( "Go Left" )
			solveMaze(curRow, curCol-1, Left)
		//look up
		} else if Maze.maze[curRow-1][curCol] == 0 {
			fmt.Println( "Go Up" )
			solveMaze(curRow-1, curCol, Up)
		//look right
		} else if Maze.maze[curRow][curCol+1] == 0 {
			fmt.Println( "Go Right" )
			solveMaze(curRow, curCol+1, Right)
		} else if Maze.maze[curRow+1][curCol] == 0 {
		//look down
			fmt.Println( "Go Down" )
			solveMaze(curRow+1, curCol, Down)
		}
	} else if forward == Left {
		if Maze.maze[curRow+1][curCol] == 0 {
		//look down
			fmt.Println( "Go Down" )
			solveMaze(curRow+1, curCol, Down)
		//look left
		} else if Maze.maze[curRow][curCol-1] == 0 {
			fmt.Println( "Go Left" )
			solveMaze(curRow, curCol-1, Left)
		//look up
		} else if Maze.maze[curRow-1][curCol] == 0 {
			fmt.Println( "Go Up" )
			solveMaze(curRow-1, curCol, Up)
		//look right
		} else if Maze.maze[curRow][curCol+1] == 0 {
			fmt.Println( "Go Right" )
			solveMaze(curRow, curCol+1, Right)
		}
	} else if forward == Right {
		//look up
		if Maze.maze[curRow-1][curCol] == 0 {
			fmt.Println( "Go Up" )
			solveMaze(curRow-1, curCol, Up)
		//look right
		} else if Maze.maze[curRow][curCol+1] == 0 {
			fmt.Println( "Go Right" )
			solveMaze(curRow, curCol+1, Right)
		} else if Maze.maze[curRow+1][curCol] == 0 {
		//look down
			fmt.Println( "Go Down" )
			solveMaze(curRow+1, curCol, Down)
		} else if Maze.maze[curRow][curCol-1] == 0 {
			// look left
			fmt.Println( "Go Left" )
			solveMaze(curRow, curCol-1, Left)
		}
	} else if forward == Down {
		//look right
		if Maze.maze[curRow][curCol+1] == 0 {
			fmt.Println( "Go Right" )
			solveMaze(curRow, curCol+1, Right)
		} else if Maze.maze[curRow+1][curCol] == 0 {
		//look down
			fmt.Println( "Go Down" )
			solveMaze(curRow+1, curCol, Down)
		} else if Maze.maze[curRow][curCol-1] == 0 {
			//look left
			fmt.Println( "Go Left" )
			solveMaze(curRow, curCol-1, Left)
		//look up
		} else if Maze.maze[curRow-1][curCol] == 0 {
			fmt.Println( "Go Up" )
			solveMaze(curRow-1, curCol, Up)
		}
	}
}

func main(){ //{{{
	filePtr := flag.String("file","","a file name")
	flag.Parse()

	if *filePtr == "" {
		usage()
		//log.Fatal("No file name provided")
		return
	}

	//fmt.Printf("File argument: %s\n", *filePtr)

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner( file )
	lineNumber := 0
	for scanner.Scan() {
		//fmt.Printf( "(%d): %s\n", lineNumber, scanner.Text() )
		if lineNumber == 0 {
			parts := strings.Split( scanner.Text() , " ")
			//fmt.Printf("First part of line 1: %s\n", parts[0] )
			//fmt.Printf("Second part of line 1: %s\n", parts[1] )
			Maze.rows,err = strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			Maze.columns,err = strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			Maze.startRow,err = strconv.Atoi(parts[2])
			if err != nil {
				log.Fatal(err)
			}
			Maze.startColumn,err = strconv.Atoi(parts[3])
			if err != nil {
				log.Fatal(err)
			}
			Maze.maze = make([][]int, Maze.rows)
			for i:=0; i<Maze.rows; i++ {
				Maze.maze[i] = make( []int, Maze.columns)
			}
		} else {
			fmt.Printf("Maze line: %s\n", scanner.Text() )
			parts := strings.Split( scanner.Text() , "")
			for index, character := range parts {
				//fmt.Printf("Found %s at %d\n" , character, index )
				if character == "#" {
					//fmt.Printf("(%d,%d) = 1 ", lineNumber, index)
					Maze.maze[lineNumber-1][index] = 1
				} else {
					//fmt.Printf("(%d,%d) = o ", lineNumber, index)
					Maze.maze[lineNumber-1][index] = 0
				}
			}
			//fmt.Println()
		}
		lineNumber++
	}
	fmt.Printf("Maze size: %d by %d\n", Maze.rows, Maze.columns)
	fmt.Printf("Start point: (%d,%d)\n", Maze.startRow, Maze.startColumn)
	fmt.Printf("%d\n",Maze.maze)

	solveMaze( Maze.startRow, Maze.startColumn, Up )
} //}}}

