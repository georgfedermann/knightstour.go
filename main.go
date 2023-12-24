package main

import (
	"flag"
	"fmt"
)

type Board [8][8]int

type KnightMove struct {
	Dx, Dy int
}

type Position struct {
	X, Y int
}

var KnightMoves [8]KnightMove = [8]KnightMove{
	{1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2},
}

var verbose bool = false

func printBoards(board1 Board, board2 Board) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			fmt.Printf("%02d ", board1[row][col])
		}
		fmt.Print("   ")
		for col := 0; col < 8; col++ {
			fmt.Printf("%d ", board2[row][col])
		}
		fmt.Println()
	}
}

func isMoveValid(position Position, move KnightMove, board Board) (bool, Position) {
	pos := Position{X: position.X + move.Dx, Y: position.Y + move.Dy}
	return 0 <= pos.X && pos.X < 8 &&
		0 <= pos.Y && pos.Y < 8 &&
		board[pos.Y][pos.X] == 0, pos
}

func calculateHeuristics(board Board) Board {
	var heuristics [8][8]int
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			var pos Position = Position{X: col, Y: row}
			for _, move := range KnightMoves {
				isValid, _ := isMoveValid(pos, move, board)
				if isValid {
					heuristics[row][col]++
				}
			}
		}
	}
	return heuristics
}

func warndorffSortByAccessibility(moves [8]KnightMove, currentBoard Board, pos Position) []KnightMove {
	heuristics := calculateHeuristics(currentBoard)
	var results [9][]KnightMove = [9][]KnightMove{}
	for _, move := range moves {
		isValid, newPos := isMoveValid(pos, move, currentBoard)
		if isValid {
			accessibilityCount := heuristics[newPos.Y][newPos.X]
			results[accessibilityCount] = append(results[accessibilityCount], move)
		}
	}
	var priorityList []KnightMove

	for _, slice := range results {
		priorityList = append(priorityList, slice...)
	}

	return priorityList
}

func applyMove(position Position, move KnightMove) Position {
	return Position{X: position.X + move.Dx, Y: position.Y + move.Dy}
}

// FindPath attempts to find a path for a knight on a chessboard that visits
// every sqaure exactly once using Warndorff's rule in a backtracking fashion.
// It is a recursive backtracking algorithm where each call attempts to find
// the next move.
// currentBoard is the current state of the board with the positions already
// visited marked with the move number. This way, the algorithm can also try
// to solve a position that has already been partially solved.
// currentPosition is the current position of the knight on the board.
// moveNumber is the current move number. It is used to determine if the
// board is solved.
// The function returns a slice of Position that represents the path of the knight.
// If not path is found, an empty slice is returned. At this point, the calling
// function should backtrack.
func FindPath(currentBoard Board, currentPosition Position, moveNumber int) []Position {
	// Mark the current position on the board with the move number
	currentBoard[currentPosition.Y][currentPosition.X] = moveNumber
	// Calculate the accessibility heuristics for the current board
	heuristics := calculateHeuristics(currentBoard)
	// Deebugging output showing the current move and the state of the board
	// and the heuristics
	if verbose {
		fmt.Printf("Move %d -> %v\n", moveNumber, currentPosition)
		printBoards(currentBoard, heuristics)
	}
	// If the movenumber is 64, we have visited every square and hence found a full path
	if moveNumber == 64 {
		// Add the currentPosition to return the full path
		return []Position{currentPosition}
	} else {
		// If we have not finished yet, iterate over all possible moves prioritized
		// by the accessibility heuristics
		for _, move := range warndorffSortByAccessibility(KnightMoves, currentBoard, currentPosition) {
			// Recursively call FindPath with the new position and the next move number
			subPath := FindPath(currentBoard, applyMove(currentPosition, move), moveNumber+1)
			// If a subpath is found, append the current position and return the path
			if len(subPath) > 0 {
				return append([]Position{currentPosition}, subPath...)
			}
		}
	}
	// If no move leads to a solution, return an empty path indicating failure to
	// find a complete path. The calling function should now backtrack.
	return []Position{}
}

func main() {
	xPtr := flag.Int("x", 0, "The X position (column) to start the knight's tour at")
	yPtr := flag.Int("y", 0, "The Y position (row) to start the knight's tour at")
	verbosePtr := flag.Bool("v", false, "Enable verbose output")
    flag.Parse()
	verbose = *verbosePtr

	var board Board
	path := FindPath(board, Position{X: *xPtr, Y: *yPtr}, 1)
	fmt.Println(path)
}
