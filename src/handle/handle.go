package handle
import (
	"net/http"
	"fmt"
)
func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleDefault")
}
