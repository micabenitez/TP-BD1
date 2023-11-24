package tablas

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"fmt"
	"log"
)

func AlumneCreateTable() {
	db.CreateTable("alumne", "legajo int, nombre text, apellido text")
}

func AlumneDropTable() {
	db.DropTable("alumne")
}

func AlumneInsert(legajo int, nombre string, apellido string) {
	values := fmt.Sprintf("%d, '%s', '%s'", legajo, nombre, apellido)
	db.Insert("alumne", values)
}

func AlumneSelect(query string) {
	rows := db.Select("alumne", "legajo, nombre, apellido", query)

	type alumne struct {
		legajo           int
		nombre, apellido string
	}

	var a alumne

	for rows.Next() {
		if err := rows.Scan(&a.legajo, &a.nombre, &a.apellido); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v\n", a.legajo, a.nombre, a.apellido)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
