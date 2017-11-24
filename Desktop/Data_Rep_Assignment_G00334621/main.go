// G00334621
// Data Representation Project 2017
// Christian Olim
// https://github.com/data-representation/eliza/blob/master/eliza.go

package main

import (
	"fmt"
	"net/http"
	"./static"
)
func main() {
	// Here is our file server 
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/askEliza", HandleAsk)
	// This begins our server for our website
	http.ListenAndServe(":8080", nil)
}

func HandleAsk(writer http.ResponseWriter, request *http.Request) {
	// Here is code which runs when the user makes a request
	userInput := request.URL.Query().Get("userInput")
	answer := eliza.AskEliza(userInput)
	// This writes back the users input to the ResponseWriter
	fmt.Fprintf(writer, answer)

}

