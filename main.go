package main

import (
	"net/http"
	"github.com/Ajiki-D/REcipekeeper/Handlers"
)

func main() {
http.HandleFunc("/", Handlers.HomePage)
http.ListenAndServe(":8000", nil)

}