package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Piece struct {
	Name  string
	Color string
}

type Square struct {
	Piece Piece
}

type Board [8][8]Square

func initializeBoard() Board {
	board := Board{}

	// Initialize white pieces
	board[0][0] = Square{Piece{"Rook", "White"}}
	board[0][1] = Square{Piece{"Knight", "White"}}
	board[0][2] = Square{Piece{"Bishop", "White"}}
	board[0][3] = Square{Piece{"Queen", "White"}}
	board[0][4] = Square{Piece{"King", "White"}}
	board[0][5] = Square{Piece{"Bishop", "White"}}
	board[0][6] = Square{Piece{"Knight", "White"}}
	board[0][7] = Square{Piece{"Rook", "White"}}
	for i := 0; i < 8; i++ {
		board[1][i] = Square{Piece{"Pawn", "White"}}
	}

	// Initialize black pieces
	board[7][0] = Square{Piece{"Rook", "Black"}}
	board[7][1] = Square{Piece{"Knight", "Black"}}
	board[7][2] = Square{Piece{"Bishop", "Black"}}
	board[7][3] = Square{Piece{"Queen", "Black"}}
	board[7][4] = Square{Piece{"King", "Black"}}
	board[7][5] = Square{Piece{"Bishop", "Black"}}
	board[7][6] = Square{Piece{"Knight", "Black"}}
	board[7][7] = Square{Piece{"Rook", "Black"}}
	for i := 0; i < 8; i++ {
		board[6][i] = Square{Piece{"Pawn", "Black"}}
	}

	// Initialize empty squares
	for i := 2; i < 6; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = Square{}
		}
	}

	return board
}

func displayBoard(board Board) {
	fmt.Println("  a b c d e f g h")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", 8-i)
		for j := 0; j < 8; j++ {
			if board[i][j].Piece.Name != "" {
				fmt.Printf("%s ", board[i][j].Piece.Name[:1])
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Printf("%d\n", 8-i)
	}
	fmt.Println("  a b c d e f g h")
}

func parseMove(move string) (int, int, int, int) {
	from := move[:2]
	to := move[2:]
	fromCol := int(from[0] - 'a')
	fromRow := 8 - int(from[1]-'0')
	toCol := int(to[0] - 'a')
	toRow := 8 - int(to[1]-'0')
	return fromRow, fromCol, toRow, toCol
}

func main() {
	board := initializeBoard()
	reader := bufio.NewReader(os.Stdin)

	for {
		displayBoard(board)
		fmt.Print("Enter move (e.g., e2e4): ")
		move, _ := reader.ReadString('\n')
		move = strings.TrimSpace(move)
		fromRow, fromCol, toRow, toCol := parseMove(move)
		board[toRow][toCol] = board[fromRow][fromCol]
		board[fromRow][fromCol] = Square{}
	}
}
