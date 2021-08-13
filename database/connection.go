package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ones sync.Once
var gdb *gorm.DB

type DBConfig struct {
	Host     string
	DBName   string
	Port     string
	UserDB   string
	Password string
}

func (d *DBConfig) getConnectionString() string {
	cadena := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		d.Host,
		d.UserDB,
		d.Password,
		d.DBName,
		d.Port,
	)
	return cadena
}

func GetDBConnextion(conf *DBConfig) *gorm.DB {
	ones.Do(func() {
		schema := schema.NamingStrategy{SingularTable: true}
		//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(conf.getConnectionString()), &gorm.Config{NamingStrategy: schema})
		if err != nil {
			log.Fatalln("No se pudo connectar la base de datos:", err.Error())
		}
		gdb = db
	})
	return gdb
}
