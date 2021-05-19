package main

import (
	"fmt"
	"log"
	"net/http"
)

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	http.Redirect(w, r, "/contact.html", 200)
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	http.Redirect(w, r, "/aboutMe.html", 303)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	http.Redirect(w, r, "/index.html", 303)
}

func webscraper(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	http.Redirect(w, r, "/webscraper.html", 303)
}

func downloadImages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	http.Redirect(w, r, "/webscraper.html", 303)
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./HTML")))
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/aboutMe", aboutMe)
	http.HandleFunc("/index", index)
	http.HandleFunc("/webscraper", webscraper)
	http.HandleFunc("/downloadImages", downloadImages)
	err := http.ListenAndServe(":80", nil) // set listen port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
