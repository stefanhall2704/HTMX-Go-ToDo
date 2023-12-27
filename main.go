
package main

import (
    "html/template"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", logRequest(serveForm))
    http.HandleFunc("/create_task", logRequest(formHandler))

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func logRequest(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
        handler(w, r)
    }
}

func serveForm(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("form.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing the form", http.StatusInternalServerError)
        return
    }

    var tmpl *template.Template
    var err error
    var task string
    var data struct {
        Task string
    }

    if r.Method == "POST" {
        task = r.FormValue("task")
        tmpl, err = template.ParseFiles("response_template.html")
        data.Task = task
    } else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
