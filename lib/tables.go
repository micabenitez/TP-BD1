package lib

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ExecuteQueryCompra(query string) {
	db := Connection()
	defer db.Close()

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Minute)

}

func ExecuteQuery(query string) {
	db := Connection()
	defer db.Close()
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTable(tableName string, values string) {
	db := Connection()
	defer db.Close()
	//fmt.Println("Create Table " + tableName)
	_, err := db.Exec(`create table if not exists ` + tableName + ` (` + values + `);`)
	if err != nil {
		log.Fatal(err)
	}
}

func DropTable(tableName string) {
	db := Connection()
	defer db.Close()
	//fmt.Println("Delete Table" + tableName)
	_, err := db.Exec(`drop table if exists ` + tableName + `;`)
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(tableName string, values string) {
	db := Connection()
	defer db.Close()
	query := fmt.Sprintf("insert into %s values(%s);", tableName, values)
	fmt.Println(query)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePK(tableName string, pkName string, columNames string) {
	db := Connection()
	defer db.Close()
	//alter table cursa add constraint cursa_legajo_fk foreign key (legajo) references alumne (legajo);
	query := fmt.Sprintf("alter table %s add constraint %s primary key (%s);", tableName, pkName, columNames)
	fmt.Println(query)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFK(tableNameOrigin string, fkName string, columNamesOrigin string, tableNameDestiny string, columNamesDestiny string) {
	db := Connection()
	defer db.Close()
	//alter table cursa add constraint cursa_legajo_fk foreign key (legajo) references alumne (legajo);
	//alter table tableOrigen add constraint fkName foreign key (columnasOrigen) references tablaDestino (columnasDestino);
	query := fmt.Sprintf("alter table %s add constraint %s foreign key (%s) references %s (%s);", tableNameOrigin, fkName, columNamesOrigin, tableNameDestiny, columNamesDestiny)
	fmt.Println(query)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteKey(tableName string, keyName string) {
	db := Connection()
	defer db.Close()
	//alter table cursa add constraint cursa_legajo_fk foreign key (legajo) references alumne (legajo);
	query := fmt.Sprintf("alter table if exists %s drop constraint if exists %s cascade;", tableName, keyName)
	fmt.Println(query)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func Select(tableName string, columName string, query string) *sql.Rows {
	db := Connection()
	defer db.Close()
	//fmt.Println("*Select Data*")
	rows, err := db.Query(`select ` + columName + ` from ` + tableName + ` ` + query + `;`)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
