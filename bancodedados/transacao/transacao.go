package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")
	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
