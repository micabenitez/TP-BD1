package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func ComercioCreateTable() {
	db.CreateTable("comercio", "nrocomercio int, nombre text, domicilio text, codigopostal char(8), telefono char(12)")
}

func ComercioCreatePK() {
	db.CreatePK("comercio", "comercio_pk", "nrocomercio")
}

func ComercioDropTable() {
	db.DropTable("alerta")
}

func ComercioDeletePK() {
	db.DeleteKey("comercio", "comercio_pk")
}

func ComercioSelectAll() {
	db.SelectAllFromBucket("comercio")
}

func ComercioInsert(nrocomercio int, nombre string, domicilio string, codigopostal string, telefono string) {
	if len(codigopostal) > 8 {
		telefono = telefono[0:7]
	}
	if len(telefono) > 12 {
		telefono = telefono[0:11]
	}
	values := fmt.Sprintf("%d,'%s','%s','%s','%s'", nrocomercio, nombre, domicilio, codigopostal, telefono)
	db.Insert("comercio", values)
}

type Comercio struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

func ComercioInsertBolt(nrocomercio int, nombre string, domicilio string, codigopostal string, telefono string) {
	bucket := "comercio"
	if len(codigopostal) > 8 {
		telefono = telefono[0:7]
	}
	if len(telefono) > 12 {
		telefono = telefono[0:11]
	}
	comercios := Comercio{
		Nrocomercio:  nrocomercio,
		Nombre:       nombre,
		Domicilio:    domicilio,
		Codigopostal: codigopostal,
		Telefono:     telefono}

	data, err := json.MarshalIndent(comercios, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	db.WriteToBucket(bucket, []byte(strconv.Itoa(nrocomercio)), data)
}

func ComercioSelectBolt(id int) {
	bucket := "comercio"
	data, _ := db.ReadUniqueFromBucket(bucket, []byte(strconv.Itoa(id)))

	var comercio Comercio

	err := json.Unmarshal(data, &comercio)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n", comercio)
}
