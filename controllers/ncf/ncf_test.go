package ncf

import (
	"testing"

	"ggstudios/solerfacturabackend/db_connection"
)

func TestCreate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful creation", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "03940395"

		created, err := Create(serie, tipo, secuencia)
		deletion := db_connection.NCF{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Unexpected error while creating NCF: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		if created.Serie != serie {
			t.Errorf("Expected serie %q, got %q", serie, created.Serie)
		}
		if created.Tipo != tipo {
			t.Errorf("Expected tipo %q, got %q", tipo, created.Tipo)
		}
		if created.Secuencia != secuencia {
			t.Errorf("Expected secuencia %q, got %q", secuencia, created.Secuencia)
		}
	})

	t.Run("invalid serie", func(t *testing.T) {
		serie := "BB"
		tipo := "01"
		secuencia := "03940395"

		_, err := Create(serie, tipo, secuencia)
		if err == nil {
			t.Error("Expected error for invalid serie, got nil")
		}
		if err.Error() != "solo se pueden ingresar 1 digito en la serie" {
			t.Errorf("Expected error message %q, got %q", "solo se pueden ingresar 1 digito en la serie", err.Error())
		}
	})

	t.Run("invalid tipo", func(t *testing.T) {
		serie := "B"
		tipo := "A1"
		secuencia := "03940395"

		_, err := Create(serie, tipo, secuencia)
		if err == nil {
			t.Error("Expected error for invalid tipo, got nil")
		}
		if err.Error() != "el tipo debe contener solo números" {
			t.Errorf("Expected error message %q, got %q", "el tipo debe contener solo números", err.Error())
		}
	})

	t.Run("invalid secuencia", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "0394039A"

		_, err := Create(serie, tipo, secuencia)
		if err == nil {
			t.Error("Expected error for invalid secuencia, got nil")
		}
		if err.Error() != "la secuencia debe contener solo números" {
			t.Errorf("Expected error message %q, got %q", "la secuencia debe contener solo números", err.Error())
		}
	})

	t.Run("invalid tipo length", func(t *testing.T) {
		serie := "B"
		tipo := "001"
		secuencia := "03940395"

		_, err := Create(serie, tipo, secuencia)
		if err == nil {
			t.Error("Expected error for invalid tipo length, got nil")
		}
		if err.Error() != "solo se pueden ingresar 2 digitos en el tipo" {
			t.Errorf("Expected error message %q, got %q", "solo se pueden ingresar 2 digitos en el tipo", err.Error())
		}
	})

	t.Run("invalid secuencia length", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "039403"

		_, err := Create(serie, tipo, secuencia)
		if err == nil {
			t.Error("Expected error for invalid secuencia length, got nil")
		}
		if err.Error() != "solo se pueden ingresar 8 digitos en la secuencia" {
			t.Errorf("Expected error message %q, got %q", "solo se pueden ingresar 8 digitos en la secuencia", err.Error())
		}
	})
}

