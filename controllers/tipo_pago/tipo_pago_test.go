package tipo_pago

import (
	"strings"
	"testing"
	"time"

	"ggstudios/solerfacturabackend/db_connection"
)

func TestCreate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful creation", func(t *testing.T) {
		descripcion := "test_valid_" + time.Now().Format("20060102150405")
		created, err := Create(descripcion)
		deletion := db_connection.TipoPago{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		if created.Descripcion != descripcion {
			t.Errorf("Expected description %q, got %q", descripcion, created.Descripcion)
		}
	})

	t.Run("empty description", func(t *testing.T) {
		_, err := Create("")
		if err == nil {
			t.Error("Expected error for empty description, got nil")
		}
	})

	t.Run("duplicate description", func(t *testing.T) {
		descripcion := "test_duplicate_" + time.Now().Format("20060102150405")

		first, err := Create(descripcion)
		deletion := db_connection.TipoPago{}
		deletion.ID = first.ID
		if err != nil {
			t.Fatalf("Setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		_, err = Create(descripcion)
		if err == nil {
			t.Error("Expected error for duplicate description, got nil")
		}
	})

	t.Run("exceeds max description length", func(t *testing.T) {
		longDesc := strings.Repeat("a", 101)
		_, err := Create(longDesc)
		if err == nil {
			t.Error("Expected error for long description, got nil")
		}
	})
}

func TestGetAll(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("retrieve multiple records", func(t *testing.T) {
		desc1 := "test_getall_1_" + time.Now().Format(time.RFC3339Nano)
		desc2 := "test_getall_2_" + time.Now().Format(time.RFC3339Nano)

		tp1, err := Create(desc1)
		deletion1 := db_connection.TipoPago{}
		deletion1.ID = tp1.ID
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion1)

		tp2, err := Create(desc2)
		deletion2 := db_connection.TipoPago{}
		deletion2.ID = tp2.ID
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion2)

		result, err := GetAll()
		if err != nil {
			t.Fatalf("GetAll failed: %v", err)
		}

		foundCount := 0
		for _, item := range result {
			if item.ID == tp1.ID && item.Descripcion == desc1 {
				foundCount++
			}
			if item.ID == tp2.ID && item.Descripcion == desc2 {
				foundCount++
			}
		}

		if foundCount != 2 {
			t.Errorf("Expected 2 test records, found %d", foundCount)
		}
	})
}

func TestGetById(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("existing record", func(t *testing.T) {
		desc := "test_getbyid_" + time.Now().Format(time.RFC3339Nano)
		created, err := Create(desc)
		deletion := db_connection.TipoPago{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		fetched, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("GetById failed: %v", err)
		}

		if fetched.ID != created.ID || fetched.Descripcion != desc {
			t.Errorf("Mismatched data\nExpected: ID=%d, Desc=%s\nGot: ID=%d, Desc=%s",
				created.ID, desc, fetched.ID, fetched.Descripcion)
		}
	})

	t.Run("non-existing record", func(t *testing.T) {
		_, err := GetById(999999999)
		if err != nil {
			t.Errorf("Expected no error for non-existing ID, got: %v", err)
		}
	})

	t.Run("invalid ID (zero)", func(t *testing.T) {
		_, err := GetById(0)
		if err != nil {
			t.Errorf("Expected no error for zero ID, got: %v", err)
		}
	})

	t.Run("deleted record", func(t *testing.T) {
		desc := "test_deleted_" + time.Now().Format(time.RFC3339Nano)
		created, err := Create(desc)
		deletion := db_connection.TipoPago{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		db_connection.Db.Unscoped().Delete(&deletion)

		// Attempt to retrieve deleted record
		fetched, err := GetById(created.ID)
		if err != nil {
			t.Errorf("Expected no error for deleted record, got: %v", err)
		}
		if fetched.ID != 0 {
			t.Errorf("Expected empty record for deleted entry, got ID %d", fetched.ID)
		}
	})
}

func TestUpdate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful update", func(t *testing.T) {
		originalDesc := "test_update_original_" + time.Now().Format(time.RFC3339Nano)
		updatedDesc := "test_update_modified_" + time.Now().Format(time.RFC3339Nano)

		created, err := Create(originalDesc)
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		err = Update(updatedDesc, created.ID)
		if err != nil {
			t.Fatalf("Update failed: %v", err)
		}

		updated, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Verification failed: %v", err)
		}

		if updated.Descripcion != updatedDesc {
			t.Errorf("Update failed\nExpected: %s\nGot: %s", updatedDesc, updated.Descripcion)
		}
	})

	t.Run("non-existent ID", func(t *testing.T) {
		invalidID := uint(999999999)
		err := Update("new_description", invalidID)
		if err == nil || err.Error() != "no se encontro ningun tipo de pago" {
			t.Errorf("Expected 'not found' error, got: %v", err)
		}
	})

	t.Run("empty description", func(t *testing.T) {
		created, err := Create("test_empty_desc_" + time.Now().Format(time.RFC3339Nano))
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		err = Update("", created.ID)
		if err == nil {
			t.Error("Expected error for empty description, got nil")
		}
	})

	t.Run("exceeds max description length", func(t *testing.T) {
		created, err := Create("test_long_desc_" + time.Now().Format(time.RFC3339Nano))
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		longDesc := strings.Repeat("a", 101)
		err = Update(longDesc, created.ID)
		if err == nil {
			t.Error("Expected error for long description, got nil")
		}
	})

	t.Run("duplicate description", func(t *testing.T) {
		desc1 := "test_dup_1_" + time.Now().Format(time.RFC3339Nano)
		desc2 := "test_dup_2_" + time.Now().Format(time.RFC3339Nano)

		record1, err := Create(desc1)
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&record1)

		record2, err := Create(desc2)
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&record2)

		err = Update(desc1, record2.ID)
		if err == nil {
			t.Error("Expected error for duplicate description, got nil")
		}
	})

	t.Run("no changes update", func(t *testing.T) {
		originalDesc := "test_nochange_" + time.Now().Format(time.RFC3339Nano)
		created, err := Create(originalDesc)
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		err = Update(originalDesc, created.ID)
		if err != nil {
			t.Errorf("Expected successful update with same description, got error: %v", err)
		}
	})
}

