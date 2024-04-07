package dep_inject

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// io.writer is interface
// which is implemented by other functions
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// in terminal os.Stdout implement io.writer
// so it will output to terminal by it
func Terminal() {
	Greet(os.Stdout, "Alnoor")
}

// in net/http module ResponseWriter function
// implement io.writer interface
// so when open page you will see the same output from Greet
func greetHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name, ok := queryParams["name"]
	if !ok {
		name[0] = "World"
	}
	Greet(w, name[0])
}

// start http server
func Web() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(greetHandler)))
}
