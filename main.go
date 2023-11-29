package main

import (
	"context"
	"github/JCesarBat/Proyect_sqlc/db"
	"github/JCesarBat/Proyect_sqlc/handler"
	"log"
	"net/http"
)

func main() {
	db.Connect()
	db.DB.InsertarEstado(context.Background(), "bueno")
	db.DB.InsertarEstado(context.Background(), "medio")
	db.DB.InsertarEstado(context.Background(), "malo")

	mux := http.NewServeMux()
	mux.Handle("/producto", &handler.Manejador{})
	mux.Handle("/producto/", &handler.Manejador{})
	mux.Handle("/calcular/", &handler.Calcu{})
	log.Fatal(http.ListenAndServe(":3000", mux))

}
