package utils

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"mux_test/model"
)

func ConnectDB() {
	pgUrl, err := pq.ParseURL("postgres://rcdgiapl:xZhrR4X2tYJfS5BI6XLbCg7dcg-_4R4i@arjuna.db.elephantsql.com/rcdgiapl")
	if err != nil {
		log.Fatal("1", err)
	}

	model.DB, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal("2", err)
	}

	err = model.DB.Ping()
}
