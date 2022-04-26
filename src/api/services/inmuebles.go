package services

import (
	"InmoGo/src/api/config"
	"InmoGo/src/api/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Inmueble models.Inmueble

func (i *Inmueble) Save(inmueble Inmueble) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	DB, err := sql.Open("mysql", "root:root@/InmoTricarico")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 10)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	err = DB.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return err
	}
	log.Printf("Connected to DB %s successfully\n", "Inmotricarico")
	if _, err := DB.Query("INSERT INTO inmueble" + inmueble.Columns() + " VALUE ( " + inmueble.Values(inmueble) + ")"); err != nil {
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
	if _, err := config.DB.Query("UPDATE inmueble SET( $1 )", inmueble.Values(inmueble)); err != nil {
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

func (i *Inmueble) Values(inmueble Inmueble) string {
	return fmt.Sprintf("'%s', %s, '%s', '%s', %s, %s, %s", inmueble.Direccion, strconv.Itoa(inmueble.Ambientes), inmueble.Tipo, inmueble.Uso, strconv.FormatFloat(inmueble.Precio, 'f', 2, 64), strconv.FormatBool(inmueble.Disponible), strconv.FormatInt(inmueble.PropietarioID, 10))
}

func (i *Inmueble) Columns() string {
	return "(direccion, ambientes, tipo, uso, precio, disponible, propietario)"
}
