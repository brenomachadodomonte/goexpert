package database

import "database/sql"

func Migrate(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS categories (id string, name string, description string)")

	return err
}
