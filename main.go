package main

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
  "path"
  "log"
  "github.com/PuerkitoBio/goquery"
)


func handle(w http.ResponseWriter, r *http.Request) {
tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  Subscriber := ""
  Url := "https://www.youtube.com/c/tasarimci-dayi"
	name := ""
  
  	if r.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(r.URL.Path)
	}
  
	  switch r.Method{
    case "GET":
     name="index.html"
    case "POST":
     Url=r.FormValue("username")
  }


	data := struct{
		Time time.Time
    Url string
    Subscriber string
	}{
		Time: time.Now(),
    Url: Url,
    Subscriber: Subscriber,
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
    }
  
  res, err := http.Get(Url)
  if err != nil {
    log.Fatal(err)
  } 
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }
  doc, err := goquery.NewDocumentFromReader(res.Body)

  Subscriber = doc.Find("#subscriber-count").Text()
		fmt.Printf("Review : ", Subscriber)
}

func main() {

	fmt.Println("http server up!")
	http.Handle(
		"/static/",
		 http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":0", nil)
}

