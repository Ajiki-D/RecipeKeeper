package main

import (
	"net/http"
	"recipekeeper/data"
	"recipekeeper/Handlers"
)

func main() {
	data.FetchAllRecipes()
	//fmt.PrintLn(data.AllRecipes)
http.HandleFunc("/", Handlers.HomePage)
http.ListenAndServe(":8000", nil)

}