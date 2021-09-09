package db_core

import (
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/yoshinori-development/simple-community-api-core/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open(config config.Config) (*gorm.DB, error) {
	dbConf := config.Database
	datasource := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDb.Close()
}

func Migrate(db *gorm.DB) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "setup/migrations",
	}
	migrate.SetTable("migrations")

	n, err := migrate.Exec(sqlDb, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
