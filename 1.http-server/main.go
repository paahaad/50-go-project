package main

import (
	"fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method in not allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "HELLO")
}
func main(){
	fmt.Println("Welcome to basic http-web-server")
	fileServer := http.FileServer(http.Dir("./static")) 

	http.Handle("/", fileServer) 
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Application is running at 8080")
	if err:= http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}