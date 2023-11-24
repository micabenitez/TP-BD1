package funciones

import (
	db "alpuin-benitez-chianese-zanardi-tp/tablas" //Tablas con su estructura definida
)

func DropTables() {
	db.ClienteDropTable()
	db.TarjetaDropTable()
	db.ComercioDropTable()
	db.CompraDropTable()
	db.RechazoDropTable()
	db.CierreDropTable()
	db.CabeceraDropTable()
	db.DetalleDropTable()
	db.AlertaDropTable()
	db.ConsumoDropTable()
}
