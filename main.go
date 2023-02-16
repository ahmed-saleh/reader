package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Body struct {
	Name  string `json:"name"`
	Description string `json:"description"`
	Event_type string `json:"event_type"`
	Origin string`json:"origin"`
	Timestamp int64`json:"timestamp"`
	Result string`json:"result"`
}

var messages = []Body {}

func Parse(s []byte) Body {
	data := Body{}
	json.Unmarshal(s, &data)
	return data
}

func collect(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var data = Parse(body)
		messages = append(messages, data)


	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
}

func printer() {

	file, _ := json.MarshalIndent(messages, "", " ")
 
	if err := ioutil.WriteFile("test.json", file, 0644); err != nil {
		fmt.Errorf("failed to write into file")
	} else {
		//consume
		messages = []Body{}
	}
}

func main() {
	http.HandleFunc("/collect", collect)    // Update this line of code
	http.HandleFunc("/print", printer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
