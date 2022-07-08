package models

import (
	"gorm.io/gorm"
	"time"
)

type TokenLogin struct {
	Token string `json:"token"`
}

type Inmuebles struct {
	Inmuebles []*Inmueble `json:"inmuebles"`
}

type Inmueble struct {
	gorm.Model
	Direccion     string  `json:"direccion"`
	Ambientes     int     `json:"ambientes"`
	Tipo          string  `json:"tipo"`
	Uso           string  `json:"uso"`
	Precio        float64 `json:"precio"`
	Disponible    bool    `json:"disponible"`
	PropietarioID int64   `json:"propietario"`
}

type Propietario struct {
	gorm.Model
	Dni      int64  `json:"dni"`
	Apellido string `json:"apellido"`
	Nombre   string `json:"nombre"`
	Telefono int64  `json:"telefono"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type Pago struct {
	gorm.Model
	NumeroPago int       `json:"numero_pago"`
	AlquilerID int       `json:"alquiler"`
	Fecha      time.Time `json:"fecha"`
	Importe    float64   `json:"importe"`
}

type Alquiler struct {
	gorm.Model
	InquilinoId string    `json:"inquilino"`
	InmuebleID  int       `json:"inmueble"`
	FechaInicio time.Time `json:"fechaInicio"`
	FechaFin    time.Time `json:"fechaFin"`
	Precio      float64   `json:"precio"`
}

type Inquilino struct {
	gorm.Model
	Dni       int64  `json:"dni"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	Telefono  int64  `json:"telefono"`
	Direccion string `json:"direccion"`
}
