package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"fmt"
	"strings"
)

func RechazoCreateTable() {
	db.CreateTable("rechazo", "nrorechazo int, nrotarjeta char(16), nrocomercio int, fecha timestamp, monto decimal(7,2), motivo text")
}

func RechazoCreatePK() {
	db.CreatePK("rechazo", "rechazo_pk", "nrorechazo")
}

func RechazoCreateFK() {
	db.CreateFK("rechazo", "rechazo_tarjeta_fk", "nrotarjeta", "tarjeta", "nrotarjeta")
	db.CreateFK("rechazo", "rechazo_comercio_fk", "nrocomercio", "comercio", "nrocomercio")
}

func RechazoDropTable() {
	db.DropTable("alerta")
}

func RechazoDeletePK() {
	db.DeleteKey("rechazo", "rechazo_pk")
}

func RechazoDeleteFK() {
	db.DeleteKey("rechazo", "rechazo_tarjeta_fk")
	db.DeleteKey("rechazo", "rechazo_comercio_fk")
}

func RechazoInsert(nrorechazo int, nrotarjeta string, nrocomercio int, fechaTimeStamp string, monto float64, motivo string) {
	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}
	values := fmt.Sprintf("%d,'%s',%d,'%s',%f,'%s'", nrorechazo, nrotarjeta, nrocomercio, fechaTimeStamp, monto, motivo)
	db.Insert("rechazo", values)
}
