package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	port := "8080"

	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	svc := os.Getenv("REDIS_CLIENT")
	if svc == "" {
		svc = "Not set"
	}

	svcUrl := fmt.Sprintf("http://%s", svc)
	resp, err := http.Get(svcUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	log.Println(sb)
	fmt.Fprintf(w, sb)
}