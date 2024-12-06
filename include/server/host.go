
package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "crypto/sha256"
  "crypto/subtle"
  "encoding/json"
  "net/http"
  "errors"
  "fmt"
)

var auth Authentication

func configauth() error {
  var username []byte = InsecureInput("[1/3] Set username for nts: ")
  var password []byte = SecureInput("[2/3] Set password for nts: ")
  var passwordCheck []byte = SecureInput("[3/3] Repeat password:      ")
  if string(password) != string(passwordCheck) {
    return errors.New("Passwords don't match.")
  }
  auth.Username = username
  auth.Password = password
  return nil
}

func okAuth(username string, password string) bool {
  usernameHash := sha256.Sum256([]byte(username))
  passwordHash := sha256.Sum256([]byte(password))
  expectedUsernameHash := sha256.Sum256(auth.Username)
  expectedPasswordHash := sha256.Sum256(auth.Password)
  usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
  passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)
  return usernameMatch && passwordMatch
}

func basicAuth(next http.Handler) http.HandlerFunc {
  var username string
  var password string
  var ok bool // used to see if basicAuth works
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    username, password, ok = r.BasicAuth()
    if ok {
      if okAuth(username, password) {
        next.ServeHTTP(w, r)
        return
      }
    }
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
  })
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
  var notebook Notebook
  var writeErr error
  var decoder *json.Decoder = json.NewDecoder(r.Body)
  var decodeErr error = decoder.Decode(&notebook)
  if decodeErr != nil {
    http.Error(w, "Invalid json.", http.StatusBadRequest); return
  }
  if len(notebook.Notes) == 0 {
    http.Error(w, "No notes found in payload.", http.StatusBadRequest); return
  }
  writeErr = WriteNotebook(notebook)
  if writeErr != nil {
    Error.Println(writeErr)
  } else {
    fmt.Fprintln(w, "Succesfully added", len(notebook.Notes), "notes.")
  }
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "pong")
}

func pullHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  if err := json.NewEncoder(w).Encode(NtNotes); err != nil {
    http.Error(w, "Error encoding json.", http.StatusInternalServerError)
  }
}

func RunServer() error {
  var serverErr error
  if configErr := configauth(); configErr != nil { return configErr }
  http.Handle("/push", basicAuth(http.HandlerFunc(pushHandler)))
  http.Handle("/pull", basicAuth(http.HandlerFunc(pullHandler)))
  http.Handle("/ping", http.HandlerFunc(healthHandler))
  Warn.Printf("server started at %s", NtConfig.Server.Port)
  serverErr = http.ListenAndServe(NtConfig.Server.Port, nil)
  return serverErr
}