func TestGetAll(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("retrieve multiple records", func(t *testing.T) {
		serie1 := "B"
		tipo1 := "01"
		secuencia1 := "03940395"

		created1, err := Create(serie1, tipo1, secuencia1)
		deletion1 := db_connection.NCF{}
		deletion1.ID = created1.ID
		if err != nil {
			t.Fatalf("Failed to create NCF 1: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion1)

		serie2 := "A"
		tipo2 := "02"
		secuencia2 := "12345678"

		created2, err := Create(serie2, tipo2, secuencia2)
		deletion2 := db_connection.NCF{}
		deletion2.ID = created2.ID
		if err != nil {
			t.Fatalf("Failed to create NCF 2: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion2)

		ncfs, err := GetAll()
		if err != nil {
			t.Fatalf("Unexpected error while retrieving NCFs: %v", err)
		}

		if len(ncfs) < 2 {
			t.Fatalf("Expected at least 2 NCFs, got %d", len(ncfs))
		}

		lastNCF1 := ncfs[len(ncfs)-2]
		lastNCF2 := ncfs[len(ncfs)-1]

		if lastNCF1.Serie != serie1 {
			t.Errorf("Expected serie %q, got %q", serie1, lastNCF1.Serie)
		}
		if lastNCF1.Tipo != tipo1 {
			t.Errorf("Expected tipo %q, got %q", tipo1, lastNCF1.Tipo)
		}
		if lastNCF1.Secuencia != secuencia1 {
			t.Errorf("Expected secuencia %q, got %q", secuencia1, lastNCF1.Secuencia)
		}

		if lastNCF2.Serie != serie2 {
			t.Errorf("Expected serie %q, got %q", serie2, lastNCF2.Serie)
		}
		if lastNCF2.Tipo != tipo2 {
			t.Errorf("Expected tipo %q, got %q", tipo2, lastNCF2.Tipo)
		}
		if lastNCF2.Secuencia != secuencia2 {
			t.Errorf("Expected secuencia %q, got %q", secuencia2, lastNCF2.Secuencia)
		}
	})

	t.Run("no records in database", func(t *testing.T) {
		db_connection.Db.Exec("DELETE FROM ncfs")

		ncfs, err := GetAll()
		if err != nil {
			t.Fatalf("Unexpected error while retrieving NCFs: %v", err)
		}

		if len(ncfs) != 0 {
			t.Errorf("Expected 0 NCFs, got %d", len(ncfs))
		}
	})
}

func TestGetById(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("existing record", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "03940395"

		created, err := Create(serie, tipo, secuencia)
		deletion := db_connection.NCF{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Failed to create NCF: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		ncf, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving NCF: %v", err)
		}

		if ncf.Serie != serie {
			t.Errorf("Expected serie %q, got %q", serie, ncf.Serie)
		}
		if ncf.Tipo != tipo {
			t.Errorf("Expected tipo %q, got %q", tipo, ncf.Tipo)
		}
		if ncf.Secuencia != secuencia {
			t.Errorf("Expected secuencia %q, got %q", secuencia, ncf.Secuencia)
		}
	})

	t.Run("non-existent record", func(t *testing.T) {
		ncf, err := GetById(999999)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving ncf: %v", err)
		}
		if ncf.ID != 0 {
			t.Errorf("Expected ncf to be empty, got ID %d", ncf.ID)
		}
	})

	t.Run("invalid ID (zero)", func(t *testing.T) {
		ncf, err := GetById(0)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving ncf: %v", err)
		}
		if ncf.ID != 0 {
			t.Errorf("Expected ncf to be empty, got ID %d", ncf.ID)
		}
	})
}

func TestUpdate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful update", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "03940395"

		created, err := Create(serie, tipo, secuencia)
		deletion := db_connection.NCF{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Failed to create NCF: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		newSerie := "A"
		newTipo := "02"
		newSecuencia := "12345678"

		err = Update(newSerie, newTipo, newSecuencia, created.ID)
		if err != nil {
			t.Fatalf("Failed to update NCF: %v", err)
		}

		updated, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Failed to retrieve updated NCF: %v", err)
		}

		if updated.Serie != newSerie {
			t.Errorf("Expected serie %q, got %q", newSerie, updated.Serie)
		}
		if updated.Tipo != newTipo {
			t.Errorf("Expected tipo %q, got %q", newTipo, updated.Tipo)
		}
		if updated.Secuencia != newSecuencia {
			t.Errorf("Expected secuencia %q, got %q", newSecuencia, updated.Secuencia)
		}
	})

	t.Run("non-existent ID", func(t *testing.T) {
		err := Update("A", "01", "12345678", 999999)
		if err == nil {
			t.Error("Expected error for non-existent ID, got nil")
		}
		if err.Error() != "no se encontro ningun producto" {
			t.Errorf("Expected error message %q, got %q", "no se encontro ningun producto", err.Error())
		}
	})

	t.Run("invalid data", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "03940395"

		created, err := Create(serie, tipo, secuencia)
		if err != nil {
			t.Fatalf("Failed to create NCF: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		invalidSerie := "BB"
		err = Update(invalidSerie, tipo, secuencia, created.ID)
		if err == nil {
			t.Error("Expected error for invalid serie, got nil")
		}
		if err.Error() != "solo se pueden ingresar 1 digito en la serie" {
			t.Errorf("Expected error message %q, got %q", "solo se pueden ingresar 1 digito en la serie", err.Error())
		}
	})
}

