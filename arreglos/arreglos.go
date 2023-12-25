package arreglos

import (
	"fmt"
)

func Arreglos() {

	numeros := [10]int{1, 2, 3, 4, 5, 6, 2}

	fmt.Println(numeros)

	animales := make([]string, 0, 3)

	names := make(map[int]string)

	names[0] = "Jose"
	names[1] = "Abdiel"
	names[2] = "Adrian"

	for posicion, name := range names {
		fmt.Printf(" %d - %s \n", posicion, name)
	}

	personas := map[int]string{
		0: "Abdiel",
		1: "Jose",
		2: "Adrian",
		3: "Federico",
	}

	delete(personas, 0)

	for po, per := range personas {
		fmt.Printf("%d -> %s \n", po, per)
	}

	animales = append(animales, "Gato", "Caballo", "Colibri", "Mariposa")

	fmt.Printf("Tamaño: %d, Cap: %d \n", len(animales), cap(animales))

	num, existe := personas[1]

	fmt.Printf("El puntaje es %s, y %t", num, existe)

}
