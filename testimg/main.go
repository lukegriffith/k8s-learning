package main
import (
  "net/http"
  "strings"
  "fmt"
)


func main() {

  fmt.Println("server started on port 8080")
  http.HandleFunc("/", greet)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}


func greet(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello robot " + message
  w.Write([]byte(message))
}


