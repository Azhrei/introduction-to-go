package main

import (
	"html/template"
	"log"
	"net/http"
)

const templ = `<html><head><title>Hello</title></head><body>
Hello {{ .RemoteAddr }}
You sent me a {{ .Method }} request for {{ .URL }}
</body></html>`

func HelloServer(w http.ResponseWriter, req *http.Request) {
    // ...
    log.Println(req.URL) // OMIT
    t := template.Must(template.New("html").Parse(templ)) // OMIT
    t.Execute(w, req) // OMIT
}

func main() {
    log.Println("Please visit http://localhost:7777/")
    http.HandleFunc("/", HelloServer)
    log.Fatal(http.ListenAndServe(":7777", nil))
}
