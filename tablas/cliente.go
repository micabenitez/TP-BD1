package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func ClienteCreateTable() {
	db.CreateTable("cliente", "nrocliente int, nombre text, apellido text, domicilio text, telefono char(12)")
}

func ClienteCreatePK() {
	db.CreatePK("cliente", "cliente_pk", "nrocliente")
}

func ClienteDropTable() {
	db.DropTable("cliente")
}

func ClienteDeletePK() {
	db.DeleteKey("cliente", "cliente_pk")
}

func ClienteInsert(nrocliente int, nombre string, apellido string, domicilio string, telefono string) {
	if len(telefono) > 12 {
		telefono = telefono[0:11]
	}
	values := fmt.Sprintf("%d,'%s','%s','%s','%s'", nrocliente, nombre, apellido, domicilio, telefono)
	db.Insert("cliente", values)
}

func ClienteSelect(query string) {
	rows := db.Select("cliente", "nrocliente, nombre, apellido, domicilio,telefono", query)

	type Cliente struct {
		nroCliente                            int
		nombre, apellido, domicilio, telefono string
	}

	var a Cliente

	for rows.Next() {
		if err := rows.Scan(&a.nroCliente, &a.nombre, &a.apellido, &a.domicilio, &a.telefono); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v %v %v\n", a.nroCliente, a.nombre, a.apellido, a.domicilio, a.telefono)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

type Cliente struct {
	NroCliente                            int
	Nombre, Apellido, Domicilio, Telefono string
}

func ClienteInsertBolt(nroclienteB int, nombreB string, apellidoB string, domicilioB string, telefonoB string) {
	bucket := "cliente"

	if len(telefonoB) > 12 {
		telefonoB = telefonoB[0:11]
	}

	clientes := Cliente{NroCliente: nroclienteB,
		Nombre:    nombreB,
		Apellido:  apellidoB,
		Domicilio: domicilioB,
		Telefono:  telefonoB}

	data, err := json.MarshalIndent(clientes, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	db.WriteToBucket(bucket, []byte(strconv.Itoa(nroclienteB)), data)
}

func ClienteSelectBolt(id int) {
	bucket := "cliente"
	data, _ := db.ReadUniqueFromBucket(bucket, []byte(strconv.Itoa(id)))

	var cliente Cliente

	err := json.Unmarshal(data, &cliente)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n", cliente)
}

func ClienteSelectAll() {
	db.SelectAllFromBucket("cliente")
}
