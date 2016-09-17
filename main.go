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
	count := 0

	fmt.Println("Tic Tac Toe")
	fmt.Println("Player 1 (X) - Player 2 (O)")

	for {
		displayTable(position)

		fmt.Printf("Player %d, enter a number: ", player)
		fmt.Scanln(&input)
		n, _ = strconv.Atoi(input)

		if position[n-1] == 'X' || position[n-1] == 'O' {
			fmt.Println("That is not a valid move! Pease try again.")
			continue
		}

		position[n-1] = value
		count++

		if position[0] == value && position[1] == value && position[2] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if position[3] == value && position[4] == value && position[5] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if position[6] == value && position[7] == value && position[8] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if position[0] == value && position[3] == value && position[6] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if position[1] == value && position[4] == value && position[7] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if position[2] == value && position[5] == value && position[8] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if position[0] == value && position[4] == value && position[8] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if position[2] == value && position[4] == value && position[6] == value {
			displayTable(position)
			fmt.Printf("Player %d is a winner!!!\n", player)
			return
		}

		if count == 9 {
			fmt.Println("It's a draw.")
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
