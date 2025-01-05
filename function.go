package f

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

// GCF request message
type Message struct {
	Body string `json:"body"`
}

// GCF entrypoint
func Trigger(response http.ResponseWriter, request *http.Request) {
	var message Message
	if err := json.NewDecoder(request.Body).Decode(&message); err != nil {
		log.Printf("Error: %v", err)
		http.Error(response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		log.Printf("Request received: %v", message)

		Service(message)

		responseBody := "Finish!"
		log.Println(responseBody)
		fmt.Fprint(response, html.EscapeString(responseBody))
	}
}
