package factura

import (
	"fmt"
	"testing"

	"ggstudios/solerfacturabackend/db_connection"
)

func TestGetAll(t *testing.T) {
	db_connection.DbOpen()
	facturas, err := GetAll()

	if err != nil {
		t.Errorf("Ha ocurrido un error al traer los datos: %v", err)
	}

	for _, factura := range facturas {
		fmt.Println(factura)
	}
	db_connection.CloseDb()
}
