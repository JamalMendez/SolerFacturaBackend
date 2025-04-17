package producto

import (
	"strings"
	"testing"
	"time"

	"ggstudios/solerfacturabackend/controllers/tipo_producto"
	"ggstudios/solerfacturabackend/db_connection"
)

func TestCreate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful creation", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&tipoProducto)

		descripcion := "test_producto_" + time.Now().Format("20060102150405")
		var costo uint = 5000
		var costoEnDolares uint = 100

		created, err := Create(descripcion, costo, costoEnDolares, tipoProducto.ID)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		if created.Descripcion != descripcion {
			t.Errorf("Expected description %q, got %q", descripcion, created.Descripcion)
		}
		if created.Costo != costo {
			t.Errorf("Expected cost %d, got %d", costo, created.Costo)
		}
		if created.CostoEnDolares != costoEnDolares {
			t.Errorf("Expected cost in dollars %d, got %d", costoEnDolares, created.CostoEnDolares)
		}
		if created.TPR_id != tipoProducto.ID {
			t.Errorf("Expected type ID %d, got %d", tipoProducto.ID, created.TPR_id)
		}
	})

	t.Run("empty description", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&tipoProducto)

		var costo uint = 5000
		var costoEnDolares uint = 100

		_, err = Create("", costo, costoEnDolares, tipoProducto.ID)
		if err == nil {
			t.Error("Expected error for empty description, got nil")
		}
	})
}

func TestGetAll(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("retrieve multiple records", func(t *testing.T) {
		tipoProductoDescripcion1 := "test_tipo_producto_1_" + time.Now().Format("20060102150405")
		tipoProducto1, err := tipo_producto.Create(tipoProductoDescripcion1)
		if err != nil {
			t.Fatalf("Failed to create type product 1: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&tipoProducto1)

		tipoProductoDescripcion2 := "test_tipo_producto_2_" + time.Now().Format("20060102150405")
		tipoProducto2, err := tipo_producto.Create(tipoProductoDescripcion2)
		if err != nil {
			t.Fatalf("Failed to create type product 2: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&tipoProducto2)

		descripcion1 := "test_producto_1_" + time.Now().Format("20060102150405")
		descripcion2 := "test_producto_2_" + time.Now().Format("20060102150405")
		var costo1, costo2 uint = 5000, 7000
		var costoEnDolares1, costoEnDolares2 uint = 100, 140

		created1, err := Create(descripcion1, costo1, costoEnDolares1, tipoProducto1.ID)
		if err != nil {
			t.Fatalf("Failed to create product 1: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created1)

		created2, err := Create(descripcion2, costo2, costoEnDolares2, tipoProducto2.ID)
		if err != nil {
			t.Fatalf("Failed to create product 2: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created2)

		productos, err := GetAll()
		if err != nil {
			t.Fatalf("Unexpected error while retrieving products: %v", err)
		}

		if len(productos) < 2 {
			t.Fatalf("Expected at least 2 products, got %d", len(productos))
		}

		lastProduct1 := productos[len(productos)-2]
		lastProduct2 := productos[len(productos)-1]

		if lastProduct1.Descripcion != descripcion1 {
			t.Errorf("Expected description %q, got %q", descripcion1, lastProduct1.Descripcion)
		}
		if lastProduct1.Costo != costo1 {
			t.Errorf("Expected cost %d, got %d", costo1, lastProduct1.Costo)
		}
		if lastProduct1.CostoEnDolares != costoEnDolares1 {
			t.Errorf("Expected cost in dollars %d, got %d", costoEnDolares1, lastProduct1.CostoEnDolares)
		}
		if lastProduct1.TipoProducto != tipoProductoDescripcion1 {
			t.Errorf("Expected type product %q, got %q", tipoProductoDescripcion1, lastProduct1.TipoProducto)
		}

		if lastProduct2.Descripcion != descripcion2 {
			t.Errorf("Expected description %q, got %q", descripcion2, lastProduct2.Descripcion)
		}
		if lastProduct2.Costo != costo2 {
			t.Errorf("Expected cost %d, got %d", costo2, lastProduct2.Costo)
		}
		if lastProduct2.CostoEnDolares != costoEnDolares2 {
			t.Errorf("Expected cost in dollars %d, got %d", costoEnDolares2, lastProduct2.CostoEnDolares)
		}
		if lastProduct2.TipoProducto != tipoProductoDescripcion2 {
			t.Errorf("Expected type product %q, got %q", tipoProductoDescripcion2, lastProduct2.TipoProducto)
		}
	})

	t.Run("no records in database", func(t *testing.T) {
		db_connection.Db.Exec("DELETE FROM SisFac.dbo.productos")

		productos, err := GetAll()
		if err != nil {
			t.Fatalf("Unexpected error while retrieving products: %v", err)
		}

		if len(productos) != 0 {
			t.Errorf("Expected 0 products, got %d", len(productos))
		}
	})
}

func TestGetById(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("existing record", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&tipoProducto)

		descripcion := "test_producto_" + time.Now().Format("20060102150405")
		var costo uint = 5000
		var costoEnDolares uint = 100

		created, err := Create(descripcion, costo, costoEnDolares, tipoProducto.ID)
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&created)

		producto, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving product by ID: %v", err)
		}

		if producto.Descripcion != descripcion {
			t.Errorf("Expected description %q, got %q", descripcion, producto.Descripcion)
		}
		if producto.Costo != costo {
			t.Errorf("Expected cost %d, got %d", costo, producto.Costo)
		}
		if producto.CostoEnDolares != costoEnDolares {
			t.Errorf("Expected cost in dollars %d, got %d", costoEnDolares, producto.CostoEnDolares)
		}
		if producto.TipoProducto != tipoProducto.Descripcion {
			t.Errorf("Expected type ID %q, got %q", tipoProducto.Descripcion, producto.TipoProducto)
		}
	})

	t.Run("non-existing record", func(t *testing.T) {
		producto, err := GetById(999999)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving product: %v", err)
		}
		if producto.ID != 0 {
			t.Errorf("Expected product to be empty, got ID %d", producto.ID)
		}
	})

	t.Run("invalid ID (zero)", func(t *testing.T) {
		producto, err := GetById(0)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving product: %v", err)
		}
		if producto.ID != 0 {
			t.Errorf("Expected product to be empty, got ID %d", producto.ID)
		}
	})

	t.Run("deleted record", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}

		descripcion := "test_producto_" + time.Now().Format("20060102150405")
		var costo uint = 5000
		var costoEnDolares uint = 100

		created, err := Create(descripcion, costo, costoEnDolares, tipoProducto.ID)
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		err = Delete(created.ID)
		if err != nil {
			t.Fatalf("Failed to delete product: %v", err)
		}

		producto, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving product: %v", err)
		}
		if producto.ID != 0 {
			t.Errorf("Expected product to be empty, got ID %d", producto.ID)
		}

		db_connection.Db.Unscoped().Delete(&created)
		db_connection.Db.Unscoped().Delete(&tipoProducto)
	})
}

