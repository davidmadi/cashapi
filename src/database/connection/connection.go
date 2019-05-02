package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Connection struct {
	DB     *sql.DB
	opened bool
}

//Open postgres connection
func Open() (Connection, error) {
	connStr := "user=root dbname=cashdb password=hagadol23 sslmode=disable port=5432"
	dbCon, err := sql.Open("postgres", connStr)

	con := Connection{
		DB:     dbCon,
		opened: false}

	if err != nil {
		log.Fatal(err)
		return con, err
	}
	con.opened = true

	return con, nil
}

//ExecuteQuery given the sql
func (con Connection) ExecuteQuery(sql string, list interface{}) bool {

	rows, err := con.DB.Query(sql)
	if err != nil {
		log.Fatal(err)
		return false
	}

	rows.Scan(list)
	return true

}
