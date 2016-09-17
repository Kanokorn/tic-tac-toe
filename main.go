package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var n int
	position := [9]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	fmt.Println("Tic Tac Toe")
	fmt.Println("Player 1 (X) - Player 2 (O)")

	fmt.Println("   :   :   ")
	fmt.Printf(" %c : %c : %c \n", position[0], position[1], position[2])
	fmt.Println("___:___:___")
	fmt.Println("   :   :   ")
	fmt.Printf(" %c : %c : %c \n", position[3], position[4], position[5])
	fmt.Println("___:___:___")
	fmt.Println("   :   :   ")
	fmt.Printf(" %c : %c : %c \n", position[6], position[7], position[8])
	fmt.Println("   :   :   ")

	fmt.Print("Player 1, enter a number: ")
	fmt.Scanln(&input)
	n, _ = strconv.Atoi(input)
	position[n-1] = 'X'

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
