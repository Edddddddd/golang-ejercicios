package user

import "time"

// Objetos en Mayuscula es publico y Privado Objeto en Minuscula

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
