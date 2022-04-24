package services

import (
	"InmoGo/src/api/config"
	"InmoGo/src/api/models"
	"database/sql"
	"fmt"
)

type Inmueble models.Inmueble

func (i *Inmueble) Save(inmueble Inmueble) error {
	if _, err := config.DB.Query("INSERT INTO inmueble VALUES( %1 )", Values(inmueble)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (i *Inmueble) Get(id string) (Inmueble, error) {
	data, err := config.DB.Query("SELECT * FROM inmueble WHERE inmuebleID = $1", id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return Scan(data)
}

func (i *Inmueble) GetAll(idPropietario string) ([]Inmueble, error) {
	data, err := config.DB.Query("SELECT * FROM inmueble WHERE propietario = $1", idPropietario)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return ScanAll(data)

}

func (i *Inmueble) Update(inmueble Inmueble) error {
	if _, err := config.DB.Query("UPDATE inmueble SET( $1 )", Values(inmueble)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Scan(data *sql.Rows) (Inmueble, error) {
	res := Inmueble{}
	if err := data.Scan(&res.InmuebleID, &res.Direccion, &res.Ambientes, &res.Tipo, &res.Uso, &res.Precio, &res.Disponible, &res.PropietarioID); err != nil {
		fmt.Println(err)
		return Inmueble{}, err
	}
	return res, nil
}

func ScanAll(data *sql.Rows) ([]Inmueble, error) {
	var inmuebles []Inmueble
	for data.Next() {
		inmueble, err := Scan(data)
		if err != nil {
			fmt.Println(err)
			return []Inmueble{}, err
		}
		inmuebles = append(inmuebles, inmueble)
	}

	return inmuebles, nil
}

func Values(inmueble Inmueble) string {
	return fmt.Sprintf("(%s,%s,%s,%s,%s,%s,%s", inmueble.Direccion, inmueble.Ambientes, inmueble.Tipo, inmueble.Uso, inmueble.Precio, inmueble.Disponible, inmueble.PropietarioID)
}
