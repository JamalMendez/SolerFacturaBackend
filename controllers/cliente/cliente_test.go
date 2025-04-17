package cliente

import (
	"strings"
	"testing"

	"ggstudios/solerfacturabackend/db_connection"
)

func TestCreate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful creation", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		created, err := Create(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular)
		deletion := db_connection.Cliente{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Unexpected error while creating cliente: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		// Validate the created ClienteDTO
		if created.RNC_Cedula != rnc_cedula {
			t.Errorf("Expected RNC/Cedula %q, got %q", rnc_cedula, created.RNC_Cedula)
		}
		if created.Nombre != nombre {
			t.Errorf("Expected Nombre %q, got %q", nombre, created.Nombre)
		}
		if created.Apellido != apellido {
			t.Errorf("Expected Apellido %q, got %q", apellido, created.Apellido)
		}
		if created.Email != email {
			t.Errorf("Expected Email %q, got %q", email, created.Email)
		}
		if created.Direccion != direccion {
			t.Errorf("Expected Direccion %q, got %q", direccion, created.Direccion)
		}
		if created.Ciudad != ciudad {
			t.Errorf("Expected Ciudad %q, got %q", ciudad, created.Ciudad)
		}
		if created.Telefono != telefono {
			t.Errorf("Expected Telefono %q, got %q", telefono, created.Telefono)
		}
		if created.Celular != celular {
			t.Errorf("Expected Celular %q, got %q", celular, created.Celular)
		}
	})
}

func TestDataValidation(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)
		if err != nil {
			t.Fatalf("Unexpected error for valid data: %v", err)
		}
	})

	t.Run("invalid RNC/Cedula", func(t *testing.T) {
		rnc_cedula := "123"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)
		if err == nil {
			t.Error("Expected error for invalid RNC/Cedula, got nil")
		}
		if err.Error() != "cedula o rnc no es valido" {
			t.Errorf("Expected error message %q, got %q", "cedula o rnc no es valido", err.Error())
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "invalid-email"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)
		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}
		if err.Error() != "el email no es valido" {
			t.Errorf("Expected error message %q, got %q", "el email no es valido", err.Error())
		}
	})

	t.Run("invalid telefono", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "12345"
		celular := "8093435893"

		err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)
		if err == nil {
			t.Error("Expected error for invalid telefono, got nil")
		}
		if err.Error() != "el telefono no es valido" {
			t.Errorf("Expected error message %q, got %q", "el telefono no es valido", err.Error())
		}
	})

	t.Run("invalid celular", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "12345"

		err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)
		if err == nil {
			t.Error("Expected error for invalid celular, got nil")
		}
		if err.Error() != "el celular no es valido" {
			t.Errorf("Expected error message %q, got %q", "el celular no es valido", err.Error())
		}
	})

	t.Run("invalid nombre length", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := ""
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)
		if err == nil {
			t.Error("Expected error for invalid nombre, got nil")
		}
		if err.Error() != "el nombre no es valido" {
			t.Errorf("Expected error message %q, got %q", "el nombre no es valido", err.Error())
		}
	})

	t.Run("invalid direccion length", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := strings.Repeat("a", 201)
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)
		if err == nil {
			t.Error("Expected error for invalid direccion, got nil")
		}
		if err.Error() != "la direccion excede el tamaño permitido" {
			t.Errorf("Expected error message %q, got %q", "la direccion excede el tamaño permitido", err.Error())
		}
	})
}

