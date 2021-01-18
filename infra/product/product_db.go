package infra

import (
	"log"

	"github.com/site/entity"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type UserSQL struct {
	db *gorm.DB
}

func NewProductSQL(db *gorm.DB) *UserSQL {
	return &UserSQL{
		db: db,
	}
}
func (d *UserSQL) Buscar(id int32) []*entity.Product {
	products := []*entity.Product{}
	d.db.Find(&products)

	return products
}
func GetDatabase() *gorm.DB {
	dsn := "sqlserver://guser:guser@localhost:1433/?database=golang"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Println("Failed to ping: ", err.Error())
	}
	return db
}
