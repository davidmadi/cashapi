package connection

import (
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Connection struct {
	DB     *gorm.DB
	opened bool
}

//Open postgres connection
func Open() (Connection, error) {

	connStr := "user=root dbname=cashdb password=hagadol23 sslmode=disable port=5432"
	dbCon, err := gorm.Open("postgres", connStr)

	con := Connection{
		DB:     dbCon,
		opened: false}

	if err != nil {
		log.Fatal(err)
		return con, err
	}

	err = dbCon.DB().Ping()
	if err != nil {
		log.Fatal(err)
		return con, err
	}
	con.opened = true
	return con, nil
}

//ExecuteQuery given the sql
func (con Connection) ExecuteQuery(sql string) (*sql.Rows, bool) {
	rows, err := con.DB.DB().Query(sql)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	//rows.Scan(list)
	return rows, true
}

//Close the connection
func (con Connection) Close() bool {
	error := con.DB.Close()
	if error != nil {
		log.Fatalf("Erro on closing db connection %v", error)
		return false
	}
	return true
}
