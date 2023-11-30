/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 13:49:49
*/
package main

import (
	"context"
	"database/sql"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	id := 53
	_, err = tx.ExecContext(ctx, "UPDATE drivers SET status = ? WHERE id = ?;", "assigned", id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
		}
		log.Fatal(err)
	}
	_, err = tx.ExecContext(ctx, "UPDATE pickups SET driver_id = $1;", id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to back: %v", err, rollbackErr)
		}
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
