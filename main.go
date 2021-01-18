package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/site/domain/product"
	infra "github.com/site/infra/product"
	_ "gorm.io/driver/sqlserver"
	_ "gorm.io/gorm"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	dbs := infra.GetDatabase()
	repo := infra.NewProductSQL(dbs)
	service := product.NewService(repo)
	produtos := service.Buscar(0)

	fmt.Println(produtos)
	temp.ExecuteTemplate(w, "index", produtos)
}
