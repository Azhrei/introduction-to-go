package main

import (
	"html/template"
	"log"
	"net/http"
)

// START OMIT
const templ = `<html><head><title>Hello</title></head><body>
Hello {{ .RemoteAddr }}
You sent me a {{ .Method }} request for {{ .URL }}
</body></html>`

// END OMIT
func HelloServer(w http.ResponseWriter, req *http.Request) {
    // ...
// START OMIT
	log.Println(req.URL)
	t := template.Must(template.New("html").Parse(templ))
	t.Execute(w, req)
// END OMIT
}

func main() {
	log.Println("please connect to http://localhost:7777/")
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
