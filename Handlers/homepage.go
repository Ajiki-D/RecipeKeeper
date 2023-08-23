package Handlers
import (
	"net/http"
	"fmt"
)

func HomePage(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "Hello")
}