func TestDelete(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful delete", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "03940395"

		created, err := Create(serie, tipo, secuencia)
		deletion := db_connection.NCF{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Failed to create NCF: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		err = Delete(created.ID)
		if err != nil {
			t.Fatalf("Failed to delete NCF: %v", err)
		}

		ncf, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving NCF: %v", err)
		}
		if ncf.ID != 0 {
			t.Errorf("Expected NCF to be empty, got ID %d", ncf.ID)
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

func TestDataValidation(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		serie := "b"
		tipo := "01"
		secuencia := "12345678"

		err := dataValidation(&serie, &tipo, &secuencia)
		if err != nil {
			t.Fatalf("Unexpected error for valid data: %v", err)
		}

		if serie != "B" {
			t.Errorf("Expected serie to be uppercase %q, got %q", "B", serie)
		}
	})

	t.Run("invalid serie length", func(t *testing.T) {
		serie := "BB"
		tipo := "01"
		secuencia := "12345678"

		err := dataValidation(&serie, &tipo, &secuencia)
		if err == nil {
			t.Error("Expected error for invalid serie length, got nil")
		}
		if err.Error() != "solo se pueden ingresar 1 digito en la serie" {
			t.Errorf("Expected error message %q, got %q", "solo se pueden ingresar 1 digito en la serie", err.Error())
		}
	})

	t.Run("invalid tipo (non-numeric)", func(t *testing.T) {
		serie := "B"
		tipo := "A1"
		secuencia := "12345678"

		err := dataValidation(&serie, &tipo, &secuencia)
		if err == nil {
			t.Error("Expected error for invalid tipo, got nil")
		}
		if err.Error() != "el tipo debe contener solo números" {
			t.Errorf("Expected error message %q, got %q", "el tipo debe contener solo números", err.Error())
		}
	})

	t.Run("invalid secuencia (non-numeric)", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "12345A78"

		err := dataValidation(&serie, &tipo, &secuencia)
		if err == nil {
			t.Error("Expected error for invalid secuencia, got nil")
		}
		if err.Error() != "la secuencia debe contener solo números" {
			t.Errorf("Expected error message %q, got %q", "la secuencia debe contener solo números", err.Error())
		}
	})

	t.Run("invalid tipo length", func(t *testing.T) {
		serie := "B"
		tipo := "001"
		secuencia := "12345678"

		err := dataValidation(&serie, &tipo, &secuencia)
		if err == nil {
			t.Error("Expected error for invalid tipo length, got nil")
		}
		if err.Error() != "solo se pueden ingresar 2 digitos en el tipo" {
			t.Errorf("Expected error message %q, got %q", "solo se pueden ingresar 2 digitos en el tipo", err.Error())
		}
	})

	t.Run("invalid secuencia length", func(t *testing.T) {
		serie := "B"
		tipo := "01"
		secuencia := "12345"

		err := dataValidation(&serie, &tipo, &secuencia)
		if err == nil {
			t.Error("Expected error for invalid secuencia length, got nil")
		}
		if err.Error() != "solo se pueden ingresar 8 digitos en la secuencia" {
			t.Errorf("Expected error message %q, got %q", "solo se pueden ingresar 8 digitos en la secuencia", err.Error())
		}
	})
}
