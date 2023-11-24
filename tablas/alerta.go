package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib" //conexion y funcionalidad basica de sql
	_ "fmt"
	_ "strings"
)

func AlertaCreateTable() {
	db.CreateTable("alerta", "nroalerta int, nrotarjeta char(16), fecha timestamp, nrorechazo int, codalerta int, descripcion text")
}

func AlertaCreatePK() {
	db.CreatePK("alerta", "alerta_pk", "nroalerta")
}

func AlertaCreateFK() {
	db.CreateFK("alerta", "alerta_tarjeta_fk", "nrotarjeta", "tarjeta", "nrotarjeta")
	db.CreateFK("alerta", "alerta_rechazo_fk", "nrorechazo", "rechazo", "nrorechazo")
}

func AlertaDropTable() {
	db.DropTable("alerta")
}

func AlertaDeletePK() {
	db.DeleteKey("alerta", "alerta_pk")
}

func AlertaDeleteFK() {
	db.DeleteKey("alerta", "alerta_tarjeta_fk")
	db.DeleteKey("alerta", "alerta_rechazo_fk")
}

/**
func AlertaInsert(nroalerta int, nrotarjeta string, fechaTimeStamp string, nrorechazo int, codalerta int, descripcion string) {
	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}
	values := fmt.Sprintf("%d,'%s','%s',%d,%d,'%s'", nroalerta, nrotarjeta, fechaTimeStamp, nrorechazo, codalerta, descripcion)
	db.Insert("alerta", values)
}
*/
