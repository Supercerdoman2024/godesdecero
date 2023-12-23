package main

import (
	"fmt"

	"github.com/godesdecero/ejercios"
)

func main() {

	/*
			variables.MuestroEnteros()
			variables.RestoDeVariables()

			estado, texto := variables.ConvertirATexto(200)

			fmt.Println(estado)
			fmt.Println(texto)

			fmt.Println("Acediendo desde el main ... ", variables.Fecha)

		variables.SaberSistemaOperativo()
	*/

	valorConvertiro, texto := ejercios.ConvertirStringANumero("10")

	fmt.Println(valorConvertiro)
	fmt.Println(texto)

}
