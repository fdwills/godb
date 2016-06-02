package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

const (
	READ_MODE = iota
	WRITE_MODE
)

type db_accessor struct {
	Read  *gorm.DB
	Write *gorm.DB
}

var master_db *db_accessor

func GetMasterDB(mode int) *gorm.DB {
	if mode == READ_MODE {
		return master_db.Read
	} else {
		return master_db.Write
	}
}

func InitDB() {
	config := GetConfig()
	// master
	{
		log.Printf("open master read")
		read_db, dberr := gorm.Open(config.DBDriver, config.DBMaster.Read)
		if dberr != nil {
			log.Fatalf("failed to open read db: %v", dberr)
		}

		read_db.DB().SetMaxOpenConns(config.DBMaster.MaxPoolSize)
		if config.Mode == "debug" {
			read_db.LogMode(true)
		}

		log.Printf("open master write")
		write_db, dberr := gorm.Open(config.DBDriver, config.DBMaster.Write)
		if dberr != nil {
			log.Fatalf("failed to open write db: %v", dberr)
		}
		write_db.DB().SetMaxOpenConns(config.DBMaster.MaxPoolSize)
		if config.Mode == "debug" {
			write_db.LogMode(true)
		}
		master_db = &db_accessor{read_db, write_db}
	}
}
