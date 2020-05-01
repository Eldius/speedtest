package persistence

import (
	"os"
	"path/filepath"

	"github.com/Eldius/speedtest/speedtest"
	"github.com/jinzhu/gorm"
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
	}
}

func getDB() *gorm.DB {
	if db, err := gorm.Open("sqlite3", filepath.Join(dbFolder, dbFile)); err != nil {
		panic(err.Error())
	} else {
		db.AutoMigrate(&speedtest.TestServer{})
		return db
	}
}

/*
Save saves an entity
*/
func Save(e interface{}) {
	getDB().Create(e)
}

/*
SaveAll saves a list of entities
*/
func SaveAll(entities ...interface{}) {
	db := getDB()
	defer db.Close()
	tx := db.Begin()
	for _, e := range entities {
		tx.Create(e)
	}
	tx.Commit()
}
