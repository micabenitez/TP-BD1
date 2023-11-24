package main

import ( //funciones de alta de tablas y datos
	//conexiones y alta de base
	fx "alpuin-benitez-chianese-zanardi-tp/funciones"
	cx "alpuin-benitez-chianese-zanardi-tp/lib"
	fmt "fmt"
	"strings"

	_ "github.com/lib/pq"
)

func main() {

	var option string
	cx.SelectAllFromBucket("cliente")
	// ------------- MENU ------------
	for {
		fmt.Println(`
		1 - Crear Base de Datos
		2 - Crear Tablas		
		3 - Crear PKs y FKs		
		4 - Crear Funciones y Triggers		
		C - Crear Rapido (2, 3, 4)

		5 - Insertar Datos
		
		T - Correr Test
		A - Generar Resumen
		N - No SQL
		
		6 - Eliminar Base de Datos
		7 - Eliminar Tablas
		8 - Eliminar PKs y FKs		
		9 - Eliminar Funciones y Triggers
		B - Borrar Rapido (9, 8, 7)
		
		0 - Terminar
		`)
		fmt.Scanln(&option)

		if option == "0" {
			cx.DropDatabase()
			break
		}

		optionPicker(option)
	}
}

func optionPicker(option string) {
	/*
		1 - Crear Base de Datos
		2 - Crear Tablas
		3 - Crear PKs y FKs
		4 - Crear Funciones y Triggers
		C - Crear Rapido (2, 3, 4)

		5 - Insertar Datos

		T - Correr Test
		A - Generar Resumen
		N - No SQL

		6 - Eliminar Base de Datos
		7 - Eliminar Tablas
		8 - Eliminar PKs y FKs
		9 - Eliminar Funciones y Triggers
		B - Borrar Rapido (9, 8, 7)

		0 - Terminar
	*/
	if option == "1" {
		cx.CreateDatabase()
		fmt.Printf("Creando database ... \n")
		return
	}
	if option == "2" {
		fx.CreateTables()
		fmt.Printf("Creando tablas ... \n")
		return
	}
	if option == "3" {
		fx.CreatePrimaryKey()
		fx.CreateForeignKey()
		fmt.Printf("Agregando PKs y FKs ... \n")
		return
	}
	if option == "4" {
		fx.CreateFuncAndTrigger()
		fmt.Printf("Creando funciones ... \n")
		return
	}
	if option == "5" {
		fx.InsertData()
		fmt.Printf("Insertando datos... \n")
		return
	}
	if option == "6" {
		cx.DropDatabase()
		return
	}
	if option == "7" {
		fx.DropTables()
		fmt.Printf("Eliminando tablas ... \n")
		return
	}
	if option == "8" {
		fx.DeleteForeignKey()
		fx.DeletePrimaryKey()
		fmt.Printf("Eliminando PKs y FKs ... \n")
		return
	}
	if option == "9" {
		fx.DeleteFuncAndTrigger()
		return
	}
	if strings.ToUpper(option) == "T" {
		fx.GenerarConsumos()
		fmt.Printf("Generando consumos ... \n")
		return
	}
	if strings.ToUpper(option) == "N" {
		fx.InsertDataBolt()
		fx.SelectDataBolt()
		return
	}
	if strings.ToUpper(option) == "C" {
		fx.CreateTables()
		fx.CreatePrimaryKey()
		fx.CreateForeignKey()
		fx.CreateFuncAndTrigger()
		return
	}
	if strings.ToUpper(option) == "B" {
		fx.DeleteFuncAndTrigger()
		fx.DeleteForeignKey()
		fx.DeletePrimaryKey()
		fx.DropTables()
		return
	}
	if strings.ToUpper(option) == "A" {
		fx.TestResumen()
		return
	}

}
