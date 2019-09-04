package main

import (
		"net/http"
		"fmt"
		"log"
)

func hello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	switch r.Method{
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		if err := r.ParseForm(); err != nil{
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w,"Post from website! r.PostForm = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s \n", address)
	default:
		fmt.Fprintf(w, "Another method doesn't exists! ")

	}
}

func main()  {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST ..\n")
	log.Fatal(http.ListenAndServe(":8888", nil))
}