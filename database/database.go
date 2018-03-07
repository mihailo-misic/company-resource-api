package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	m "github.com/mihailo-misic/company-resource-api/models"
	c "github.com/mihailo-misic/company-resource-api/controllers"
	"log"
)

func Init() (db *gorm.DB) {
	// Open a db connection
	var err error
	db, err = gorm.Open("postgres", `
		host=localhost
		port=5432
		user=postgres
		dbname=crm
		password=secret
		sslmode=disable
`)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&m.User{}, &m.Product{})

	// Connect the controllers to the Database
	c.ConnectControllers(db)

	return
}
