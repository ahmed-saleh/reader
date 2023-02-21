package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
//	"strings"
)

type Body struct {
	Name  string `json:"name"`
	Description string `json:"description"`
	Event_type string `json:"event_type"`
	Origin string`json:"origin"`
	Timestamp time.Time`json:"timestamp"`
	Result string`json:"result"`
}

/*
func (b *Body) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse(time.RFC1123, strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	b.Timestamp = tt
	return nil
}
*/
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
		fmt.Println(data, time.Now().UTC())
		messages = append(messages, data)

	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
}

func printer(w http.ResponseWriter, r *http.Request) {

	fmt.Println("printing ....")
	file, _ := json.MarshalIndent(messages, "", " ")
	filename, _ := time.Now().UTC().MarshalText()
 
	if err := ioutil.WriteFile("output/report "+ string(filename) + ".json", file, 0644); err != nil {
		fmt.Println("failed to write into file")
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
