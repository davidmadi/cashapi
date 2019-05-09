package connection

import (
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//Open postgres connection
func Open() (db *gorm.DB, err error) {

	connStr := "user=root dbname=cashdb password=hagadol23 sslmode=disable port=5432"
	dbCon, err := gorm.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = dbCon.DB().Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return dbCon, nil
}

//ExecuteQuery given the sql
func ExecuteQuery(sql string, con *gorm.DB) (*sql.Rows, bool) {
	rows, err := con.CommonDB().Query(sql, nil)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	return rows, true
}
