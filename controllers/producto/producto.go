package producto

import (
	"errors"

	"ggstudios/solerfacturabackend/db_connection"
)

type ProductoDTO struct {
	ID             uint
	TipoProducto   string
	Descripcion    string
	Costo          uint
	CostoEnDolares uint
}

func Create(descripcion string, costo, costoEnDolares, tpr_id uint) (db_connection.Producto, error) {
	producto := db_connection.Producto{Descripcion: descripcion, Costo: costo, TPR_id: tpr_id, CostoEnDolares: costoEnDolares}

	descripcionError := invalidDescripcion(&descripcion)
	costoError := invalidCosto(&costo)
	costoEnDolaresError := invalidCosto(&costoEnDolares)

	switch {
	case descripcionError != nil:
		return producto, descripcionError
	case costoError != nil:
		return producto, costoError
	case costoEnDolaresError != nil:
		return producto, costoEnDolaresError
	case tpr_id == 0:
		return producto, errors.New("el tipo de producto no puede ser cero")
	}

	result := db_connection.Db.Create(&producto)

	if result.Error != nil {
		return producto, result.Error
	}

	return producto, nil
}

func GetAll() ([]ProductoDTO, error) {
	productos := make([]ProductoDTO, 0)

	result := db_connection.Db.Model(&db_connection.Producto{}).
		Select("productos.id, productos.descripcion, productos.costo, productos.costo_en_dolares, tipo_productos.descripcion TipoProducto").
		Joins("left join tipo_productos on tipo_productos.id = productos.tpr_id").Scan(&productos)

	if result.Error != nil {
		return productos, result.Error
	}

	return productos, nil
}

func GetById(id uint) (ProductoDTO, error) {
	producto := new(ProductoDTO)

	result := db_connection.Db.Model(&db_connection.Producto{}).
		Select("productos.id, productos.descripcion, productos.costo, productos.costo_en_dolares, tipo_productos.descripcion TipoProducto").
		Joins("left join tipo_productos on tipo_productos.id = productos.tpr_id").
		Where("productos.id = ?", id).Scan(producto)

	if result.Error != nil {
		return *producto, result.Error
	}

	return *producto, nil
}

func Update(descripcion string, costo, costoEnDolares, tpr_id, id uint) error {
	producto := new(db_connection.Producto)

	result := db_connection.Db.Find(producto, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	descripcionError := invalidDescripcion(&descripcion)
	costoError := invalidCosto(&costo)
	costoEnDolaresError := invalidCosto(&costoEnDolares)

	switch {
	case descripcionError != nil:
		return descripcionError
	case costoError != nil:
		return costoError
	case costoEnDolaresError != nil:
		return costoEnDolaresError
	case tpr_id == 0:
		return errors.New("el tipo de producto no puede ser cero")
	}

	producto.Descripcion = descripcion
	producto.Costo = costo
	producto.CostoEnDolares = costoEnDolares
	producto.TPR_id = tpr_id

	result = db_connection.Db.Save(producto)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id uint) error {
	producto := new(db_connection.Producto)

	result := db_connection.Db.Find(producto, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	result = db_connection.Db.Delete(producto)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func invalidDescripcion(descripcion *string) error {
	if *descripcion == "" {
		return errors.New("la descripcion no puede estar vacia")
	}
	if len(*descripcion) > 100 {
		return errors.New("la descripcion no puede tener mas de 100 caracteres")
	}

	return nil
}

func invalidCosto(costo *uint) error {
	if *costo == 0 {
		return errors.New("el costo no puede ser cero")
	}
	if *costo > 99999999 {
		return errors.New("el costo no puede ser mayor a 99,999,999")
	}

	return nil
}
