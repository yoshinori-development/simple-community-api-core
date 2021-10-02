package repositories

import (
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/yoshinori-development/simple-community-api-main/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDbMain() error {
	config := config.Get()
	err := open(config)
	if err != nil {
		return err
	}

	err = migrateExec()
	if err != nil {
		return err
	}
	return nil
}

func open(config config.Config) error {
	var err error
	dbConf := config.Database
	datasource := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name)
	db, err = gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database")
	}

	return nil
}

func Close() error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDb.Close()
}

func migrateExec() error {
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
