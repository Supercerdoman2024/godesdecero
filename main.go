package main

import (
	"fmt"

	"github.com/godesdecero/variables"
)

func main() {

	variables.MuestroEnteros()
	variables.RestoDeVariables()

	estado, texto := variables.ConvertirATexto(200)

	fmt.Println(estado)
	fmt.Println(texto)

	fmt.Println("Acediendo desde el main ... ", variables.Fecha)

}
