package webserver

import (
	"fmt"
	"net/http"
	"time"
)


// Serve is for init the simple hello world in a web server
func Serve(port string){
	http.HandleFunc("/", handleHelloWorld)
	http.ListenAndServe(port, nil)
}


func handleHelloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}