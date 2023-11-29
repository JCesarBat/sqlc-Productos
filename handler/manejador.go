package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github/JCesarBat/Proyect_sqlc/db"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var producto = regexp.MustCompile(`^/producto/*$`)
var producto_id = regexp.MustCompile(`^/producto/([0-9]+)$`)

type Manejador struct {
}

func (p *Manejador) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && producto_id.MatchString(r.URL.Path):
		obtenerid(w, r)
		return
	case r.Method == http.MethodGet && producto.MatchString(r.URL.Path):
		obtenerT(w, r)
		return
	case r.Method == http.MethodPost && producto.MatchString(r.URL.Path):
		CrearP(w, r)
		return
	case r.Method == http.MethodDelete && producto_id.MatchString(r.URL.Path):
		Eliminar(w, r)
		return

	default:
		io.WriteString(w, " no encotro la direccion")
		DefaultError(w, r)
	}

}

func obtenerid(w http.ResponseWriter, r *http.Request) {
	number_path := strings.Split(r.URL.Path, "/")
	ID, err := strconv.ParseUint(number_path[2], 10, 32)

	if err != nil {
		io.WriteString(w, "se encontro un problema al obtener la direccion de el url")
		DefaultError(w, r)

	} else {
		producto, err := db.DB.BuscarProducto(context.Background(), int32(ID))

		if err != nil {
			io.WriteString(w, "no se encontro ese producto")
			DefaultError(w, r)
		} else {

			json_producto, err := json.Marshal(producto)
			if err != nil {
				io.WriteString(w, " error de conversion")
				DefaultError(w, r)
			} else {

				w.Write(json_producto)
				w.WriteHeader(http.StatusOK)
			}

		}
	}
}
func obtenerT(w http.ResponseWriter, r *http.Request) {
	productos, err := db.DB.BuscarTProductos(context.Background())
	if err != nil {
		io.WriteString(w, " no se encontro ningun producto")
		DefaultError(w, r)
	} else {
		json_productos, err := json.Marshal(productos)
		if err != nil {
			io.WriteString(w, " error al convertir json")
			DefaultError(w, r)
		} else {
			w.Write(json_productos)
			w.WriteHeader(http.StatusOK)
		}
	}
}
func CrearP(w http.ResponseWriter, r *http.Request) {

	estados, err := db.DB.BuscarTodosEstados(context.Background())
	if err != nil {
		io.WriteString(w, "no encontro ningun estado ")
		DefaultError(w, r)
	} else {
		existe_status := false
		mapa := map[string]string{}
		json.NewDecoder(r.Body).Decode(&mapa)
		estado := mapa["estado"]
		fmt.Println(estado)
		for _, estatus := range estados {

			if estatus.Descripcion == estado {
				existe_status = true

				Nombre := mapa["Nombre"]
				Precio := mapa["Precio"]
				Cantidad := mapa["Cantidad"]

				if Nombre == "" || Precio == "" || Precio == "" || Cantidad == "" {
					io.WriteString(w, " los datos ingresados son incorrectos un producto tiene que tener"+
						" un Nombre un Precio una cantidad con un numero mayor o igual que 1 y el estado")
					DefaultError(w, r)
				} else {
					_precio, err := strconv.ParseUint(Precio, 10, 32)
					if err != nil || _precio <= 0 {
						io.WriteString(w, "el precio no puede ser un string ni menor que 0 ")
						DefaultError(w, r)
					} else {
						_canitdad, err := strconv.ParseUint(Cantidad, 10, 32)
						if err != nil || _canitdad <= 0 {
							io.WriteString(w, " la cantidad  no puede ser un string ni menor que 0 ")
							DefaultError(w, r)
						} else {
							_, err := db.DB.InsertarProducto(context.Background(), db.InsertarProductoParams{
								Nombre:   Nombre,
								Precio:   int32(_precio),
								Cantidad: int32(_canitdad),
								IDEstado: estatus.ID,
							})
							if err != nil {
								io.WriteString(w, err.Error())
								DefaultError(w, r)
							} else {
								io.WriteString(w, " se ingreso el producto con exito")
								w.WriteHeader(http.StatusOK)
								break
							}
						}
					}
				}
			}
		}
		if existe_status == false {
			io.WriteString(w, "el estado ingresado es incorrecto o no se ingreso este")
			DefaultError(w, r)
		}
	}
}
func Eliminar(w http.ResponseWriter, r *http.Request) {
	number_path := strings.Split(r.URL.Path, "/")
	ID, err := strconv.ParseUint(number_path[2], 10, 32)
	if err != nil {
		io.WriteString(w, " error en la direccion")
	} else {
		_, err := db.DB.Eliminar(context.Background(), int32(ID))
		if err != nil {
			io.WriteString(w, " no se pudo eliminar a el producto o no se ingreso")
			DefaultError(w, r)
		} else {
			io.WriteString(w, "se elimino con exito el producto")
			w.WriteHeader(http.StatusOK)
		}
	}
}

func DefaultError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
