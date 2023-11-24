package funciones

import (
	db "alpuin-benitez-chianese-zanardi-tp/tablas"
	"fmt"
)

func CreatePrimaryKey() {
	fmt.Println("* Creando Claves Primarias *")
	db.AlertaCreatePK()
	db.CabeceraCreatePK()
	db.CierreCreatePK()
	db.ClienteCreatePK()
	db.ComercioCreatePK()
	db.CompraCreatePK()
	db.DetalleCreatePK()
	db.RechazoCreatePK()
	db.TarjetaCreatePK()
}

func DeletePrimaryKey() {
	fmt.Println("* Borrando Claves Primarias *")
	db.AlertaDeletePK()
	db.CabeceraDeletePK()
	db.CierreDeletePK()
	db.ClienteDeletePK()
	db.ComercioDeletePK()
	db.CompraDeletePK()
	db.DetalleDeletePK()
	db.RechazoDeletePK()
	db.TarjetaDeletePK()
}
