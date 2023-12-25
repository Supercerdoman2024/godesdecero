package modelos

import (
	"time"
)

type Usuario struct {
	Id     int
	Name   string
	Fecha  time.Time
	Estado bool
}

func (usario *Usuario) AgregarUsuario(id int, name string, fecha time.Time, estado bool) {

	usario.Id = id
	usario.Name = name
	usario.Fecha = fecha
	usario.Estado = estado

}
