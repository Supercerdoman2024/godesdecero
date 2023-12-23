package variables

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoDeVariables() {

	Nombre = "Abdiel"
	Estado = false
	Sueldo = 20.500
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

func ConvertirATexto(numero int) (bool, string) {

	texto := strconv.Itoa(numero)

	return true, texto

}

func SaberSistemaOperativo() {

	sistema := runtime.GOOS

	if sistema == "darwin" {
		fmt.Println("Si es macos")
	} else {
		fmt.Println("No es macos")
	}

}
