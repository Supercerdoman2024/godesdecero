package ejercios

import (
	"fmt"
	"strconv"
)

func ConvertirStringANumero(valor string) (int, string) {

	nuevoValor, err := strconv.Atoi(valor)

	if err != nil {
		fmt.Println("Error durante la convecion: ", err.Error())
		return nuevoValor, "Error"
	}

	switch xValor := 20; xValor {
	case 10:
		fmt.Println("Es menor a 10")
	default:
		fmt.Println("Es cual quiera ")
	}

	if nuevoValor > 100 {
		return nuevoValor, "es mayor a 100"
	} else {
		return nuevoValor, "es menor a 100"
	}

}

func Calculadora(operador string, valor1 int, valor2 int) int {

	switch operador {
	case "+":
		fmt.Println("Tipo de operacion: Suma")
		return valor1 + valor2
	case "-":
		if valor1 >= valor2 {
			fmt.Println("Tipo de operacion: Resta")
			return valor1 - valor2
		} else {
			return 0
		}
	case "*":
		{
			fmt.Println("Tipo de operacion:  Multiplicacion")
			return valor1 * valor2
		}
	default:
		return 0
	}

}
