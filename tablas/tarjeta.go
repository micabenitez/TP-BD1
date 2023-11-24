package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func TarjetaCreateTable() {
	db.CreateTable("tarjeta", "nrotarjeta char(16), nrocliente int, validadesde char(6), validahasta char(6), codseguridad char(4),limitecompra decimal(8,2),estado char(10)")
}

func TarjetaCreatePK() {
	db.CreatePK("tarjeta", "tarjeta_pk", "nrotarjeta")
}

func TarjetaCreateFK() {
	db.CreateFK("tarjeta", "tarjeta_nrocliente_fk", "nrocliente", "cliente", "nrocliente")
}

func TarjetaDeletePK() {
	db.DeleteKey("tarjeta", "tarjeta_pk")
}

func TarjetaDeleteFK() {
	db.DeleteKey("tarjeta", "tarjeta_nrocliente_fk")
}

func TarjetaDropTable() {
	db.DropTable("tarjeta")
}

func TarjetaInsert(
	nrotarjeta string,
	nrocliente int,
	validadesde string,
	validahasta string,
	codseguridad string,
	limitecompra float64,
	estado string) {

	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}
	if len(validadesde) > 6 {
		validadesde = validadesde[0:5]
	}
	if len(validahasta) > 6 {
		validahasta = validahasta[0:5]
	}
	if len(codseguridad) > 3 {
		codseguridad = codseguridad[0:3]
	}
	if len(estado) > 10 {
		estado = estado[0:10]
	}

	values := fmt.Sprintf("'%s', %d,'%s','%s','%s',%f,'%s'",
		nrotarjeta, nrocliente, validadesde, validahasta, codseguridad, limitecompra, estado)
	db.Insert("tarjeta", values)
}

func TarjetaSelect(query string) {
	rows := db.Select("tarjeta", "nroTarjeta, nrotarjeta, nrocliente, validadesde, validahasta, codseguridad, limitecompra, estado", query)

	type tarjeta struct {
		nrotarjeta   string
		nrocliente   int
		validadesde  string
		validahasta  string
		codseguridad string
		limitecompra float64
		estado       string
	}

	var a tarjeta

	for rows.Next() {
		if err := rows.Scan(a.nrotarjeta, a.nrocliente, a.validadesde, a.validahasta, a.codseguridad, a.limitecompra, a.estado); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v %v %v %v %v\n", a.nrotarjeta, a.nrocliente, a.validadesde, a.validahasta, a.codseguridad, a.limitecompra, a.estado)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

type Tarjeta struct {
	Nrotarjeta   string
	Nrocliente   int
	Validadesde  string
	Validahasta  string
	Codseguridad string
	Limitecompra float64
	Estado       string
}

func TarjetaInsertBolt(nrotarjeta string,
	nrocliente int,
	validadesde string,
	validahasta string,
	codseguridad string,
	limitecompra float64,
	estado string) {
	bucket := "tarjeta"

	nrotarjeta = strings.ReplaceAll(nrotarjeta, " ", "")
	if len(nrotarjeta) > 16 {
		nrotarjeta = nrotarjeta[0:15]
	}
	if len(validadesde) > 6 {
		validadesde = validadesde[0:5]
	}
	if len(validahasta) > 6 {
		validahasta = validahasta[0:5]
	}
	if len(codseguridad) > 3 {
		codseguridad = codseguridad[0:3]
	}
	if len(estado) > 10 {
		estado = estado[0:10]
	}

	tarjetas := Tarjeta{Nrotarjeta: nrotarjeta,
		Nrocliente:   nrocliente,
		Validadesde:  validadesde,
		Validahasta:  validahasta,
		Codseguridad: codseguridad,
		Limitecompra: limitecompra,
		Estado:       estado}

	data, err := json.MarshalIndent(tarjetas, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	db.WriteToBucket(bucket, []byte(nrotarjeta), data)
}

func TarjetaSelectBolt(id int) {
	bucket := "tarjeta"
	data, _ := db.ReadUniqueFromBucket(bucket, []byte(strconv.Itoa(id)))

	var tarjeta Tarjeta

	err := json.Unmarshal(data, &tarjeta)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n", tarjeta)
}

func TarjetaSelectAll() {
	db.SelectAllFromBucket("tarjeta")
}
