package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	disk_db_path := flag.String("leveldb", "", "Path to LevelDB")
	verbose := flag.Bool("verbose", false, "Verbose output")
	flag.Parse()

	if *verbose {
		log.SetLevel(log.TraceLevel)
	}

	disk_db, err := leveldb.OpenFile(*disk_db_path, nil)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to open LevelDB: %s", err))
	}
	defer disk_db.Close()
	store := KeyValueStore{disk: disk_db}

	server := Server{store: &store}
	http.ListenAndServe(":3000", server)
}
