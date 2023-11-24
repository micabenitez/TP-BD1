package funciones

import (
	db "alpuin-benitez-chianese-zanardi-tp/tablas" //Tablas con su estructura definida
	"fmt"
	"time"
)

// funcion que hay que llamar desde el main
func InsertData() {
	fmt.Println("* Agregando Datos *")
	CierreInsertData()
	ClienteInsertData()
	TarjetaInsertData()
	ComercioInsertData()
	ConsumoInsertData()
}

/// Se deben cargar los datos en las tablas
// Fecha sin timestamp como string -> ejemplo : "2022/12/31"
// Fecha

func CierreInsertData() {
	/*
		estrucutra sql : anio int, mes int, terminacion int, fechainicio date, fechacierre date, fechavto date
	*/
	for anio := 2021; anio <= 2022; anio++ { //anios
		fecha := time.Date(anio, 1, 5, 0, 0, 0, 0, time.UTC)    //dia de inicio = 5, mes de inicio 1
		for terminacion := 0; terminacion < 10; terminacion++ { //terminacion
			for mes := 1; mes <= 12; mes++ {
				fechaInicio := fmt.Sprint(fecha.Format("2006/01/02"))
				fechaCierre := fmt.Sprint(fecha.AddDate(0, 0, 30).Format("2006/01/02"))      //30 dias luego del inicio
				fechaVencimiento := fmt.Sprint(fecha.AddDate(0, 0, 37).Format("2006/01/02")) //7 dias luego del cierre
				db.CierreInsert(anio, mes, terminacion, fechaInicio, fechaCierre, fechaVencimiento)
				//fmt.Printf("inicio %s : cierre %s : vto %s \n", fechaInicio, fechaCierre, fechaVencimiento)
				fecha = fecha.AddDate(0, 1, 0) //suma 1 al mes
			}
			fecha = fecha.AddDate(0, 0, 1) //suma 1 al dia
		}
	}
}

func ClienteInsertData() {
	/*
		estrucutra sql : nrocliente int, nombre text, apellido text, domicilio text ,telefono char(12)
	*/
	db.ClienteInsert(1, "Matias", "Alpuin", "Cavallari 1235", "1533445566")
	db.ClienteInsert(2, "Micaela", "Benitez", "San Martin 1926", "1127545213")
	db.ClienteInsert(3, "Valentin", "Chianese", "Vacarezza 216", "1122469852")
	db.ClienteInsert(4, "Joaquin", "Zanardi", "Mitre 1118", "1511329647")
	db.ClienteInsert(5, "Carlos", "Martinez", "San Martin 2146", "1189654713")
	db.ClienteInsert(6, "Carmen", "Moreno", "Bartolome Mitre 5146", "1545823696")
	db.ClienteInsert(7, "Angel", "Jimenez", "Ceballos 172", "1178992111")
	db.ClienteInsert(8, "Javier", "Perez", "Gral Paz 111", "1133674899")
	db.ClienteInsert(9, "Dolores", "Sanchez", "Cervantes 1583", "1179452155")
	db.ClienteInsert(10, "Ana", "Ruiz", "Alcorta 501", "1512457988")
	db.ClienteInsert(11, "Ian", "Muñoz", "Gorriti 1216", "1199271233")
	db.ClienteInsert(12, "Pilar", "Santoro", "Rivadavia 16275", "1177652144")
	db.ClienteInsert(13, "Milagros", "Rinaldi", "Av Balbin 1727", "1187224777")
	db.ClienteInsert(14, "Luis", "Coppola", "La Plata 4022", "1146572114")
	db.ClienteInsert(15, "Laura", "Martini", "Hipolito Yrigoyen 543", "1544786655")
	db.ClienteInsert(16, "Martina", "Ricci", "25 De Mayo 302", "1521135688")
	db.ClienteInsert(17, "Daniel", "Amato", "Mitre 4211", "1199784521")
	db.ClienteInsert(18, "Pablo", "Vitale", "Parana 6370", "1155228237")
	db.ClienteInsert(19, "Raquel", "Bianchi", "Laprida 4475", "1132374556")
	db.ClienteInsert(20, "Irene", "Lombardi", "Avellaneda 3348", "1185465521")
}

