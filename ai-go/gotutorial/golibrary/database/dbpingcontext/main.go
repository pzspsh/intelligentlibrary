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
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	// Ping and PingContext may be used to determine if communication with
	// the database server is still possible.
	//
	// When used in a command line application Ping may be used to establish
	// that further queries are possible; that the provided DSN is valid.
	//
	// When used in long running service Ping may be part of the health
	// checking system.

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println(status)
}
