package webserver

import (
	"log"
	"fmt"
	"net/http"
	"time"
)


// Serve is for init the simple hello world in a web server
func Serve(port string){
	http.HandleFunc("/", handleHelloWorld)
	log.Printf("Listen on port %s", port)
	http.ListenAndServe(port, nil)
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request)  {
	log.Printf("%v, Request to: %v, Host: %v, Remote Address: %v", r.Method, r.RequestURI, r.Host, r.RemoteAddr)
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}