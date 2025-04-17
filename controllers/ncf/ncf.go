package ncf

import (
	"errors"
	"strconv"
	"strings"

	"ggstudios/solerfacturabackend/db_connection"
)

type NCFDTO struct {
	ID        uint
	Serie     string
	Tipo      string
	Secuencia string
}

func Create(serie, tipo, secuencia string) (NCFDTO, error) {
	if err := dataValidation(&serie, &tipo, &secuencia); err != nil {
		return NCFDTO{}, err
	}

	ncf := db_connection.NCF{
		Serie:     serie,
		Tipo:      tipo,
		Secuencia: secuencia,
	}

	result := db_connection.Db.Create(&ncf)

	if result.Error != nil {
		return NCFDTO{}, result.Error
	}

	return NCFDTO{
		ID:        ncf.ID,
		Serie:     ncf.Serie,
		Tipo:      ncf.Tipo,
		Secuencia: ncf.Secuencia,
	}, nil
}

func GetAll() ([]NCFDTO, error) {
	ncfs := make([]NCFDTO, 0)

	result := db_connection.Db.Model(&db_connection.NCF{}).
		Select("id, serie, tipo, secuencia").
		Scan(&ncfs)

	if result.Error != nil {
		return nil, result.Error
	}

	return ncfs, nil
}

func GetById(id uint) (NCFDTO, error) {
	ncf := new(NCFDTO)

	result := db_connection.Db.Model(&db_connection.NCF{}).
		Select("id, serie, tipo, secuencia").
		Where("id = ?", id).
		Scan(ncf)

	if result.Error != nil {
		return NCFDTO{}, result.Error
	}

	return *ncf, nil
}

func Update(serie, tipo, secuencia string, id uint) error {
	ncf := new(db_connection.NCF)

	if err := dataValidation(&serie, &tipo, &secuencia); err != nil {
		return err
	}

	result := db_connection.Db.Find(ncf, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	ncf.Serie = serie
	ncf.Tipo = tipo
	ncf.Secuencia = secuencia
	result = db_connection.Db.Save(ncf)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id uint) error {
	ncf := new(db_connection.NCF)

	result := db_connection.Db.Find(ncf, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	result = db_connection.Db.Delete(ncf)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func dataValidation(serie, tipo, secuencia *string) error {
	*serie = strings.ToUpper(*serie)

	if len(*serie) != 1 {
		return errors.New("solo se pueden ingresar 1 digito en la serie")
	}

	if _, err := strconv.Atoi(*tipo); err != nil {
		return errors.New("el tipo debe contener solo números")
	}

	if _, err := strconv.Atoi(*secuencia); err != nil {
		return errors.New("la secuencia debe contener solo números")
	}

	if len(*tipo) != 2 {
		return errors.New("solo se pueden ingresar 2 digitos en el tipo")
	}

	if len(*secuencia) != 8 {
		return errors.New("solo se pueden ingresar 8 digitos en la secuencia")
	}

	return nil
}
