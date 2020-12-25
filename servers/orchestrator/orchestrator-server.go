package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	texts := []string{
		"somehow got this text",
		"this is another text",
		"jinx is an adc",
		"shen is a tank",
		"something something",
		"this is a text of text",
		"i like playing shen",
	}

	c := make(chan string)

	for _, text := range texts {
		go callServer(text, c)
	}

	for l := range c {
		go func(text string) {
			time.Sleep(3 * time.Second)
			callServer(text, c)
		}(l)
	}

	return
}
