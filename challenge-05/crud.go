package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/*
Acceptance Criteria
In order to successfully complete this challenge, your project will have to:

Expose a CRUD API that will have an in-memory list of shopping cart items.
Whenever an endpoint is hit within your API, it will have to use logrus to log that event to a file in JSON format.
*/

//Produto é utilizada para armazenar as consultas ao banco de dados
type Produto struct {
	ID      int    `json:"id"`
	Prod    string `json:"prod"`
	ProdQtd int    `json:"prod_qtd"`
}

func insert(prod string, prodQtd int) {
	db := openDB()
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into cart(prod, prod_qtd) values (?,?)")

	_, err := stmt.Exec(prod, prodQtd)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("Uma linha inserida.")
}

func update(prod string, prodQtd int, id int) {
	db := openDB()
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("update cart set prod = ?, prod_qtd = ? where id = ?")

	_, err := stmt.Exec(prod, prodQtd, id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("Uma linha atualizada.")
}

func delete(id int) {
	db := openDB()
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("delete from cart where id = ?")
	_, err := stmt.Exec(id)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	fmt.Println("Uma linha deletada.")
}

func (p *Produto) selectOne(w http.ResponseWriter, r *http.Request, id int) {
	db := openDB()
	defer db.Close()

	tx, _ := db.Begin()

	tx.QueryRow("select * from cart where id = ?", id).Scan(&p.ID, &p.Prod, &p.ProdQtd)
	//fmt.Println(*p)

	json, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func (p *Produto) selectAll(w http.ResponseWriter, r *http.Request) {
	db := openDB()
	defer db.Close()

	tx, _ := db.Begin()

	rows, err := tx.Query("select * from cart")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pList []Produto

	for rows.Next() {
		rows.Scan(&p.ID, &p.Prod, &p.ProdQtd)
		pList = append(pList, *p)
	}

	json, _ := json.Marshal(pList)
	//fmt.Println(string(json))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Project@1522@/store")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//Orquestrador é responsável por identificar qual operação de CRUD será disparada a partir da chamada do server
func Orquestrador(w http.ResponseWriter, r *http.Request) {
	var p Produto
	crudType := strings.TrimPrefix(r.URL.Path, "/cart/")

	switch {
	case r.Method == "GET" && crudType == "selectone":
		p.selectOne(w, r, 1)
	case r.Method == "GET" && crudType == "selectall":
		p.selectAll(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Método de requisição diferente ou URL inválida.")
	}

	//delete(9)
	//insert("anador", 6)
	//update("fio dental", 3, 8)
	//p.selectOne(1)
	//p.selectAll()
}