func TestUpdate(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful update", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}

		descripcion := "test_producto_" + time.Now().Format("20060102150405")
		var costo uint = 5000
		var costoEnDolares uint = 100

		created, err := Create(descripcion, costo, costoEnDolares, tipoProducto.ID)
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		newDescripcion := "updated_producto_" + time.Now().Format("20060102150405")
		var newCosto uint = 7000
		var newCostoEnDolares uint = 150
		newTipoProductoDescripcion := "updated_tipo_producto_" + time.Now().Format("20060102150405")

		newTipoProducto, err := tipo_producto.Create(newTipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create new type product: %v", err)
		}

		err = Update(newDescripcion, newCosto, newCostoEnDolares, newTipoProducto.ID, created.ID)
		if err != nil {
			t.Fatalf("Failed to update product: %v", err)
		}

		updated, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Failed to retrieve updated product: %v", err)
		}

		if updated.Descripcion != newDescripcion {
			t.Errorf("Expected description %q, got %q", newDescripcion, updated.Descripcion)
		}
		if updated.Costo != newCosto {
			t.Errorf("Expected cost %d, got %d", newCosto, updated.Costo)
		}
		if updated.CostoEnDolares != newCostoEnDolares {
			t.Errorf("Expected cost in dollars %d, got %d", newCostoEnDolares, updated.CostoEnDolares)
		}
		if updated.TipoProducto != newTipoProductoDescripcion {
			t.Errorf("Expected type product %q, got %q", newTipoProductoDescripcion, updated.TipoProducto)
		}

		db_connection.Db.Unscoped().Delete(&created)
		db_connection.Db.Unscoped().Delete(&tipoProducto)
		db_connection.Db.Unscoped().Delete(&newTipoProducto)
	})

	t.Run("non-existent product ID", func(t *testing.T) {
		err := Update("non_existent", 5000, 100, 1, 999999)
		if err == nil {
			t.Error("Expected error for non-existent product ID, got nil")
		}
	})

	t.Run("invalid type product ID", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}

		descripcion := "test_producto_" + time.Now().Format("20060102150405")
		var costo uint = 5000
		var costoEnDolares uint = 100

		created, err := Create(descripcion, costo, costoEnDolares, tipoProducto.ID)
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		err = Update("updated_producto", 7000, 150, 0, created.ID)
		if err == nil {
			t.Error("Expected error for invalid type product ID, got nil")
		}

		db_connection.Db.Unscoped().Delete(&created)
		db_connection.Db.Unscoped().Delete(&tipoProducto)
	})

	t.Run("empty description", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}

		descripcion := "test_producto_" + time.Now().Format("20060102150405")
		var costo uint = 5000
		var costoEnDolares uint = 100

		created, err := Create(descripcion, costo, costoEnDolares, tipoProducto.ID)
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		err = Update("", 7000, 150, tipoProducto.ID, created.ID)
		if err == nil {
			t.Error("Expected error for empty description, got nil")
		}

		db_connection.Db.Unscoped().Delete(&created)
		db_connection.Db.Unscoped().Delete(&tipoProducto)
	})
}

