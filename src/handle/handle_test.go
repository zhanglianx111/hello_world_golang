package handle
import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)
func Test_HandleDefault(t *testing.T) {
	fmt.Println("Test_HandleDefault")
}

func Test_HandleHi(t *testing.T) {
	fmt.Println("Test_HandleHi")
}

func Test_HandleName(t *testing.T) {
	fmt.Println("Test_HandleName")
}

func Test_HandleAge(t *testing.T) {
	fmt.Println("Test_HandleAge")
}

func Test_HandleSex(t *testing.T) {
	fmt.Println("Test_HandleSex")
	assert.Equal(t, "1", "2")
}

