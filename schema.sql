
CREATE TABLE estado (
                        ID SERIAL PRIMARY KEY ,
                        descripcion TEXT NOT NULL UNIQUE
);




CREATE TABLE producto  (
    ID SERIAL PRIMARY KEY ,
    Nombre TEXT NOT NULL ,
    Precio INTEGER NOT NULL ,
    Cantidad INTEGER NOT NULL,
    ID_Estado INTEGER NOT NULL

);
ALTER TABLE Producto ADD CONSTRAINT FK_Estado
FOREIGN KEY (ID_Estado) REFERENCES Estado(ID);
