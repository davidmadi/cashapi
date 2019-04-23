package migration

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Main() {

}

func init() {
	connStr := "user=root dbname=cashdb password=hagadol23 sslmode=disable port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT version FROM migration")
	if err != nil {
		createMigration(db)
	} else {
		var version int64
		rows.Next()
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("version is %d", version)
			if version == 1 {
				version1(db)
			}
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
	db.Exec(`update migration set version=2;`)
	log.Printf("version updated to %d", 2)
}
