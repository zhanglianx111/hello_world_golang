package handle
import (
	"net/http"
	"fmt"
)
func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleDefault")
}

func HandleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleHi")
}

func HandleName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleName")
}

func HandleAge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleAge")
}
