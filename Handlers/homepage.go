package Handlers

import (
	"html/template"
	"log"
	"net/http"

	"recipekeeper/data"
)

func HomePage(w http.ResponseWriter, r *http.Request){
tmpl, err :=template.ParseFiles("FrontEnd/homepage.html")
if err != nil {
	log.Fatal(err, "ERROR: problem with home file path")
}
tmpl.Execute(w, data.AllRecipes)
}
