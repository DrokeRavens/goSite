package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/site/db"
	"github.com/site/domain/product"
	_ "gorm.io/driver/sqlserver"
	_ "gorm.io/gorm"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	dbs := db.GetDatabase()
	repo := db.NewProductSQL(dbs)
	service := product.NewService()
	produtos := service.Buscar(0)

	fmt.Println(produtos)
	temp.ExecuteTemplate(w, "index", produtos)
}
