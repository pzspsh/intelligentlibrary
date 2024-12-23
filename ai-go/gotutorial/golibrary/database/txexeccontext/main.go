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
	id := 37
	_, execErr := tx.ExecContext(ctx, "UPDATE users SET status = ? WHERE id = ?", "paid", id)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
