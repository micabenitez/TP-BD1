package funciones

import (
	db "alpuin-benitez-chianese-zanardi-tp/tablas" //Tablas con su estructura definida
)

func CreateTables() {
	db.ClienteCreateTable()
	db.TarjetaCreateTable()
	db.ComercioCreateTable()
	db.CompraCreateTable()
	db.RechazoCreateTable()
	db.CierreCreateTable()
	db.CabeceraCreateTable()
	db.DetalleCreateTable()
	db.AlertaCreateTable()
	db.ConsumoCreateTable()
}
