package main

import "fmt"

func rotateArray(arr *[]int, n int, dir int) {
	length := len(*arr)
	if n >= length {
		n = n % length
	}
	if n < 0 {
		n = length + n
	}

	if dir == 0 {
		// Rotar a la izquierda
		*arr = append((*arr)[n:], (*arr)[:n]...)
	} else if dir == 1 {
		// Rotar a la derecha
		*arr = append((*arr)[length-n:], (*arr)[:length-n]...)
	}
}

func main() {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Secuencia original:", original)

	rotaciones := []struct {
		cantidad  int
		direccion int
	}{
		{3, 0},
		{2, 1},
		{1, 0},
	}

	for _, r := range rotaciones {
		rotateArray(&original, r.cantidad, r.direccion)
		fmt.Printf("Secuencia rotada %d veces a la %s: %v\n", r.cantidad, map[int]string{0: "izquierda", 1: "derecha"}[r.direccion], original)
	}
}
