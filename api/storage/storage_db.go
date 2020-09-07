package storage

import (
	"database/sql"
	"log"

	dbh "github.com/ndv6/tnotif/database"
	"github.com/ndv6/tnotif/helper"
	"github.com/ndv6/tnotif/models"
)

type database struct {
	db *sql.DB
}

func newConnection() database {
	DbURL := helper.GetEnv("DATABASE_URL")
	db, err := sql.Open("postgres", DbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = dbh.LoadSQLFile(db, "database/log_mail.sql")
	if err != nil {
		log.Fatal(err)
	}

	return database{
		db: db,
	}
}

func (o database) Create(obj models.LogMail) error {
	_, err := o.db.Exec("INSERT INTO log_mail(email, sent_at) VALUES ($1, $2);", obj.Email, obj.SentAt)
	return err
}

func (o database) List() ([]models.LogMail, error) {
	rows, err := o.db.Query("SELECT email, sent_at FROM log_mail LIMIT 10;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sentMails []models.LogMail
	for rows.Next() {
		var sentMail models.LogMail
		err = rows.Scan(&sentMail.Email, &sentMail.SentAt)
		if err != nil {
			return nil, err
		}
		sentMails = append(sentMails, sentMail)
	}

	return sentMails, nil
}
