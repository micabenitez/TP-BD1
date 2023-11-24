package funciones

import (
	db "alpuin-benitez-chianese-zanardi-tp/tablas"
	"fmt"
)

func CreateForeignKey() {
	//Cabecera + Detalle
	fmt.Println("* Creando Claves Externas *")
	db.AlertaCreateFK()
	db.CabeceraCreateFK()
	db.CompraCreateFK()
	db.DetalleCreateFK()
	db.RechazoCreateFK()
	db.TarjetaCreateFK()
	db.ConsumoCreateFK()
}

func DeleteForeignKey() {
	//alter table %s drop constraint %s
	fmt.Println("* Borrando Claves Externas *")
	db.AlertaDeleteFK()
	db.CabeceraDeleteFK()
	db.ComercioDeletePK()
	db.CompraDeleteFK()
	db.DetalleDeleteFK()
	db.RechazoDeleteFK()
	db.TarjetaDeleteFK()
	db.ConsumoDeleteFK()
}
