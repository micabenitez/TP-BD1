package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func CompraCreateTable() {
	db.CreateTable("compra", "nrooperacion int, nrotarjeta char(16), nrocomercio int, fecha timestamp, monto decimal(7,2), pagado boolean")
}

func CompraCreatePK() {
	db.CreatePK("compra", "compra_pk", "nrooperacion")
}

func CompraCreateFK() {
	db.CreateFK("compra", "compra_tarjeta_fk", "nrotarjeta", "tarjeta", "nrotarjeta")
	db.CreateFK("compra", "compra_comercio_fk", "nrocomercio", "comercio", "nrocomercio")
}

func CompraDeletePK() {
	db.DeleteKey("compra", "compra_pk")
}

func CompraDeleteFK() {
	db.DeleteKey("compra", "compra_tarjeta_fk")
	db.DeleteKey("compra", "compra_comercio_fk")
}

func CompraDropTable() {
	db.DropTable("alerta")
}

func CompraInsert(nrooperacion int, nrotarjeta string, nrocomercio int, fechaTimeStamp string, monto float64, pagado bool) {
	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}
	values := fmt.Sprintf("%d,'%s',%d,'%s',%f,%t", nrocomercio, nrotarjeta, nrocomercio, fechaTimeStamp, monto, pagado)
	db.Insert("compra", values)
}

type Compra struct {
	Nrooperacion   int
	Nrotarjeta     string
	Nrocomercio    int
	FechaTimeStamp string
	Monto          float64
	Pagado         bool
}

func CompraInsertBolt(nrooperacion int, nrotarjeta string, nrocomercio int, fechaTimeStamp string, monto float64, pagado bool) {
	bucket := "compra"
	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}

	compras := Compra{Nrooperacion: nrooperacion,
		Nrotarjeta:     nrotarjeta,
		Nrocomercio:    nrocomercio,
		FechaTimeStamp: fechaTimeStamp,
		Monto:          monto,
		Pagado:         pagado}
	data, err := json.MarshalIndent(compras, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	db.WriteToBucket(bucket, []byte(strconv.Itoa(nrooperacion)), data)
}

func CompraSelectBolt(id int) {
	bucket := "compra"
	data, _ := db.ReadUniqueFromBucket(bucket, []byte(strconv.Itoa(id)))

	var compra Compra

	err := json.Unmarshal(data, &compra)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n", compra)
}

func CompraSelectAll() {
	db.SelectAllFromBucket("compra")
}
