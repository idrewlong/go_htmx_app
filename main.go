// net/http module
// html/template module

// create web server

// https://pkg.go.dev/net/http
// https://pkg.go.dev/html/template
// https://golangforall.com/en/post/templates.html

// Notes:
// w is an 'http.ResponseWriter' object. Used to write the HTTP response back to the client. Used to set headers, write response data and send the response status code.

// r is an http.Request object. Represents the http request received from the client. Object contains information about the request, such as the HTTP method, URL, headers and any data sent from the request.

// io.WriteString(w, "Hello, world!") writes the string hello world to the http.ResponseWriter object w, sending it as part of the HTTP response body.

// io.WriteString(w, r.Method) write the HTTP method(GET, POST, etc) of the incoming request r to the http.ResponseWriter object w, sending it as part of the HTTP response body

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// func main() {

// 	h1 := func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Hello World!")
// 		io.WriteString(w, r.Method)
// 	}
// 	http.HandleFunc("/", h1)
// 	log.Fatal(http.ListenAndServe("8000", nil))
// 	starts an HTTP server on port 8000 and logs any errors that occur, causing the program to exit if an error is encountered.
// }

type Music struct {
	Song string
	Artist string
}

func main() {
	fmt.Println("Go App...")


	h1 := func(w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("index.html"))
		songs := map[string][]Music{
			"Songs" : {
				{Song: "What a Pleasure", Artist: "Beach Fossils"},
				{Song: "Taker", Artist: "DIIV"},
				{Song: "Pain", Artist: "The War on Drugs"},
			},
		}
		tmpl.Execute(w, songs)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		song := r.PostFormValue("title")
		artist := r.PostFormValue("artist")
	// 	htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", song, artist)
	//       tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl := template.Must(template.ParseFiles("index.html"))
	       tmpl.ExecuteTemplate(w, "song-list-element", Music{Song: song, Artist: artist})
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-song/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}