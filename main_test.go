package main

import (
	"testing"
)

func verifyEquality(moves1 []KnightMove, moves2 []KnightMove) bool {
	if len(moves1) != len(moves2) {
		return false
	} else {
		for idx, move := range moves1 {
			if move != moves2[idx] {
				return false
			}
		}
	}
	return true
}

var singleMoveBoard Board = Board{
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 0, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 0, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

var unbiasedMoveBoard Board = Board{
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 0, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 0, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 0, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

var biasedMoveBoard Board = Board{
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 0, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 0, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 0, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 0, 1, 0, 1, 1, 1},
}

func TestWarndorffSortByAccessibility(t *testing.T) {
	tests := []struct {
		moves                  [8]KnightMove
		currentBoard           Board
		currentPosition        Position
		expectedSuggestedMoves []KnightMove
	}{
		{KnightMoves, Board{}, Position{X: 0, Y: 0},
			[]KnightMove{KnightMoves[2], KnightMoves[3]}},
		{KnightMoves, singleMoveBoard, Position{X: 4, Y: 3}, []KnightMove{KnightMoves[7]}},
		{KnightMoves, unbiasedMoveBoard, Position{X: 4, Y: 3},
			[]KnightMove{KnightMoves[3], KnightMoves[7]}},
		{KnightMoves, biasedMoveBoard, Position{X: 4, Y: 3},
			[]KnightMove{KnightMoves[7], KnightMoves[4]}},
	}

	for _, test := range tests {
		suggestedMoves := warndorffSortByAccessibility(test.moves, test.currentBoard, test.currentPosition)
		if !verifyEquality(suggestedMoves, test.expectedSuggestedMoves) {
			t.Errorf("warndorffSortByAccessibility(%v, %v, %v) is %v, expected %v",
				test.moves, test.currentBoard, test.currentPosition, suggestedMoves, test.expectedSuggestedMoves)
		}
	}
}

func TestIsMoveValidOld(t *testing.T) {
	position := Position{X: 0, Y: 0}
	move := KnightMoves[2]
	var board Board
	isValid, newPos := isMoveValid(position, move, board)

	if !isValid {
		t.Errorf("Valid move was not accepted")
	}

	expectedPos := Position{X: 2, Y: 1}
	if newPos != expectedPos {
		t.Errorf("Move did not end at the expected square")
	}
}

func TestIsMoveValid(t *testing.T) {
	tests := []struct {
		position         Position
		board            Board
		move             KnightMove
		expectedValid    bool
		expectedPosition Position
	}{
		// left upper corner has two valid moves, verify that they are validated
		{Position{X: 0, Y: 0}, Board{}, KnightMoves[2], true, Position{X: 2, Y: 1}},
		{Position{X: 0, Y: 0}, Board{}, KnightMoves[3], true, Position{X: 1, Y: 2}},
		// verify that a move off the grid will not be validated
		{Position{X: 6, Y: 6}, Board{}, KnightMoves[2], false, Position{X: 8, Y: 7}},
	}

	for _, test := range tests {
		isValid, newPosition := isMoveValid(test.position, test.move, test.board)
		if isValid != test.expectedValid {
			t.Errorf("isMoveValid(%v, %v, %v) is %v, expected %v", test.position, test.move, test.board, isValid, test.expectedValid)
		}
		if newPosition != test.expectedPosition {
			t.Errorf("isMoveValid(%v, %v, %v) is %v, expected %v", test.position, test.move, test.board, newPosition, test.expectedPosition)
		}
	}
}
