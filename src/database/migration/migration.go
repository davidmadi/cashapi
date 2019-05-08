package migration

import (
	"cashapi/src/database/connection"
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
	//_ "github.com/lib/pq"
)

type Group struct {
	gorm.Model
	Code  string
	Price uint
}

func Main() {

}

func init() {

	//connStr := "user=root dbname=cashdb password=hagadol23 sslmode=disable port=5432"
	con, err := connection.Open()
	db := con.DB.DB()
	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err := db.Query("SELECT version FROM migration")
	if err != nil {
		createMigration(db)
		version1(db)
		version2(db)
		version3(db)
	} else {
		var version int64
		rows.Next()
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("version is %d", version)
			if version == 1 {
				version2(db)
				version3(db)
			} else if version == 2 {
				version3(db)
			}
			//2
		}
		rows.Close()
	}
	db.Close()
}

func createMigration(db *sql.DB) {
	_, err1 := db.Exec(`CREATE TABLE migration(version int);`)
	if err1 != nil {
		log.Fatalf("tabela migration nao foi criada, %v", err1)
	} else {
		_, err2 := db.Exec("insert into migration values(1);")
		if err2 != nil {
			log.Fatal("Não foi possível iniciar migration")
		}
	}
}

func version1(db *sql.DB) {
	db.Exec(`update migration set version=1;`)
	log.Printf("version updated to %d", 2)
}

func version2(db *sql.DB) {
	db.Exec(`CREATE TABLE users(
				id serial PRIMARY KEY,
				login VARCHAR (50) UNIQUE NOT NULL,
				password VARCHAR (50) NOT NULL,
				email VARCHAR (355) UNIQUE NOT NULL,
				created_on TIMESTAMP NOT NULL,
				last_login TIMESTAMP
			);`)
	db.Exec(`update migration set version=2;`)
	log.Printf("version updated to %d", 2)
}

func version3(db *sql.DB) {
	db.Exec(`alter TABLE users add column groupid int;`)
	db.Exec(`CREATE TABLE groups(id int, name varchar(10));`)
	db.Exec(`update migration set version=3;`)
	log.Printf("version updated to %d", 2)
}
