package cliente

import (
	"errors"
	"strings"

	"ggstudios/solerfacturabackend/db_connection"
)

type ClienteDTO struct {
	ID         uint
	RNC_Cedula string
	Nombre     string
	Apellido   string
	Email      string
	Direccion  string
	Ciudad     string
	Telefono   string
	Celular    string
}

func Create(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular string) (ClienteDTO, error) {
	err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)

	if err != nil {
		return ClienteDTO{}, err
	}

	cliente := db_connection.Cliente{
		RNC_Cedula: rnc_cedula,
		Nombre:     nombre,
		Apellido:   apellido,
		Email:      email,
		Direccion:  direccion,
		Ciudad:     ciudad,
		Telefono:   telefono,
		Celular:    celular,
	}

	result := db_connection.Db.Create(&cliente)

	if result.Error != nil {
		return ClienteDTO{}, result.Error
	}

	return ClienteDTO{
		ID:         cliente.ID,
		RNC_Cedula: cliente.RNC_Cedula,
		Nombre:     cliente.Nombre,
		Apellido:   cliente.Apellido,
		Email:      cliente.Email,
		Direccion:  cliente.Direccion,
		Ciudad:     cliente.Ciudad,
		Telefono:   cliente.Telefono,
		Celular:    cliente.Celular,
	}, nil
}

func GetAll() ([]ClienteDTO, error) {
	clientes := make([]ClienteDTO, 0)

	result := db_connection.Db.Model(&db_connection.Cliente{}).
		Select("id, rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular").
		Scan(&clientes)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no se encontro ningun cliente")
	}

	return clientes, nil
}

func GetById(id uint) (ClienteDTO, error) {
	cliente := new(ClienteDTO)

	result := db_connection.Db.Model(&db_connection.Cliente{}).
		Select("id, rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular").
		Where("id = ?", id).
		Scan(cliente)

	if result.Error != nil {
		return ClienteDTO{}, result.Error
	}

	if result.RowsAffected == 0 {
		return ClienteDTO{}, errors.New("no se encontro ningun cliente")
	}

	return *cliente, nil
}

func Update(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular string, id uint) error {
	cliente := new(db_connection.Cliente)

	err := dataValidation(&rnc_cedula, &nombre, &apellido, &email, &direccion, &ciudad, &telefono, &celular)

	if err != nil {
		return err
	}

	result := db_connection.Db.Find(cliente, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	cliente.RNC_Cedula = rnc_cedula
	cliente.Nombre = nombre
	cliente.Apellido = apellido
	cliente.Email = email
	cliente.Direccion = direccion
	cliente.Ciudad = ciudad
	cliente.Telefono = telefono
	cliente.Celular = celular
	result = db_connection.Db.Save(cliente)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Delete(id uint) error {
	cliente := new(db_connection.Cliente)

	result := db_connection.Db.Find(cliente, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no se encontro ningun producto")
	}

	result = db_connection.Db.Delete(cliente)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func dataValidation(rnc_cedula, nombre, apellido, email, direccion, ciudad, telefono, celular *string) error {
	isDotAtTheEdge := strings.HasPrefix(*email, ".") || strings.HasSuffix(*email, ".")
	isArrAtTheEdge := strings.HasPrefix(*email, "@") || strings.HasSuffix(*email, "@")
	emailContainArrDot := strings.Contains(*email, "@") && strings.Contains(*email, ".")
	isEmailInvalid := !emailContainArrDot || isArrAtTheEdge || isDotAtTheEdge

	switch {
	case len(*rnc_cedula) < 8 || len(*rnc_cedula) > 11:
		return errors.New("cedula o rnc no es valido")
	case len(*nombre) == 0 || len(*nombre) > 100:
		return errors.New("el nombre no es valido")
	case len(*apellido) == 0 || len(*apellido) > 100:
		return errors.New("el apellido no es valido")
	case len(*email) > 0 && isEmailInvalid:
		return errors.New("el email no es valido")
	case len(*email) > 150:
		return errors.New("el email excede el tamaño permitido")
	case len(*direccion) > 200:
		return errors.New("la direccion excede el tamaño permitido")
	case len(*ciudad) > 200:
		return errors.New("la ciudad excede el tamaño permitido")
	case len(*telefono) > 0 && len(*telefono) != 10:
		return errors.New("el telefono no es valido")
	case len(*celular) > 0 && len(*celular) != 10:
		return errors.New("el celular no es valido")
	}

	return nil
}
