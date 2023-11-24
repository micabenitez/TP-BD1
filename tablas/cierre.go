package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"fmt"
)

func CierreCreateTable() {
	db.CreateTable("cierre", "anio int, mes int, terminacion int, fechainicio date, fechacierre date, fechavto date")
}

func CierreCreatePK() {
	db.CreatePK("cierre", "cierre_pk", "anio,mes,terminacion")
}

func CierreDropTable() {
	db.DropTable("alerta")
}

func CierreDeletePK() {
	db.DeleteKey("cierre", "cierre_pk")
}

func CierreInsert(anio int, mes int, terminacion int, fechainicio string, fechacierre string, fechavto string) {
	values := fmt.Sprintf("%d,%d,%d,'%s','%s','%s'", anio, mes, terminacion, fechainicio, fechacierre, fechavto)
	db.Insert("cierre", values)
}
