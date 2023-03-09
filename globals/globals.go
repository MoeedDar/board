package globals

import (
	"context"
	"log"
	"path/filepath"

	"github.com/genjidb/genji"
)

var DB *genji.DB
var TemplatesDir string

func initDB() {
	db, err := genji.Open("database")
	if err != nil {
		log.Fatal(err)
	}

	db = db.WithContext(context.Background())

	err = db.Exec(`
		CREATE TABLE IF NOT EXISTS boards (
			title TEXT PRIMARY KEY
		);

		CREATE SEQUENCE IF NOT EXISTS post_id_seq;
		
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY DEFAULT(NEXT VALUE FOR post_id_seq),
			content TEXT,
			timestamp TEXT,
			board TEXT
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func initTemplatesDir() {
	absPath, err := filepath.Abs("./templates")
	if err != nil {
		log.Fatal(err)
	}
	TemplatesDir = absPath
}

func init() {
	initDB()
	initTemplatesDir()
}
