package persistence

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/Eldius/speedtest/speedtest"
	"github.com/jinzhu/gorm"

	// it's to use SQLite db
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
)

const (
	dbFolder = "~/.speedtest"
	dbFile   = "speedtest.db"
)

func init() {
	if path, err := homedir.Expand(dbFolder); err == nil {
		os.MkdirAll(path, os.ModePerm)
	} else {
		panic(err.Error())
	}
}

func getDbFile() string {
	if path, err := homedir.Expand(dbFolder); err != nil {
		panic(err.Error())
	} else {
		return filepath.Join(path, dbFile)
	}
}

func getDB() *gorm.DB {
	databasePath := getDbFile()
	log.Println("db file:", databasePath)
	if db, err := gorm.Open("sqlite3", databasePath); err != nil {
		panic(err.Error())
	} else {
		db.AutoMigrate(&speedtest.TestServer{}, &speedtest.SelectedServer{})
		return db
	}
}

/*
SaveServer saves a server
*/
func SaveServer(s speedtest.TestServer) {
	getDB().Save(s)
}

/*
SaveAllServers saves a list of servers
*/
func SaveAllServers(servers []speedtest.TestServer) {
	db := getDB()
	defer db.Close()
	tx := db.Begin()
	for _, s := range servers {
		debug(s)
		db.Save(s)
	}
	tx.Commit()
}

func debug(o interface{}) {
	if b, err := json.Marshal(o); err == nil {
		log.Println("---")
		log.Println(string(b))
		log.Println("---")
	} else {
		panic(err.Error())
	}
}
