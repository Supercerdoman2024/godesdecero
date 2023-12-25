package funciones

import "fmt"

var funOp = func() {
	println("Funciona fantasmas")
}

func tabla(valor int) func() int {

	numero := valor
	secuencia := 0

	fmt.Println("\n Me ejecuto una vez \n")

	return func() int {
		secuencia++
		funOp()
		fmt.Println("\n Me ejecuto varias veces \n")

		return numero * secuencia
	}

}

func LlamarClosure() {
	tablaDel := 3
	Mitabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", i, tablaDel, Mitabla())
	}
}
