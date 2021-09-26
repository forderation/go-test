package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GreetIO(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHandler(rw http.ResponseWriter, r *http.Request) {
	GreetIO(rw, "Again")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler)))
	// var strBuffer bytes.Buffer
	// strBuffer.WriteString("Ranjan ")
	// strBuffer.WriteString("Kumar")
	// fmt.Println("The string buffer output is", strBuffer.String())
	// fmt.Println("Len buffer", strBuffer.Len())
}
