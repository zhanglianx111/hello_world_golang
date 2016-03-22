package main
import (
	"net/http"
	"handle"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handle.HandlerDefault)
	r.HandleFunc("/hi", handle.HandlerHi)
	r.HandleFunc("/name", handle.HandlerName)
	r.HandleFunc("/age", handle.HandlerAge)
	http.ListenAndServe(":8000", r)
}
