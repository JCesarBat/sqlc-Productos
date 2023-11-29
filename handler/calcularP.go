package handler

import (
	"context"
	"encoding/json"
	"github/JCesarBat/Proyect_sqlc/db"
	"io"
	"net/http"
)

type Calcu struct {
}

func (c *Calcu) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	productos, err := db.DB.BuscarTProductos(context.Background())
	if err != nil {
		io.WriteString(w, "no se encontro  ningun producto")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		var total uint
		for _, v := range productos {
			total += uint(v.Precio * v.Cantidad)
		}
		json_productos, _ := json.Marshal(total)

		w.Write(json_productos)
		w.WriteHeader(http.StatusOK)
	}

}
