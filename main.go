package main

import (
	"html/template"
	"log"
	"net/http"
)

type Car struct {
	Name  string
	Model string
	Year  int
	Price string
}

var (
	// Pre-parse the template sets
	homeTmpl = template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/index.html",
		"templates/car-card.html",
	))

	compTmpl = template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/comparisons.html",
	))
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Cars": []Car{
				{Name: "Tesla", Model: "Model 3", Year: 2024, Price: "$38,990"},
				{Name: "Rivian", Model: "R1S", Year: 2024, Price: "$74,900"},
			},
		}
		// Home Page
		homeTmpl.ExecuteTemplate(w, "layout.html", data)
	})

	http.HandleFunc("/compare", func(w http.ResponseWriter, r *http.Request) {
		// Comparison Page
		compTmpl.ExecuteTemplate(w, "layout.html", nil)
	})

	log.Println("Server active at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
