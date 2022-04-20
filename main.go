package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter,r * http.Request){
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"Parse From err :%v",err)
		return
	}
	fmt.Fprintf(w,"Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)



}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	port := 4040
	fmt.Printf("Starting server at port %v\n", port)
	fmt.Printf("go to the link: http://localhost:%v", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}

}
