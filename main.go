package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var n int
	position := [9]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	player := 1
	value := 'X'

	fmt.Println("Tic Tac Toe")
	fmt.Println("Player 1 (X) - Player 2 (O)")

	for {
		displayTable(position)

		fmt.Printf("Player %d, enter a number: ", player)
		fmt.Scanln(&input)
		n, _ = strconv.Atoi(input)
		position[n-1] = value

		if position[0] == 'X' && position[1] == 'X' && position[2] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[3] == 'X' && position[4] == 'X' && position[5] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[6] == 'X' && position[7] == 'X' && position[8] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[0] == 'X' && position[3] == 'X' && position[6] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[1] == 'X' && position[4] == 'X' && position[7] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[2] == 'X' && position[5] == 'X' && position[8] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[0] == 'X' && position[4] == 'X' && position[8] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[2] == 'X' && position[4] == 'X' && position[6] == 'X' {
			displayTable(position)
			fmt.Println("Player 1 win!!!!")
			return
		}

		if position[0] == 'O' && position[1] == 'O' && position[2] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if position[3] == 'O' && position[4] == 'O' && position[5] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if position[6] == 'O' && position[7] == 'O' && position[8] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if position[0] == 'O' && position[3] == 'O' && position[6] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if position[1] == 'O' && position[4] == 'O' && position[7] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if position[2] == 'O' && position[5] == 'O' && position[8] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if position[0] == 'O' && position[4] == 'O' && position[8] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if position[2] == 'O' && position[4] == 'O' && position[6] == 'O' {
			displayTable(position)
			fmt.Println("Player 2 win!!!!")
			return
		}

		if player == 1 {
			player = 2
			value = 'O'
		} else {
			player = 1
			value = 'X'
		}
	}
}

func displayTable(position [9]rune) {
	fmt.Println("   :   :   ")
	fmt.Printf(" %c : %c : %c \n", position[0], position[1], position[2])
	fmt.Println("___:___:___")
	fmt.Println("   :   :   ")
	fmt.Printf(" %c : %c : %c \n", position[3], position[4], position[5])
	fmt.Println("___:___:___")
	fmt.Println("   :   :   ")
	fmt.Printf(" %c : %c : %c \n", position[6], position[7], position[8])
	fmt.Println("   :   :   ")
}
