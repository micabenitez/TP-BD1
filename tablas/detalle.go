package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"fmt"
)

func DetalleCreateTable() {
	db.CreateTable("detalle", "nroresumen int, nrolinea int, fecha date, nombrecomercio text, monto decimal(7,2)")
}

func DetalleCreatePK() {
	db.CreatePK("detalle", "detalle_pk", "nroresumen,nrolinea")
}

func DetalleCreateFK() {
	db.CreateFK("detalle", "detalle_cabecera_fk", "nroresumen", "cabecera", "nroresumen")
}

func DetalleDeletePK() {
	db.DeleteKey("detalle", "detalle_pk")
}

func DetalleDeleteFK() {
	db.DeleteKey("detalle", "detalle_cabecera_fk")
}

func DetalleDropTable() {
	db.DropTable("alerta")
}

func DetalleInsert(nroresumen int, nrolinea int, fecha string, nombrecomercio string, monto float64) {
	values := fmt.Sprintf("%d,%d,'%s','%s',%f", nroresumen, nrolinea, fecha, nombrecomercio, monto)
	db.Insert("detalle", values)
}
