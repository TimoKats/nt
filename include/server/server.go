
package include

import (
  "fmt"
  "net/http"
  "crypto/sha256"
  "crypto/subtle"
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

func protectedHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "You have access to the protected resource!")
}

func RunServer() {
  http.Handle("/put", basicAuth(http.HandlerFunc(protectedHandler)))
  fmt.Println("Server started at :8080")
  http.ListenAndServe(":8080", nil)
}
