package ejercios

import (
	"fmt"
	"strconv"
)

func ConvertirStringANumero(valor string) (int, string) {

	nuevoValor, err := strconv.Atoi(valor)

	if err != nil {
		fmt.Println("Error durante la convecion: ", err)
		return nuevoValor, "Error"
	}

	if nuevoValor > 100 {
		return nuevoValor, "es mayor a 100"
	} else {
		return nuevoValor, "es menor a 100"
	}

}
