package models

import (
	"gorm.io/gorm"
	"time"
)

type Inmueble struct {
	gorm.Model
	InmuebleID    int64   `json:"inmuebleID"`
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
	PropietarioID int64  `json:"propietarioID"`
	Dni           int64  `json:"dni"`
	Apellido      string `json:"apellido"`
	Nombre        string `json:"nombre"`
	Telefono      int64  `json:"telefono"`
	Mail          string `json:"mail"`
	Password      bool   `json:"password"`
}

type Pagos struct {
	gorm.Model
	PagoID     int64     `json:"inmuebleID"`
	NumeroPago int       `json:"numero_pago"`
	AlquilerID int       `json:"alquiler"`
	Fecha      time.Time `json:"fecha"`
	Importe    float64   `json:"importe"`
}

type Alquiler struct {
	gorm.Model
	AlquilerID  int64     `json:"alquilerID"`
	InquilinoId string    `json:"inquilino"`
	InmuebleID  int       `json:"inmueble"`
	FechaInicio time.Time `json:"fechaInicio"`
	FechaFin    time.Time `json:"fechaFin"`
	Precio      float64   `json:"precio"`
}

type Inquilino struct {
	gorm.Model
	InquilinoID int64  `json:"inquilinoID"`
	Dni         int64  `json:"dni"`
	Apellido    string `json:"apellido"`
	Nombre      string `json:"nombre"`
	Telefono    int64  `json:"telefono"`
	Direccion   string `json:"direccion"`
}
