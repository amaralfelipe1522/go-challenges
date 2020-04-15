package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
Acceptance Criteria
In order to successfully complete this challenge, your project will have to:

Expose a CRUD API that will have an in-memory list of shopping cart items.
Whenever an endpoint is hit within your API, it will have to use logrus to log that event to a file in JSON format.
*/

func insert(tx *sql.Tx, prod string, prodQtd int) {
	stmt, _ := tx.Prepare("insert into cart(prod, prod_qtd) values (?,?)")

	_, err := stmt.Exec(prod, prodQtd)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("Uma linha inserida.")
}

func update(tx *sql.Tx, prod string, prodQtd int, id int) {
	stmt, _ := tx.Prepare("update cart set prod = ?, prod_qtd = ? where id = ?")

	_, err := stmt.Exec(prod, prodQtd, id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("Uma linha atualizada.")
}

func delete(tx *sql.Tx, id int) {
	stmt, _ := tx.Prepare("delete from cart where id = ?")

	_, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("Uma linha deletada.")
}

func main() {
	db, err := sql.Open("mysql", "root:Project@1522@/store")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	//delete(tx, 6)
	//insert(tx, "anador", 6)
	update(tx, "desodorante", 2, 5)
}
