package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresarNumeros() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Ingresa un numero: ")
		if scanner.Scan() {
			numero1, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Se encontro un error " + err.Error())
			} else {
				for i := 1; i <= 10; i++ {

					fmt.Println(numero1, " x ", i, " = ", i*numero1)

				}
				break
			}

		}

	}
}
