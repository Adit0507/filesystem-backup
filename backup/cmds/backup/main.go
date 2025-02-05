package backup

import (
	"errors"
	"flag"
	"log"
)

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()

	var dbPath = flag.String("db", "./backupdata", "path to DB directory")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("invalid usage; must specify command")
		return
	}

}