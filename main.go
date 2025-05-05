package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("[アクセスログ] %s %s\n", r.Method, r.URL.Path)
        next(w, r)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprint(w, `<form method="POST">
            name: <input name="name">
            <input type="submit">
        </form>`)
    } else if r.Method == http.MethodPost {
        r.ParseForm()
        name := r.FormValue("name")
        fmt.Fprintf(w, "こんにちは、%sさん！", name)
    }
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{"message": "これはJSONレスポンスです"}
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/", loggingMiddleware(helloHandler))
    http.HandleFunc("/form", loggingMiddleware(formHandler))
    http.HandleFunc("/json", loggingMiddleware(jsonHandler))

    fmt.Println("サーバーを http://localhost:8080 で起動中...")
    http.ListenAndServe(":8080", nil)
}
