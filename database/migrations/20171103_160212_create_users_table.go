package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsersTable_20171103_160212 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20171103_160212{}
	m.Created = "20171103_160212"

	migration.Register("CreateUsersTable_20171103_160212", m)
}

// Run the migrations
func (m *CreateUsersTable_20171103_160212) Up() {
	m.SQL(`
		CREATE TABLE users(
			id          BIGSERIAL PRIMARY KEY NOT NULL,
			email       TEXT NOT NULL,
			password    TEXT NOT NULL,
			first_name  TEXT NOT NULL,
			last_name   TEXT NOT NULL,
			image       TEXT,
			updated_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

// Reverse the migrations
func (m *CreateUsersTable_20171103_160212) Down() {
	m.SQL("DROP TABLE users")
}
