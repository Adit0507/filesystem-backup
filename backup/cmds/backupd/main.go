package main

import (
	"flag"
	"log"
	"time"

	"github.com/Adit0507/filesystem-backup/backup"
	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		// represents no. of seconds b/w chekcs to see whether folders have changed 
		interval  = flag.Duration("interval", 10*time.Second, "interval between checks")
		
		// path to archve location where ZIP files go
		archive = flag.String("archive", "archive", "path to archive location")
		
		// path to same filedb database
		dbPath = flag.String("db", "./db", "path to filedb database")
	)

	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver: backup.ZIP,
		Paths: make(map[string]string),
	}

	db, err := filedb.Dial(*dbPath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()

	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}
}
