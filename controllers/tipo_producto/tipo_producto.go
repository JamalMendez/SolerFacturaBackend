package tipo_producto

import (
	"errors"

	"ggstudios/solerfacturabackend/db_connection"
)

type TipoProductoDTO struct {
	ID          uint
	Descripcion string
}

func Create(descripcion string) (TipoProductoDTO, error) {
	descripcionError := invalidDescripcion(&descripcion)
	if descripcionError != nil {
		return TipoProductoDTO{}, descripcionError
	}

	tipoProducto := db_connection.TipoProducto{Descripcion: descripcion}
	result := db_connection.Db.Create(&tipoProducto)

	if result.Error != nil {
		return TipoProductoDTO{}, result.Error
	}

	return TipoProductoDTO{
		ID:          tipoProducto.ID,
		Descripcion: tipoProducto.Descripcion,
	}, nil
}

func GetAll() ([]TipoProductoDTO, error) {
	tipoProductos := make([]TipoProductoDTO, 0)

	result := db_connection.Db.Model(&db_connection.TipoProducto{}).
		Select("id, descripcion").
		Scan(&tipoProductos)

	if result.Error != nil {
		return tipoProductos, result.Error
	}

	if result.RowsAffected == 0 {
		return tipoProductos, errors.New("no se encontraron productos")
	}

	return tipoProductos, nil
}

func GetById(id uint) (TipoProductoDTO, error) {
	tipoProducto := new(TipoProductoDTO)

	result := db_connection.Db.Model(&db_connection.TipoProducto{}).
		Select("id, descripcion").
		Where("id = ?", id).
		Scan(tipoProducto)

	if result.Error != nil {
		return *tipoProducto, result.Error
	}

	if result.RowsAffected == 0 {
		return *tipoProducto, errors.New("no se encontro ningun producto")
	}

	return *tipoProducto, nil
}

func Update(descripcion string, id uint) error {
	tipoProducto := new(db_connection.TipoProducto)

	result := db_connection.Db.Find(tipoProducto, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	descripcionError := invalidDescripcion(&descripcion)
	if descripcionError != nil {
		return descripcionError
	}

	tipoProducto.Descripcion = descripcion

	result = db_connection.Db.Save(tipoProducto)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id uint) error {
	tipoProducto := new(db_connection.TipoProducto)

	result := db_connection.Db.Find(tipoProducto, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	result = db_connection.Db.Delete(tipoProducto)

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
