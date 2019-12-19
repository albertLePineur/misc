package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func random() string {
	b := make([]byte, 55)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}

// Test function
func Test(affichage http.ResponseWriter, requete *http.Request) {
	var a int = 10
	io.WriteString(affichage, "coucou les lapinous !!!")
}

func main() {
	fmt.Println("coucou")
	http.HandleFunc("/test", Test)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
