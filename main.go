package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pablotrianda/porfolioTui/src/constants"
)

func porfolio(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	if strings.Contains(strings.ToLower(userAgent), "curl") {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, constants.Page())
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data, err := os.ReadFile("public/index.html")
		if err != nil {
			http.Error(w, "Error loading page", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

func main() {

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/curl.txt", fs)

	http.HandleFunc("/", porfolio)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	log.Printf("Server runnig at :%s port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("Error when try to run the server")
	}
}