func TestDelete(t *testing.T) {
	db_connection.DbOpen()
	defer db_connection.CloseDb()

	t.Run("successful delete", func(t *testing.T) {
		tipoProductoDescripcion := "test_tipo_producto_" + time.Now().Format("20060102150405")
		tipoProducto, err := tipo_producto.Create(tipoProductoDescripcion)
		if err != nil {
			t.Fatalf("Failed to create type product: %v", err)
		}
		defer db_connection.Db.Unscoped().Delete(&tipoProducto)

		descripcion := "test_producto_" + time.Now().Format("20060102150405")
		var costo uint = 5000
		var costoEnDolares uint = 100

		created, err := Create(descripcion, costo, costoEnDolares, tipoProducto.ID)
		if err != nil {
			t.Fatalf("Failed to create product: %v", err)
		}

		err = Delete(created.ID)
		if err != nil {
			t.Fatalf("Failed to delete product: %v", err)
		}

		producto, err := GetById(created.ID)
		if err != nil {
			t.Fatalf("Unexpected error while retrieving product: %v", err)
		}
		if producto.ID != 0 {
			t.Errorf("Expected product to be empty, got ID %d", producto.ID)
		}
	})

	t.Run("non-existent product ID", func(t *testing.T) {
		err := Delete(999999)
		if err == nil {
			t.Error("Expected error for non-existent product ID, got nil")
		}
	})

	t.Run("invalid product ID (zero)", func(t *testing.T) {
		err := Delete(0)
		if err == nil {
			t.Error("Expected error for invalid product ID (zero), got nil")
		}
	})
}

func TestInvalidDescripcion(t *testing.T) {
	t.Run("empty description", func(t *testing.T) {
		descripcion := ""
		err := invalidDescripcion(&descripcion)
		if err == nil {
			t.Error("Expected error for empty description, got nil")
		}
		if err.Error() != "la descripcion no puede estar vacia" {
			t.Errorf("Expected error message %q, got %q", "la descripcion no puede estar vacia", err.Error())
		}
	})

	t.Run("description exceeds max length", func(t *testing.T) {
		descripcion := "a" + strings.Repeat("a", 100) // 101 characters
		err := invalidDescripcion(&descripcion)
		if err == nil {
			t.Error("Expected error for description exceeding max length, got nil")
		}
		if err.Error() != "la descripcion no puede tener mas de 100 caracteres" {
			t.Errorf("Expected error message %q, got %q", "la descripcion no puede tener mas de 100 caracteres", err.Error())
		}
	})

	t.Run("valid description", func(t *testing.T) {
		descripcion := "Valid description"
		err := invalidDescripcion(&descripcion)
		if err != nil {
			t.Errorf("Unexpected error for valid description: %v", err)
		}
	})
}

func TestInvalidCosto(t *testing.T) {
	t.Run("zero cost", func(t *testing.T) {
		var costo uint = 0
		err := invalidCosto(&costo)
		if err == nil {
			t.Error("Expected error for zero cost, got nil")
		}
		if err.Error() != "el costo no puede ser cero" {
			t.Errorf("Expected error message %q, got %q", "el costo no puede ser cero", err.Error())
		}
	})

	t.Run("cost exceeds max value", func(t *testing.T) {
		var costo uint = 100000000 // Greater than 99,999,999
		err := invalidCosto(&costo)
		if err == nil {
			t.Error("Expected error for cost exceeding max value, got nil")
		}
		if err.Error() != "el costo no puede ser mayor a 99,999,999" {
			t.Errorf("Expected error message %q, got %q", "el costo no puede ser mayor a 99,999,999", err.Error())
		}
	})

	t.Run("valid cost", func(t *testing.T) {
		var costo uint = 5000
		err := invalidCosto(&costo)
		if err != nil {
			t.Errorf("Unexpected error for valid cost: %v", err)
		}
	})
}
