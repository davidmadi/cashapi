package migration

import (
	"cashapi/src/database/connection"
	"cashapi/src/model"
	"log"

	"github.com/jinzhu/gorm"
	//_ "github.com/lib/pq"
)

// Migration table model
type Migration struct {
	gorm.Model
	Version uint
}

func Main() {

}

func init() {

	//connStr := "user=root dbname=cashdb password=hagadol23 sslmode=disable port=5432"
	con, err := connection.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err := con.DB().Query("SELECT Version FROM migrations")
	if err != nil {
		createMigration(con)
		con.CreateTable(model.User{}, model.PaymentModel{})
		version1(con)
	} else {
		var version int64
		rows.Next()
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Version is %d", version)
			if version == 1 {
				version2(con)
			} else if version == 2 {
				version3(con)
			} else if version == 3 {
				version4(con)
			}
		}
		rows.Close()
	}
	con.Close()
}

func createMigration(con *gorm.DB) {
	con.CreateTable(Migration{})
	con.Save(&Migration{
		Version: 1,
	})

	thisMigration := Migration{}
	con.Table("migrations").Scan(&thisMigration)
	if thisMigration.Version == 0 {
		log.Fatalf("tabela migrations nao foi criada")
	}
}

func version1(con *gorm.DB) {
	con.DB().Exec(`update migrations set version=1;`)
	log.Printf("version updated to %d", 2)
	version2(con)
}

func version2(con *gorm.DB) {
	con.CreateTable(model.Group{})
	con.DB().Exec(`update migrations set version=2;`)
	log.Printf("version updated to %d", 2)
	version3(con)
}

func version3(con *gorm.DB) {
	con.DropTableIfExists(model.Group{})
	con.CreateTable(model.Group{})
	con.Save(&model.Group{
		Name: "Administradores",
	})
	con.Save(&model.Group{
		Name: "Consultores",
	})
	con.DB().Exec(`update migrations set version=3;`)
	log.Printf("version updated to %d", 3)
	version4(con)
}

func version4(con *gorm.DB) {
	con.Create(model.Group{
		Name: "Administradores",
	})
	con.Create(model.Group{
		Name: "Consultores",
	})
	con.DB().Exec(`update migrations set version=4;`)
	log.Printf("version updated to %d", 4)
}
