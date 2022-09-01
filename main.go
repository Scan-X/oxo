package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/inancgumus/screen"
)

const dimension = 3

var gameBoard [dimension][dimension]Symbol
var player1, player2 string

type Symbol int

const (
	Space Symbol = iota
	X
	O
)

func init() {
	for i := range gameBoard {
		for j := range gameBoard[i] {
			gameBoard[i][j] = Space
		}
	}

}

// Display board status
func displayBoard() {
	color.Yellowln("    a  b  c ")
	for i := range gameBoard {
		color.Yellow.Printf("%d [", i)
		for j := range gameBoard[i] {
			switch gameBoard[i][j] {
			case 0:
				fmt.Print("   ")
			case 1:
				color.Redp(" X ")
			case 2:
				color.Greenp(" O ")
			}

		}
		color.Yellow.Printf("] %d \n", i)
	}
	color.Yellowln("    a  b  c ")
}

// input validation
func validateCoords(coords string) (int, int, error) {

	if (strings.Count(coords, "") - 1) != 2 {
		return 0, 0, errors.New("invalid coords")
	}

	x := int(coords[0]) - 97

	y, err := strconv.Atoi(string(coords[1]))
	if err != nil {
		return 0, 0, errors.New("invalid coords")
	}

	if (x < 0 || x > 2) || (y < 0 || y > 2) {
		return 0, 0, errors.New("invalid coords")
	}

	if gameBoard[y][x] != Space {
		return 0, 0, errors.New("already something here")
	}

	return x, y, nil
}

// return total of a given line
func countLines(board [dimension]Symbol) int {
	totLines := 0
	for i := range board {
		totLines += int(board[i])
	}

	return totLines
}

// Check if any space into a line
func isASpace(line [dimension]Symbol) bool {
	for i := range line {
		if line[i] == Space {
			return true
		}
	}

	return false
}

// Test if winning condition
func isWin(board [dimension][dimension]Symbol) (bool, string) {

	space := false
	for j := 0; j < dimension; j++ {

		craftLine := [dimension]Symbol{board[0][j], board[1][j], board[2][j]}

		if !isASpace(craftLine) {
			if countLines(craftLine) == 6 {
				return true, player1
			} else if countLines(craftLine) == 3 {
				return true, player2
			}
		} else {
			space = true
		}
	}

	for j := range board {
		if !isASpace(board[j]) {
			if countLines(board[j]) == 6 {
				return true, player1
			} else if countLines(board[j]) == 3 {
				return true, player2
			}
		} else {
			space = true
		}
	}

	diag1 := [dimension]Symbol{board[0][0], board[1][1], board[2][2]}
	diag2 := [dimension]Symbol{board[0][2], board[1][1], board[2][0]}

	if !isASpace(diag1) {
		if countLines(diag1) == 6 {
			return true, player1
		} else if countLines(diag1) == 3 {
			return true, player2
		}
	} else {
		space = true
	}

	if !isASpace(diag2) {
		if countLines(diag2) == 6 {
			return true, player1
		} else if countLines(diag2) == 3 {
			return true, player2
		}
	} else {
		space = true
	}

	if !space {
		return true, "No winner !"
	}
	return false, ""
}

func main() {
	var coords string

	firstPlayerTurn := true

	screen.Clear()

	color.Redp("Player1 (O) name: ")
	fmt.Scanf("%s\n", &player1)

	color.Greenp("Player2 (X) name: ")
	fmt.Scanf("%s\n", &player2)

	screen.Clear()
	displayBoard()

	for {

		if firstPlayerTurn {
			color.Red.Printf("%s turn ", player1)
			fmt.Scanf("%s\n", &coords)

			x, y, err := validateCoords(coords)

			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				gameBoard[y][x] = O
				firstPlayerTurn = false
			}

		} else {
			color.Green.Printf("%s turn ", player2)
			fmt.Scanf("%s\n", &coords)

			x, y, err := validateCoords(coords)

			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				gameBoard[y][x] = X
				firstPlayerTurn = true
			}
		}

		screen.Clear()
		displayBoard()
		win, player := isWin(gameBoard)
		if win {
			color.Yellowln("Game over, the winner is : ", player)
			os.Exit(0)
		}

	}
}

