
package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "fmt"
  "net/http"
  "crypto/sha256"
  "crypto/subtle"
  "encoding/json"
)

func basicAuth(next http.Handler) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    username, password, ok := r.BasicAuth()
    if ok {
      usernameHash := sha256.Sum256([]byte(username))
      passwordHash := sha256.Sum256([]byte(password))
      expectedUsernameHash := sha256.Sum256([]byte("username"))
      expectedPasswordHash := sha256.Sum256([]byte("password"))
      usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
      passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)
      if usernameMatch && passwordMatch {
        next.ServeHTTP(w, r)
        return
      }
    }
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
  })
}

func putHandler(w http.ResponseWriter, r *http.Request) {
  var notebook Notebook
  decoder := json.NewDecoder(r.Body)
  decodeErr := decoder.Decode(&notebook)
  if decodeErr != nil { http.Error(w, "Invalid json.", http.StatusBadRequest)
    return
  }
  if len(notebook.Notes) == 0 {
    http.Error(w, "No notes found in payload.", http.StatusBadRequest)
    return
  }
  WriteNotebook(notebook)
  fmt.Fprintln(w, "Succesfully added", len(notebook.Notes), "notes.")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "pong")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  notebook, loadErr := LoadNotebook()
  if loadErr != nil {
    http.Error(w, "Error loading json from notebook.", http.StatusInternalServerError)
  }
  if err := json.NewEncoder(w).Encode(notebook); err != nil {
    http.Error(w, "Error encoding json.", http.StatusInternalServerError)
  }
}

func RunServer() {
  http.Handle("/push", basicAuth(http.HandlerFunc(putHandler)))
  http.Handle("/pull", basicAuth(http.HandlerFunc(getHandler)))
  http.Handle("/ping", http.HandlerFunc(healthHandler))
  Warn.Println("server started at :8080")
  http.ListenAndServe(":8080", nil)
}

