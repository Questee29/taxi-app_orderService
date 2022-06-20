package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upOrders, downOrders)
}

func upOrders(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS orders(
		"id" SERIAL PRIMARY KEY,
		"user_id" int,
		"driver_id" int,
		"start_point" text,
		"end_point" text,
		"taxi_type" text,
		"order_date" date,
		"status" text,
		"user_rate" real,
		"driver_rate" real);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downOrders(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS orders;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
