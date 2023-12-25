package files

import (
	"fmt"
	"os"

	"github.com/godesdecero/teclado"
)

var filename string = "./files/txt/tabla.txt"

func GrabarTabla() {

	texto := teclado.IngresarNumeros()

	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("error al crear el archivo " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = teclado.IngresarNumeros()
	if Append(filename, texto) == false {
		fmt.Println("Error a lconcatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append", err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("error duranete el Append")
		return false
	}

	arch.Close()
	return true

}