func TestDelete(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful delete", func(t *testing.T) {
		desc := "test_delete_valid_" + time.Now().Format(time.RFC3339Nano)
		created, err := Create(desc)
		deletion := db_connection.TipoPago{}
		deletion.ID = created.ID
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&deletion)

		err = Delete(created.ID)
		if err != nil {
			t.Fatalf("Delete failed: %v", err)
		}

		var deletedRecord db_connection.TipoPago
		result := db_connection.Db.Unscoped().Where("id = ?", created.ID).First(&deletedRecord)
		if result.Error != nil {
			t.Fatal("Record should exist after delete")
		}
		if deletedRecord.DeletedAt.Time.IsZero() {
			t.Error("DeletedAt timestamp not set")
		}
	})

	t.Run("non-existent ID", func(t *testing.T) {
		invalidID := uint(999999999)
		err := Delete(invalidID)
		expectedErr := "no se encontro ningun tipo de pago"

		if err == nil || err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})

	t.Run("already deleted record", func(t *testing.T) {
		desc := "test_delete_twice_" + time.Now().Format(time.RFC3339Nano)
		created, err := Create(desc)
		if err != nil {
			t.Fatalf("Test setup failed: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		if err := Delete(created.ID); err != nil {
			t.Fatalf("First delete failed: %v", err)
		}

		err = Delete(created.ID)
		expectedErr := "no se encontro ningun tipo de pago"
		if err == nil || err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})

	t.Run("invalid ID (zero)", func(t *testing.T) {
		err := Delete(0)
		expectedErr := "no se encontro ningun tipo de pago"

		if err == nil || err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})
}

func TestInvalidDescripcion(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	t.Run("empty description", func(t *testing.T) {
		emptyDesc := ""
		expectedErr := "la descripcion no puede estar vacia"

		err := invalidDescripcion(strPtr(emptyDesc))
		if err == nil || err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})

	t.Run("description too long", func(t *testing.T) {
		longDesc := strings.Repeat("a", 101)
		expectedErr := "la descripcion no puede tener mas de 100 caracteres"

		err := invalidDescripcion(strPtr(longDesc))
		if err == nil || err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})

	t.Run("exact length 100", func(t *testing.T) {
		validDesc := strings.Repeat("a", 100)

		err := invalidDescripcion(strPtr(validDesc))
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	})

	t.Run("valid description", func(t *testing.T) {
		validDesc := "Valid description"

		err := invalidDescripcion(strPtr(validDesc))
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	})

	t.Run("borderline case (99 characters)", func(t *testing.T) {
		validDesc := strings.Repeat("a", 99)

		err := invalidDescripcion(strPtr(validDesc))
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	})

	t.Run("borderline case (100 characters)", func(t *testing.T) {
		validDesc := strings.Repeat("a", 100)

		err := invalidDescripcion(strPtr(validDesc))
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	})

	t.Run("borderline case (101 characters)", func(t *testing.T) {
		invalidDesc := strings.Repeat("a", 101)
		expectedErr := "la descripcion no puede tener mas de 100 caracteres"

		err := invalidDescripcion(strPtr(invalidDesc))
		if err == nil || err.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
		}
	})
}
