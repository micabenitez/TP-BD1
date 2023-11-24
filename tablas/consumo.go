package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"fmt"
	"strings"
)

func ConsumoCreateTable() {
	db.CreateTable("consumo", "nrotarjeta char(16), codseguridad char(4), nrocomercio int, monto decimal(7,2)")
}

func ConsumoDropTable() {
	db.DropTable("alerta")
}

func ConsumoCreateFK() {
	db.CreateFK("consumo", "consumo_tarjeta_fk", "nrotarjeta", "tarjeta", "nrotarjeta")
	db.CreateFK("consumo", "consumo_comercio_fk", "nrocomercio", "comercio", "nrocomercio")
}

func ConsumoDeleteFK() {
	db.DeleteKey("consumo", "consumo_tarjeta_fk")
	db.DeleteKey("consumo", "consumo_comercio_fk")
}

func ConsumoInsert(nrotarjeta string, codseguridad string, nrocomercio int, monto float64) {
	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}
	if len(codseguridad) > 4 {
		nrotarjeta = nrotarjeta[0:3]
	}
	values := fmt.Sprintf("'%s','%s',%d,%f", nrotarjeta, codseguridad, nrocomercio, monto)
	db.Insert("consumo", values)
}
