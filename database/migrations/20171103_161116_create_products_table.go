package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateProductsTable_20171103_161116 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateProductsTable_20171103_161116{}
	m.Created = "20171103_161116"

	migration.Register("CreateProductsTable_20171103_161116", m)
}

// Run the migrations
func (m *CreateProductsTable_20171103_161116) Up() {
	m.SQL(`
		CREATE TABLE products(
			id                  BIGSERIAL PRIMARY KEY NOT NULL,
			name                TEXT NOT NULL,
			type                INTEGER DEFAULT 0,
			price               INTEGER DEFAULT 0,
			url                 TEXT,
			image               TEXT,
			breakpoint_up       INT,
			breakpoint_down     INT,
			amount              INT,
			updated_at	        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at	        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

// Reverse the migrations
func (m *CreateProductsTable_20171103_161116) Down() {
	m.SQL("DROP TABLE products")
}
