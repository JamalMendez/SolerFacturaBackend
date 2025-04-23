package cotizacion_descripcion

import (
	"errors"

	"ggstudios/solerfacturabackend/db_connection"
)

type ProductoDTO struct {
	ID             uint
	Descripcion    string
	Costo          uint
	CostoEnDolares uint
	CostoUnitario  uint
	Cantidad       uint
	TotalUnitario  uint
	ITBIS          bool
}

func Create(cotizacionId uint, productos []ProductoDTO) ([]db_connection.CotizacionDesc, error) {
	var cotizacions []db_connection.CotizacionDesc

	for _, producto := range productos {
		cotizacions = append(cotizacions, db_connection.CotizacionDesc{
			CotizacionID:  cotizacionId,
			ProductoID:    producto.ID,
			CostoUnitario: producto.CostoUnitario,
			Cantidad:      producto.Cantidad,
			TotalUnitario: producto.TotalUnitario,
			ITBIS:         producto.ITBIS,
		})
	}

	if err := db_connection.Db.Create(&cotizacions).Error; err != nil {
		return nil, err
	}

	return cotizacions, nil
}

func GetById(id uint) ([]ProductoDTO, error) {
	productos := make([]ProductoDTO, 0)
	result := db_connection.Db.Model(&db_connection.CotizacionDesc{}).
		Select("productos.id, productos.descripcion, productos.costo, productos.costo_en_dolares, cotizacion_descs.costo_unitario, cotizacion_descs.cantidad, cotizacion_descs.total_unitario, cotizacion_descs.itbis").
		Joins("left join productos on productos.id = cotizacion_descs.producto_id").
		Where("cotizacion_id = ?", id).Scan(&productos)

	if result.Error != nil {
		return productos, result.Error
	}

	if result.RowsAffected == 0 {
		return productos, errors.New("no se encontraron productos para la cotizacion")
	}

	return productos, nil
}

func Update(cotizacionId uint, productos []ProductoDTO) error {
	if Delete(cotizacionId) != nil {
		return errors.New("no se pudo eliminar los productos de la cotizacion")
	}

	if _, err := Create(cotizacionId, productos); err != nil {
		return errors.New("no se pudo crear los productos de la cotizacion")
	}

	return nil
}

func Delete(id uint) error {
	db_connection.Db.Unscoped().Where("cotizacion_id = ?", id).Delete(&db_connection.CotizacionDesc{})
	result, _ := GetById(id)
	if len(result) > 0 {
		return errors.New("no se pudo eliminar los productos de la cotizacion")
	}
	return nil
}
