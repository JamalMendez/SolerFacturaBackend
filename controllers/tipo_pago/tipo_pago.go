package tipo_pago

import (
	"errors"

	"ggstudios/solerfacturabackend/db_connection"
)

type TipoPagoDTO struct {
	ID          uint
	Descripcion string
}

func Create(descripcion string) (TipoPagoDTO, error) {
	descripcionError := invalidDescripcion(&descripcion)
	if descripcionError != nil {
		return TipoPagoDTO{}, descripcionError
	}

	tipoPago := db_connection.TipoPago{Descripcion: descripcion}
	result := db_connection.Db.Create(&tipoPago)

	if result.Error != nil {
		return TipoPagoDTO{}, result.Error
	}

	return TipoPagoDTO{
		ID:          tipoPago.ID,
		Descripcion: tipoPago.Descripcion,
	}, nil
}

func GetAll() ([]TipoPagoDTO, error) {
	tipoPagos := make([]TipoPagoDTO, 0)

	result := db_connection.Db.Model(&db_connection.TipoPago{}).
		Select("id, descripcion").
		Scan(&tipoPagos)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no se encontraron tipos de pago")
	}

	return tipoPagos, nil
}

func GetById(id uint) (TipoPagoDTO, error) {
	tipoPago := new(TipoPagoDTO)

	result := db_connection.Db.Model(&db_connection.TipoPago{}).
		Select("id, descripcion").
		Where("id = ?", id).
		Scan(tipoPago)

	if result.Error != nil {
		return TipoPagoDTO{}, result.Error
	}

	if result.RowsAffected == 0 {
		return TipoPagoDTO{}, errors.New("no se encontro ningun tipo de pago")
	}

	return *tipoPago, nil
}

func Update(descripcion string, id uint) error {
	tipoPago := new(db_connection.TipoPago)

	result := db_connection.Db.Find(tipoPago, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun tipo de pago")
	}

	descripcionError := invalidDescripcion(&descripcion)
	if descripcionError != nil {
		return descripcionError
	}

	tipoPago.Descripcion = descripcion

	result = db_connection.Db.Save(tipoPago)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id uint) error {
	tipoPago := new(db_connection.TipoPago)

	result := db_connection.Db.Find(tipoPago, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun tipo de pago")
	}

	result = db_connection.Db.Delete(tipoPago)

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
