package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"fmt"
	"strings"
)

func CabeceraCreateTable() {
	db.CreateTable("cabecera", "nroresumen int, nombre text, apellido text, domicilio text, nrotarjeta char(16), desde date, hasta date, vence date, total decimal(8,2)")
}

func CabeceraCreatePK() {
	db.CreatePK("cabecera", "cabecera_pk", "nroresumen")
}

func CabeceraCreateFK() {
	db.CreateFK("cabecera", "cabecera_tarjeta_fk", "nrotarjeta", "tarjeta", "nrotarjeta")
}

func CabeceraDropTable() {
	db.DropTable("alerta")
}

func CabeceraDeletePK() {
	db.DeleteKey("cabecera", "cabecera_pk")
}

func CabeceraDeleteFK() {
	db.DeleteKey("cabecera", "cabecera_tarjeta_fk")
}

func CabeceraInsert(nroresumen int, nombre string, apellido string, domicilio string, nrotarjeta string, desde string, hasta string, vence string, total float64) {
	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}
	values := fmt.Sprintf("%d,'%s','%s','%s','%s','%s','%s','%s',%f", nroresumen, nombre, apellido, domicilio, nrotarjeta, desde, hasta, vence, total)
	db.Insert("cabecera", values)
}
