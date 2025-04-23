package factura_descripcion

import (
	"fmt"
	"testing"

	"ggstudios/solerfacturabackend/db_connection"
)

func TestGetAll(t *testing.T) {
	db_connection.DbOpen()
	facturas, err := GetById(1)

	if err != nil {
		t.Errorf("Ha ocurrido un error al traer los datos: %v", err)
	}

	for _, factura := range facturas {
		fmt.Println(factura)
	}
	db_connection.CloseDb()
}

func TestCreate(t *testing.T) {
	db_connection.DbOpen()
	productos := []ProductoDTO{
		{
			ID:             1,
			Descripcion:    "Producto 1",
			Costo:          100,
			CostoEnDolares: 10,
			CostoUnitario:  90,
			Cantidad:       2,
			TotalUnitario:  180,
			ITBIS:          true,
		},
		{
			ID:             2,
			Descripcion:    "Producto 2",
			Costo:          200,
			CostoEnDolares: 20,
			CostoUnitario:  180,
			Cantidad:       1,
			TotalUnitario:  180,
			ITBIS:          false,
		},
	}
	_, err := Create(2, productos)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Facturas created successfully")
	}
}

func TestDelete(t *testing.T) {
	db_connection.DbOpen()
	err := Delete(1)
	if err != nil {
		t.Errorf("Ha ocurrido un error al eliminar los datos: %v", err)
	} else {
		fmt.Println("Factura deleted successfully")
	}
	db_connection.CloseDb()
}

func TestUpdate(t *testing.T) {
	db_connection.DbOpen()
	productos := []ProductoDTO{
		{
			ID:             1,
			Descripcion:    "Producto 1",
			Costo:          100,
			CostoEnDolares: 10,
			CostoUnitario:  90,
			Cantidad:       2,
			TotalUnitario:  180,
			ITBIS:          true,
		},
		{
			ID:             2,
			Descripcion:    "Producto 2",
			Costo:          200,
			CostoEnDolares: 20,
			CostoUnitario:  180,
			Cantidad:       1,
			TotalUnitario:  180,
			ITBIS:          false,
		},
	}

	err := Update(1, productos)
	if err != nil {
		t.Errorf("Ha ocurrido un error al actualizar los datos: %v", err)
	} else {
		fmt.Println("Factura updated successfully")
	}
	db_connection.CloseDb()
}