func ComercioInsertData() {
	/*
		estrucutra sql : nrocomercio int, nombre text, domicilio text, codigopostal char(8), telefono char(12)
	*/
	db.ComercioInsert(1, "Vitalcer", "Alsina 298", "1744", "2914742012")
	db.ComercioInsert(2, "Walmart", "Av T Cosentino 617 ", "1744", "1125961378")
	db.ComercioInsert(3, "Panaderia Danesa", "Thompson 315", "1664", "2914510333")
	db.ComercioInsert(4, "Supermercado Dia", "Belgrano 224", "1751", "4888626")
	db.ComercioInsert(5, "Pet Shop", "Undiano 630", "1824", "1145668279")
	db.ComercioInsert(6, "Club Digital", "Av colon 256", "8109", "4851111")
	db.ComercioInsert(7, "Lavanderia Laverap", "Belgrano 318", "1002", "1185132298")
	db.ComercioInsert(8, "Cerrajeria Master Lock", "Inglaterra 1750", "1752", "2914232375")
	db.ComercioInsert(9, "McDonalds", "Sarmiento 727", "1004", "2915661297")
	db.ComercioInsert(10, "Quimica Cromax", "Ricardo Balbin 5512", "1650", "1127134270")
	db.ComercioInsert(11, "Cerveceria Ocaso", "Saavedra 19", "1044", "1135678811")
	db.ComercioInsert(12, "Grana Pasteleria", "Zufriategui 3698", "1603", "1522888134")
	db.ComercioInsert(13, "COTO", "Albeniz 3099", "1661", "1146682636")
	db.ComercioInsert(14, "Cotillon Cotyland", "Parana 6552", "1607", "1124981710")
	db.ComercioInsert(15, "Adidas Store", "Juan Manuel de Rosas 658", "1708", "03327424261")
	db.ComercioInsert(16, "Yolive Perfumeria", "Triunvirato 1555", "1611", "1147410076")
	db.ComercioInsert(17, "Heladeria Bambola", "Av Santa fe 848", "1641", "1147984943")
	db.ComercioInsert(18, "Farmacity", "Las Heras 2055", "1127", "1148090273")
	db.ComercioInsert(19, "Burger King", "Av Presidente Perón 1622", "1663", "1144584689")
	db.ComercioInsert(20, "Ikea Textil", "Lavalle 2562", "1052", "2915661297")
}

func ConsumoInsertData() {
	/*
		estrucutra sql : nrotarjeta char(16), codseguridad char(4), nrocomercio int, monto decimal(7,2)
	*/
	db.ConsumoInsert("2020202020202020", "123", 1, 930.13)
	db.ConsumoInsert("4909957014010164", "205", 20, 5000)
	db.ConsumoInsert("4002681920219661", "597", 11, 2500)

	db.ConsumoInsert("4221979556273084", "785", 1, 2567.50) // dos compras en comercios con cp iguales
	db.ConsumoInsert("4221979556273084", "785", 2, 5687.90)

	db.ConsumoInsert("2020202020202020", "345", 2, 1250.13) // cod de seguridad distinto

	db.ConsumoInsert("4259185660563593", "536", 15, 31248.13) // dos compras en distintos comercios con cp distintos
	db.ConsumoInsert("4259185660563593", "536", 4, 4334.13)

	db.ConsumoInsert("4765352851617417", "487", 6, 50000) // excede limite de tarjeta en el mismo dia
	db.ConsumoInsert("4765352851617417", "487", 15, 26000)

	db.ConsumoInsert("4962429668094682", "114", 7, 1312.13) // <- tarjeta suspendida
	db.ConsumoInsert("4840943520768117", "950", 8, 929.13)  // <- tarjeta vencida
	db.ConsumoInsert("4522632177487734", "240", 10, 950.15) // <- tarjeta anulada
}

func TarjetaInsertData() {
	/*
		estrucutra sql : nrotarjeta char(16), nrocliente int, validadesde char(6), validahasta char(6), codseguridad char(4),limitecompra decimal(8,2),estado char(10)
		estados :
			vigente
			anulada
			suspendida
	*/
	db.TarjetaInsert("2020202020202020", 1, "201012", "202012", "123", 125000, "vigente")
	db.TarjetaInsert("7586578112433125", 1, "201012", "203012", "319", 125000, "vigente")
	db.TarjetaInsert("4259185660563593", 2, "201907", "202307", "536", 75000, "vigente")
	db.TarjetaInsert("4671996838874773", 3, "202004", "202404", "687", 75000, "suspendida")
	db.TarjetaInsert("4522632177487734", 4, "201712", "202112", "240", 75000, "anulada")
	db.TarjetaInsert("4103636905206677", 5, "202006", "202406", "504", 75000, "vigente")
	db.TarjetaInsert("4765082003981193", 6, "201906", "202306", "675", 75000, "vigente")
	db.TarjetaInsert("4765352851617417", 7, "201808", "202208", "487", 75000, "vigente")
	db.TarjetaInsert("4962429668094682", 8, "202107", "202507", "114", 75000, "suspendida")
	db.TarjetaInsert("4626392065351394", 9, "202007", "202307", "868", 75000, "vigente")
	db.TarjetaInsert("4513252154152394", 10, "201805", "202205", "868", 75000, "vigente")
	db.TarjetaInsert("4504601299467926", 11, "201605", "202105", "651", 60000, "vigente")
	db.TarjetaInsert("4805069202283143", 11, "201711", "202211", "993", 100000, "vigente")
	db.TarjetaInsert("4840943520768117", 12, "201705", "202204", "950", 70000, "vigente") // <- vencida
	db.TarjetaInsert("4944359396514524", 13, "201905", "202303", "257", 100000, "vigente")
	db.TarjetaInsert("4839952707405503", 14, "201905", "202307", "703", 80000, "vigente")
	db.TarjetaInsert("4221979556273084", 15, "201705", "202210", "785", 90000, "vigente")
	db.TarjetaInsert("4851957935938210", 16, "201705", "202209", "219", 95000, "vigente")
	db.TarjetaInsert("4783625183391266", 17, "202105", "202606", "913", 40000, "vigente")
	db.TarjetaInsert("4909957014010164", 18, "202105", "202604", "205", 50000, "vigente")
	db.TarjetaInsert("4326416285050792", 19, "201806", "202306", "830", 100000, "vigente")
	db.TarjetaInsert("4002681920219661", 20, "201905", "202405", "597", 75000, "vigente")

}
