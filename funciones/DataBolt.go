package funciones

import (
	db "alpuin-benitez-chianese-zanardi-tp/tablas" //Tablas con su estructura definida
	"fmt"
	"time"
)

func InsertDataBolt() {
	ClienteInsertDataBolt()
	TarjetaInsertDataBolt()
	ComercioInsertDataBolt()
	CompraInsertDataBolt()
}

func SelectDataBolt() {
	db.ClienteSelectAll()
	db.TarjetaSelectAll()
	db.ComercioSelectAll()
	db.CompraSelectAll()
}

func ClienteInsertDataBolt() {
	/*
		estrucutra sql : nrocliente int, nombre text, apellido text, domicilio text ,telefono char(12)
	*/
	db.ClienteInsertBolt(1, "Matias", "Alpuin", "Cavallari 1235", "1533445566")
	db.ClienteInsertBolt(2, "Micaela", "Benitez", "San Martin 1926", "1127545213")
	db.ClienteInsertBolt(3, "Valentin", "Chianese", "Vacarezza 216", "1122469852")
	db.ClienteInsertBolt(4, "Joaquin", "Zanardi", "Mitre 1118", "1511329647")
}

func TarjetaInsertDataBolt() {
	/*
		estrucutra sql : nrotarjeta char(16), nrocliente int, validadesde char(6), validahasta char(6), codseguridad char(4),limitecompra decimal(8,2),estado char(10)
		estados :
			vigente
			anulada
			suspendida
	*/
	db.TarjetaInsertBolt("7586578112433125", 1, "201012", "203012", "319", 125000, "vigente")
	db.TarjetaInsertBolt("4259185660563593", 2, "201907", "202307", "536", 75000, "vigente")
	db.TarjetaInsertBolt("4671996838874773", 3, "202004", "202404", "687", 75000, "suspendida")
	db.TarjetaInsertBolt("4522632177487734", 4, "201712", "202112", "240", 75000, "anulada")

}

func ComercioInsertDataBolt() {
	/*
		estrucutra sql : nrocomercio int, nombre text, domicilio text, codigopostal char(8), telefono char(12)
	*/
	db.ComercioInsertBolt(1, "Vitalcer", "Alsina 298", "1744", "2914742012")
	db.ComercioInsertBolt(2, "Walmart", "Av T Cosentino 617 ", "1744", "1125961378")
	db.ComercioInsertBolt(3, "Panaderia Danesa", "Thompson 315", "1664", "2914510333")
	db.ComercioInsertBolt(4, "Supermercado Dia", "Belgrano 224", "1751", "4888626")
}

func CompraInsertDataBolt() {
	db.CompraInsertBolt(1, "7586578112433125", 1, fmt.Sprintf("%s", time.Now()), 1252.0, true)
	db.CompraInsertBolt(2, "4259185660563593", 2, fmt.Sprintf("%s", time.Now()), 4157.1, true)
	db.CompraInsertBolt(3, "4671996838874773", 3, fmt.Sprintf("%s", time.Now()), 13012, true)
	db.CompraInsertBolt(4, "4522632177487734", 4, fmt.Sprintf("%s", time.Now()), 51642.7, true)

}
