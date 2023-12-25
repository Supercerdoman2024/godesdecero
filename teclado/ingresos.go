package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var texto string

func IngresarNumeros() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Ingresa un numero: ")

		if scanner.Scan() {

			numero1, err := strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Se encontro un error " + err.Error())
			} else {
				for i := 1; i <= 10; i++ {
					op := i * numero1
					texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, op)

				}
				return texto
			}

		}

	}
}
