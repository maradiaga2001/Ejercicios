package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Print("Indique el nombre del archivo de texto: ")
	var filename string
	fmt.Scanln(&filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	numCaracteres := len(data) // incluye los saltos de línea
	numPalabras := len(strings.Fields(string(data)))
	numLineas := strings.Count(string(data), "\n")

	fmt.Printf("Cantidad de caracteres: %d\nCantidad de palabras: %d\nNúmero de líneas: %d\n", numCaracteres, numPalabras, numLineas)
}
