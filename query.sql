-- name: BuscarTProductos :many
SELECT * FROM producto ;

-- name: BuscarProducto :one
SELECT * FROM Producto WHERE ID = $1 LIMIT 1;

-- name: InsertarProducto :one
INSERT INTO producto (Nombre,Precio,Cantidad,ID_Estado)VALUES ($1,$2,$3,$4)RETURNING *;

-- name: InsertarEstado :one
INSERT INTO estado(descripcion)VALUES ($1) RETURNING *;


-- name: BuscarTodosEstados :many
SELECT * FROM estado ;

-- name: Eliminar :one
DELETE FROM producto WHERE ID = $1  RETURNING *;