func TestGetAll(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("retrieve multiple records", func(t *testing.T) {
		// Create first Cliente
		rnc_cedula1 := "40221323212"
		nombre1 := "pepito"
		apellido1 := "veras"
		email1 := "pepitoveras@gmail.com"
		direccion1 := "calle ppe"
		ciudad1 := "Santo Domingo"
		telefono1 := "8093489232"
		celular1 := "8093435893"

		created1, err := Create(rnc_cedula1, nombre1, apellido1, email1, direccion1, ciudad1, telefono1, celular1)
		deletion1 := db_connection.Cliente{}
		deletion1.ID = created1.ID
		if err != nil {
			t.Fatalf("Failed to create Cliente 1: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion1)

		// Create second Cliente
		rnc_cedula2 := "40221323213"
		nombre2 := "juanito"
		apellido2 := "perez"
		email2 := "juanitoperez@gmail.com"
		direccion2 := "calle juan"
		ciudad2 := "Santiago"
		telefono2 := "8093489233"
		celular2 := "8093435894"

		created2, err := Create(rnc_cedula2, nombre2, apellido2, email2, direccion2, ciudad2, telefono2, celular2)
		deletion2 := db_connection.Cliente{}
		deletion2.ID = created2.ID
		if err != nil {
			t.Fatalf("Failed to create Cliente 2: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion2)

		// Retrieve all Clientes
		clientes, err := GetAll()
		if err != nil {
			t.Fatalf("Unexpected error while retrieving Clientes: %v", err)
		}

		// Validate the retrieved Clientes
		if len(clientes) < 2 {
			t.Fatalf("Expected at least 2 Clientes, got %d", len(clientes))
		}

		lastCliente1 := clientes[len(clientes)-2]
		lastCliente2 := clientes[len(clientes)-1]

		if lastCliente1.RNC_Cedula != rnc_cedula1 {
			t.Errorf("Expected RNC/Cedula %q, got %q", rnc_cedula1, lastCliente1.RNC_Cedula)
		}
		if lastCliente2.RNC_Cedula != rnc_cedula2 {
			t.Errorf("Expected RNC/Cedula %q, got %q", rnc_cedula2, lastCliente2.RNC_Cedula)
		}
	})
}

func TestGetById(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("existing record", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		created, err := Create(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular)
		deletion := db_connection.Cliente{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Failed to create Cliente: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		// Retrieve the Cliente by ID
		cliente, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving Cliente: %v", err)
		}

		// Validate the retrieved Cliente
		if cliente.RNC_Cedula != rnc_cedula {
			t.Errorf("Expected RNC/Cedula %q, got %q", rnc_cedula, cliente.RNC_Cedula)
		}
		if cliente.Nombre != nombre {
			t.Errorf("Expected Nombre %q, got %q", nombre, cliente.Nombre)
		}
	})
}

func TestUpdate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful update", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		created, err := Create(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular)
		deletion := db_connection.Cliente{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Failed to create Cliente: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		// New data for update
		newRncCedula := "40221323213"
		newNombre := "juanito"
		newApellido := "perez"
		newEmail := "juanitoperez@gmail.com"
		newDireccion := "calle juan"
		newCiudad := "Santiago"
		newTelefono := "8093489233"
		newCelular := "8093435894"

		// Perform update
		err = Update(newRncCedula, newNombre, newApellido, newEmail, newDireccion, newCiudad, newTelefono, newCelular, created.ID)
		if err != nil {
			t.Fatalf("Failed to update Cliente: %v", err)
		}

		// Retrieve updated Cliente
		updated, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Failed to retrieve updated Cliente: %v", err)
		}

		// Validate updated Cliente
		if updated.RNC_Cedula != newRncCedula {
			t.Errorf("Expected RNC/Cedula %q, got %q", newRncCedula, updated.RNC_Cedula)
		}
		if updated.Nombre != newNombre {
			t.Errorf("Expected Nombre %q, got %q", newNombre, updated.Nombre)
		}
		if updated.Apellido != newApellido {
			t.Errorf("Expected Apellido %q, got %q", newApellido, updated.Apellido)
		}
		if updated.Email != newEmail {
			t.Errorf("Expected Email %q, got %q", newEmail, updated.Email)
		}
		if updated.Direccion != newDireccion {
			t.Errorf("Expected Direccion %q, got %q", newDireccion, updated.Direccion)
		}
		if updated.Ciudad != newCiudad {
			t.Errorf("Expected Ciudad %q, got %q", newCiudad, updated.Ciudad)
		}
		if updated.Telefono != newTelefono {
			t.Errorf("Expected Telefono %q, got %q", newTelefono, updated.Telefono)
		}
		if updated.Celular != newCelular {
			t.Errorf("Expected Celular %q, got %q", newCelular, updated.Celular)
		}
	})

	t.Run("non-existent ID", func(t *testing.T) {
		err := Update("40221323212", "pepito", "veras", "pepitoveras@gmail.com", "calle ppe", "Santo Domingo", "8093489232", "8093435893", 999999)
		if err == nil {
			t.Error("Expected error for non-existent ID, got nil")
		}
		if err.Error() != "no se encontro ningun producto" {
			t.Errorf("Expected error message %q, got %q", "no se encontro ningun producto", err.Error())
		}
	})

	t.Run("invalid data", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		created, err := Create(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular)
		deletion := db_connection.Cliente{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Failed to create Cliente: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		// Invalid email for update
		invalidEmail := "invalid-email"
		err = Update(rnc_cedula, nombre, apellido, invalidEmail, direccion, ciudad, telefono, celular, created.ID)
		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}
		if err.Error() != "el email no es valido" {
			t.Errorf("Expected error message %q, got %q", "el email no es valido", err.Error())
		}
	})
}

func TestDelete(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful delete", func(t *testing.T) {
		rnc_cedula := "40221323212"
		nombre := "pepito"
		apellido := "veras"
		email := "pepitoveras@gmail.com"
		direccion := "calle ppe"
		ciudad := "Santo Domingo"
		telefono := "8093489232"
		celular := "8093435893"

		created, err := Create(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular)
		deletion := db_connection.Cliente{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Failed to create Cliente: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		// Delete the Cliente
		err = Delete(created.ID)
		if err != nil {
			t.Fatalf("Failed to delete Cliente: %v", err)
		}

		// Verify the Cliente is deleted
		cliente, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving Cliente: %v", err)
		}
		if cliente.ID != 0 {
			t.Errorf("Expected Cliente to be empty, got ID %d", cliente.ID)
		}
	})

	t.Run("non-existent ID", func(t *testing.T) {
		err := Delete(999999)
		if err == nil {
			t.Error("Expected error for non-existent ID, got nil")
		}
		if err.Error() != "no se encontro ningun producto" {
			t.Errorf("Expected error message %q, got %q", "no se encontro ningun producto", err.Error())
		}
	})

	t.Run("invalid ID (zero)", func(t *testing.T) {
		err := Delete(0)
		if err == nil {
			t.Error("Expected error for invalid ID (zero), got nil")
		}
		if err.Error() != "no se encontro ningun producto" {
			t.Errorf("Expected error message %q, got %q", "no se encontro ningun producto", err.Error())
		}
	})
}
