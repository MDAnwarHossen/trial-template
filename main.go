package main

import (
	"html/template"
	"log"
	"net/http"
)

type Specs struct {
	Engine       string
	Horsepower   string
	Transmission string
	Drivetrain   string
}

type Car struct {
	ID             string
	Name           string
	Model          string // Added to match car-card.html
	Year           int
	Price          string // Added to match car-card.html
	Specifications Specs
}

var (
	// FIX 1: Include car-card.html in the parsing
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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Cars": []Car{
				// FIX 2: Provide the Model and Price fields
				{ID: "1", Name: "Tesla", Model: "Model 3", Year: 2024, Price: "$38,990"},
				{ID: "2", Name: "Ford", Model: "Mustang", Year: 2023, Price: "$45,000"},
			},
		}
		// Render
		if err := homeTmpl.ExecuteTemplate(w, "layout.html", data); err != nil {
			log.Println("Error executing home template:", err)
		}
	})

	http.HandleFunc("/comparisons", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Models": []Car{
				{
					ID:    "1",
					Name:  "Tesla",
					Model: "Model 3",
					Year:  2024,
					Specifications: Specs{
						Engine:     "Electric",
						Horsepower: "283 hp",
					},
				},
			},
		}
		if err := compTmpl.ExecuteTemplate(w, "layout.html", data); err != nil {
			log.Println("Error executing comparison template:", err)
		}
	})

	log.Println("Test server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
