package main

import "fmt"
import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

func main(){
	fmt.Println("hello")
	funcAsVariable := func (k string) {}
	funcAsVariable("hello")
}

H