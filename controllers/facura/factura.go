package facura

import (
	"errors"
	"time"

	"ggstudios/solerfacturabackend/db_connection"
)

type FacturaDTO struct {
	ID               uint
	NCFSerie         string
	NCFTipo          string
	NCFSecuencia     string
	TipoPago         string
	Cliente          string
	CostoSubtotal    uint
	CostoTotal       uint
	Descuento        uint
	Envio            uint
	Descripcion      string
	FechaCreacion    time.Time
	FechaVencimiento time.Time
	EnDolares        bool
}

func Create(ncf_id, cli_id, tpo_id, costoSubtotal, costoTotal, descuento, envio uint, cliente, descripcion string, enDolares bool, fehcaVencimiento time.Time) (db_connection.Factura, error) {
	factura := db_connection.Factura{
		NCF_id:           ncf_id,
		CLI_id:           cli_id,
		TPO_id:           tpo_id,
		Cliente:          cliente,
		CostoSubtotal:    costoSubtotal,
		CostoTotal:       costoTotal,
		Descuento:        descuento,
		Envio:            envio,
		EnDolares:        enDolares,
		Descripcion:      descripcion,
		FechaVencimiento: fehcaVencimiento,
	}

	err := dataValidation(&ncf_id, &cli_id, &tpo_id, &costoSubtotal, &costoTotal, &cliente, &descripcion, &fehcaVencimiento)

	if err != nil {
		return factura, err
	}

	result := db_connection.Db.Create(&factura)

	if result.Error != nil {
		return factura, result.Error
	}

	return factura, nil
}

func GetAll() ([]FacturaDTO, error) {
	facturas := make([]FacturaDTO, 0)

	result := db_connection.Db.Model(&db_connection.Factura{}).
		Select("facturas.id, " +
			"facturas.cliente, " +
			"facturas.costo_subtotal, " +
			"facturas.costo_total, " +
			"facturas.descuento, " +
			"facturas.envio, " +
			"facturas.descripcion," +
			"facturas.en_dolares," +
			"facturas.created_at," +
			"facturas.fecha_vencimiento," +
			"ncfs.serie NCFSerie," +
			"ncfs.tipo NCFTipo," +
			"ncfs.secuencia NCFSecuencia," +
			"tipo_pagos.descripcion TipoPago").
		Joins("left join ncfs on ncfs.id = facturas.ncf_id").
		Joins("left join tipo_pagos on tipo_pagos.id = facturas.tpo_id").
		Scan(&facturas)

	if result.Error != nil {
		return facturas, result.Error
	}

	return facturas, nil
}

func GetById(id uint) (FacturaDTO, error) {
	factura := new(FacturaDTO)

	result := db_connection.Db.Model(&db_connection.Factura{}).
		Select("facturas.id, "+
			"facturas.cliente, "+
			"facturas.costo_subtotal, "+
			"facturas.costo_total, "+
			"facturas.descuento, "+
			"facturas.envio, "+
			"facturas.descripcion,"+
			"facturas.en_dolares,"+
			"facturas.created_at,"+
			"facturas.fecha_vencimiento,"+
			"ncfs.serie NCFSerie,"+
			"ncfs.tipo NCFTipo,"+
			"ncfs.secuencia NCFSecuencia,"+
			"tipo_pagos.descripcion TipoPago").
		Joins("left join ncfs on ncfs.id = facturas.ncf_id").
		Joins("left join tipo_pagos on tipo_pagos.id = facturas.tpo_id").
		Where("facturas.id = ?", id).Scan(factura)

	if result.Error != nil {
		return *factura, result.Error
	}

	return *factura, nil
}

func Update(ncf_id, cli_id, tpo_id, costoSubtotal, costoTotal, descuento, envio, id uint, cliente, descripcion string, enDolares bool, fechaVencimiento time.Time) error {
	factura := new(db_connection.Factura)

	result := db_connection.Db.Find(factura, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun factura")
	}

	err := dataValidation(&ncf_id, &cli_id, &tpo_id, &costoSubtotal, &costoTotal, &cliente, &descripcion, &fechaVencimiento)
	if err != nil {
		return err
	}

	factura.NCF_id = ncf_id
	factura.CLI_id = cli_id
	factura.TPO_id = tpo_id
	factura.CostoSubtotal = costoSubtotal
	factura.CostoTotal = costoTotal
	factura.Descripcion = descripcion
	factura.Envio = envio
	factura.Descuento = descuento
	factura.Cliente = cliente
	factura.EnDolares = enDolares
	factura.FechaVencimiento = fechaVencimiento

	result = db_connection.Db.Save(factura)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id uint) error {
	factura := new(db_connection.Factura)

	result := db_connection.Db.Find(factura, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun factura")
	}

	result = db_connection.Db.Delete(factura)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func dataValidation(ncf_id, cli_id, tpo_id, costoSubtotal, costoTotal *uint, cliente, descripcion *string, fechaVencimiento *time.Time) error {

	if *ncf_id == 0 {
		return errors.New("el ncf no puede ser cero")
	}
	if *cli_id == 0 {
		return errors.New("el cliente no puede ser cero")
	}
	if *tpo_id == 0 {
		return errors.New("el tipo de pago no puede ser cero")
	}
	if *costoSubtotal == 0 {
		return errors.New("costoSubtotal no puede ser cero")
	}
	if *costoTotal == 0 {
		return errors.New("costoTotal no puede ser cero")
	}
	if *cliente == "" {
		return errors.New("cliente no puede estar vacío")
	}
	if len(*cliente) > 150 {
		return errors.New("cliente no puede exceder 150 caracteres")
	}
	if *descripcion == "" {
		return errors.New("descripcion no puede estar vacía")
	}
	if len(*descripcion) > 150 {
		return errors.New("descripcion no puede exceder 150 caracteres")
	}
	return nil
}
