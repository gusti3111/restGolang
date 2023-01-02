package restgolang

import (
	"io"
	"log"
	"net/http"
	"testing"
)

func Myserver(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello\n")

}

func TestServer(t *testing.T) {
	http.HandleFunc("/hello", Myserver)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
