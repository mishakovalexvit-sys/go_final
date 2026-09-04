package main

import (
	"net/http"
)


func main(){
	port := "7540"
	http.Handle("/", http.FileServer(http.Dir("web")))
	err := http.ListenAndServe(":" + port, nil)
	if err != nil{
		return
	}
}