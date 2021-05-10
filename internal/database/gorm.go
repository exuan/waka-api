package database

import (
	"fmt"
	"log"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	Cfg struct {
		Name   string `toml:"name"`
		Dsn    string `toml:"dsn"`
		Type   string `toml:"type"`
		Prefix string `toml:"prefix"`
	}
)

var (
	NotFoundDB  = fmt.Errorf("not found DB")
	defaultName = "default"
	dbs         = make(map[string]*gorm.DB)
)

func Connect(cfgs []*Cfg) error {

	for index, c := range cfgs {
		var dialect gorm.Dialector
		switch c.Type {
		case "mysql":
			dialect = mysql.Open(c.Dsn)
		case "postgres":
			dialect = postgres.Open(c.Dsn)
		case "sqlite":
			dialect = sqlite.Open(c.Dsn)
		//case "sqlserver":
		//	dialect = sqlserver.Open(c.Dsn)
		default:
			return fmt.Errorf("unavailable db type: %s", c.Type)
		}

		db, err := gorm.Open(dialect, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   c.Prefix,
				SingularTable: true,
			}})

		if err != nil {
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		if err = sqlDB.Ping(); err != nil {
			return fmt.Errorf("ping %s database err: %s", c.Name, err)
		}

		dbs[c.Name] = db

		// first or  name is default
		if index == 0 || c.Name == "default" {
			defaultName = c.Name
		}
	}
	log.Println(dbs)
	if len(dbs) == 0 {
		return fmt.Errorf("none database")
	}

	return nil
}

func Default() string {
	return defaultName
}

func DB(n string) *gorm.DB {
	return dbs[n]
}

func Close() error {
	for _, db := range dbs {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		sqlDB.Close()
	}
	return nil
}
