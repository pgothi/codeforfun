package main

import (
    "log"
    "net/http"
    "io"
)

func main() {

    http.HandleFunc("/post", HandlePost)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
func HandlePost(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "done\n")
}
