package factura_descripcion

import (
	"errors"
	"fmt"

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

func Create(facturaId uint, productos []ProductoDTO) error {
	facturas := make([]db_connection.FacturaDesc, 0)

	for _, producto := range productos {
		factura := db_connection.FacturaDesc{
			FacturaID:     facturaId,
			ProductoID:    producto.ID,
			CostoUnitario: producto.CostoUnitario,
			Cantidad:      producto.Cantidad,
			TotalUnitario: producto.TotalUnitario,
			ITBIS:         producto.ITBIS,
		}
		facturas = append(facturas, factura)
	}

	result := db_connection.Db.Create(&facturas)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Filas: ", result.RowsAffected)

	return nil
}

func GetById(id uint) ([]ProductoDTO, error) {
	productos := make([]ProductoDTO, 0)
	result := db_connection.Db.Model(&db_connection.FacturaDesc{}).
		Select("productos.id, productos.descripcion, productos.costo, productos.costo_en_dolares, factura_descs.costo_unitario, factura_descs.cantidad, factura_descs.total_unitario, factura_descs.itbis").
		Joins("left join productos on productos.id = factura_descs.producto_id").
		Where("factura_id = ?", id).Scan(&productos)

	if result.Error != nil {
		return productos, result.Error
	}

	fmt.Println(productos)

	return productos, nil
}

func Update(facturaId uint, productos []ProductoDTO) error {
	if Delete(facturaId) != nil {
		return errors.New("no se pudo eliminar los productos de la factura")
	}

	if err := Create(facturaId, productos); err != nil {
		return errors.New("no se pudo crear los productos de la factura")
	}

	return nil
}

func Delete(id uint) error {
	db_connection.Db.Unscoped().Where("factura_id = ?", id).Delete(&db_connection.FacturaDesc{})
	result, _ := GetById(id)
	if len(result) > 0 {
		return errors.New("no se pudo eliminar los productos de la factura")
	}
	return nil
}
