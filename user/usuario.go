package user

import (
	"fmt"
	"time"

	"github.com/godesdecero/modelos"
)

func RegistrarUsuario() {

	user := new(modelos.Usuario)
	user.AgregarUsuario(23, "Abdiel Jhadday", time.Now(), true)
	fmt.Println(user)

}
