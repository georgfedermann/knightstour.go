package main

import (
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

func printBoards(board1 Board, board2 Board) {
	for row := 0; row < len(board1); row++ {
		for col := 0; col < len(board1[0]); col++ {
			fmt.Printf("%d ", board1[row][col])
		}
		fmt.Print("   ")
		for col := 0; col < len(board2[0]); col++ {
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
			accessibilityCount := heuristics[newPos.X][newPos.Y]
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

func findPath(currentBoard Board, currentPosition Position, moveNumber int) []Position {
    currentBoard[currentPosition.Y][currentPosition.X] = moveNumber
    heuristics := calculateHeuristics(currentBoard)
    fmt.Printf("Move %d\n", moveNumber)
    printBoards(currentBoard, heuristics)

	if moveNumber == 64 {
		return []Position{currentPosition}
	} else {
		for _, move := range warndorffSortByAccessibility(KnightMoves, currentBoard, currentPosition) {
            subPath := findPath(currentBoard, applyMove(currentPosition, move), moveNumber + 1)
            if len(subPath) > 0 {
                return append([]Position{currentPosition}, subPath...)
            }
		}
	}
	return []Position{}
}

func main() {
	var board Board
	heuristics := calculateHeuristics(board)

	printBoards(board, heuristics)
	fmt.Println(KnightMoves)
	fmt.Println(len(KnightMoves))
	fmt.Println(warndorffSortByAccessibility(KnightMoves, board, Position{X: 1, Y: 2}))
    fmt.Println()
    findPath(board, Position{X: 0, Y: 0}, 1)
}
