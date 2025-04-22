package cotizacion

import (
	"fmt"
	"testing"

	"ggstudios/solerfacturabackend/db_connection"
)

func TestGetAll(t *testing.T) {
	db_connection.DbOpen()
	cotizacion, err := GetAll()

	if err != nil {
		t.Errorf("Ha ocurrido un error al traer los datos: %v", err)
	}

	for _, factura := range cotizacion {
		fmt.Println(factura)
	}
	db_connection.CloseDb()
}

func TestCreate(t *testing.T) {
	db_connection.DbOpen()
	cli_id := uint(1)
	tpo_id := uint(1)
	costoSubtotal := uint(1000)
	costoTotal := uint(1100)
	descuento := uint(100)
	envio := uint(50)
	secuencia := "12345678"
	cliente := "Cliente Example"
	descripcion := "Descripcion Example"
	enDolares := false

	cotizacion, err := Create(cli_id, tpo_id, costoSubtotal, costoTotal, descuento, envio, secuencia, cliente, descripcion, enDolares)
	if err != nil {
		t.Errorf("Error creating cotizacion: %v", err)
	}

	fmt.Println("Created cotizacion:", cotizacion)
	db_connection.CloseDb()
}

// Data to be inserted
