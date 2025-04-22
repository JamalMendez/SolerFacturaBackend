package cotizacion

import (
	"errors"
	"time"

	"ggstudios/solerfacturabackend/db_connection"
)

type CotizacionDTO struct {
	ID               uint
	Serie            string
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

type CotizacionDTOSend struct {
	ID            uint
	NCF           uint
	TipoPago      uint
	Cliente       uint
	CostoSubtotal uint
	CostoTotal    uint
	Descuento     uint
	Envio         uint
	Descripcion   string
	EnDolares     bool
}

func Create(cli_id, tpo_id, costoSubtotal, costoTotal, descuento, envio uint, secuencia, cliente, descripcion string, enDolares bool) (db_connection.Cotizacion, error) {
	cotizacion := db_connection.Cotizacion{
		Secuencia:     secuencia,
		CLI_id:        cli_id,
		TPO_id:        tpo_id,
		Cliente:       cliente,
		CostoSubtotal: costoSubtotal,
		CostoTotal:    costoTotal,
		Descuento:     descuento,
		Envio:         envio,
		EnDolares:     enDolares,
		Descripcion:   descripcion,
	}

	err := dataValidation(&cli_id, &tpo_id, &costoSubtotal, &costoTotal, &secuencia, &cliente, &descripcion)
	if err != nil {
		return cotizacion, err
	}

	result := db_connection.Db.Create(&cotizacion)

	if result.Error != nil {
		return cotizacion, result.Error
	}

	return cotizacion, nil
}

func GetAll() ([]CotizacionDTO, error) {
	cotizacions := make([]CotizacionDTO, 0)

	result := db_connection.Db.Model(&db_connection.Cotizacion{}).
		Select("cotizacions.id, " +
			"cotizacions.cliente, " +
			"cotizacions.costo_subtotal, " +
			"cotizacions.costo_total, " +
			"cotizacions.descuento, " +
			"cotizacions.envio, " +
			"cotizacions.descripcion," +
			"cotizacions.en_dolares," +
			"cotizacions.secuencia," +
			"cotizacions.created_at," +
			"cotizacions.fecha_vencimiento," +
			"tipo_pagos.descripcion TipoPago").
		Joins("left join tipo_pagos on tipo_pagos.id = cotizacions.tpo_id").
		Scan(&cotizacions)

	if result.Error != nil {
		return cotizacions, result.Error
	}

	return cotizacions, nil
}

func GetById(id uint) (CotizacionDTO, error) {
	cotizacion := new(CotizacionDTO)

	result := db_connection.Db.Model(&db_connection.Cotizacion{}).
		Select("cotizacions.id, "+
			"cotizacions.cliente, "+
			"cotizacions.costo_subtotal, "+
			"cotizacions.costo_total, "+
			"cotizacions.descuento, "+
			"cotizacions.envio, "+
			"cotizacions.descripcion,"+
			"cotizacions.en_dolares,"+
			"cotizacions.secuencia,"+
			"cotizacions.created_at,"+
			"cotizacions.fecha_vencimiento,"+
			"tipo_pagos.descripcion TipoPago").
		Joins("left join tipo_pagos on tipo_pagos.id = cotizacions.tpo_id").
		Where("cotizacions.id = ?", id).Scan(cotizacion)

	if result.Error != nil {
		return *cotizacion, result.Error
	}

	return *cotizacion, nil
}

func Update(cli_id, tpo_id, costoSubtotal, costoTotal, descuento, envio, id uint, secuencia, cliente, descripcion string, enDolares bool, fechaVencimiento time.Time) error {
	cotizacion := new(db_connection.Cotizacion)

	result := db_connection.Db.Find(cotizacion, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun cotizacion")
	}

	err := dataValidation(&cli_id, &tpo_id, &costoSubtotal, &costoTotal, &secuencia, &cliente, &descripcion)
	if err != nil {
		return err
	}

	cotizacion.Secuencia = secuencia
	cotizacion.CLI_id = cli_id
	cotizacion.TPO_id = tpo_id
	cotizacion.CostoSubtotal = costoSubtotal
	cotizacion.CostoTotal = costoTotal
	cotizacion.Descripcion = descripcion
	cotizacion.Envio = envio
	cotizacion.Descuento = descuento
	cotizacion.Cliente = cliente
	cotizacion.EnDolares = enDolares

	result = db_connection.Db.Save(cotizacion)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id uint) error {
	cotizacion := new(db_connection.Cotizacion)

	result := db_connection.Db.Find(cotizacion, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun cotizacion")
	}

	result = db_connection.Db.Delete(cotizacion)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func dataValidation(cli_id, tpo_id, costoSubtotal, costoTotal *uint, secuencia, cliente, descripcion *string) error {
	if *secuencia == "" {
		return errors.New("secuencia no puede estar vacío")
	}
	if len(*secuencia) > 8 {
		return errors.New("secuencia no puede exceder 8 caracteres")
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
