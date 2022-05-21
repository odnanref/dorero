package main
 
import (
"fmt"
"net/http"
"html/template"
)

type Page struct {
	Title string
	Message string
	RemoteIP string
}

func main() {
    http.HandleFunc("/", indexPage)
    http.ListenAndServe(":8000", nil)
	fmt.Println("... Started...")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	var templates *template.Template	
	templates = template.Must(templates.ParseGlob("templates/*.html"))
	
	var remoteIP = r.RemoteAddr

	p := &Page{
		Title: "Monkey Business",
		Message: "Hello!",
		RemoteIP: remoteIP,
	}

	fmt.Println("... Started: " + r.RemoteAddr)
	
	templates.ExecuteTemplate(w, "index.html", p)
}

