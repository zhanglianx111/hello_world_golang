package handle
import (
	"net/http"
	"fmt"
)
func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleDefault")
}

func HandlerHi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleHi")
}

func HandlerName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleName")
}

func HandlerAge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleAge")
}
