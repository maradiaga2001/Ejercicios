package main

import (
	"fmt"
)

func main() {
	n := 5 // Cantidad de elementos 
	drawDiamond(n)
}

func drawDiamond(n int) {
	if n%2 == 0 || n <= 0 {
		fmt.Println("El número debe ser impar y positivo")
		return
	}

	// Dibuja la mitad superior del rombo
	for i := 1; i <= n/2+1; i++ {
		for j := 1; j <= n/2+1-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Dibuja la mitad inferior del rombo
	for i := n/2; i >= 1; i-- {
		for j := 1; j <= n/2+1-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}