// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import ()

type Estado struct {
	ID          int32
	Descripcion string
}

type Producto struct {
	ID       int32
	Nombre   string
	Precio   int32
	Cantidad int32
	IDEstado int32
}
