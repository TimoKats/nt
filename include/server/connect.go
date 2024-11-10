package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "encoding/json"
  "net/http"
  "errors"
  "bytes"
)

func PushNotebook(notebook Notebook) error {
  jsonData, jsonErr := json.Marshal(notebook)
  req, requestErr := http.NewRequest("POST", "http://localhost:8080/put", bytes.NewBuffer(jsonData))
  if err := errors.Join(jsonErr, requestErr); err != nil {
    return err
  }
  req.Header.Set("Content-Type", "application/json")
  req.SetBasicAuth("username", "password")
  client := &http.Client{}
  resp, responseErr := client.Do(req)
  Info.Println(resp)
  return responseErr
}

func PullNotebook() (Notebook, error) {
  var notebook Notebook
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "http://localhost:8080/get", nil)
  req.SetBasicAuth("username", "password")
  resp, requestErr := client.Do(req)
  if requestErr != nil {
    return notebook, requestErr
  }
  if resp.StatusCode != http.StatusOK {
    Info.Println(resp.StatusCode, resp)
    return notebook, errors.New("Status code implies failed request.")
  }
  jsonErr := json.NewDecoder(resp.Body).Decode(&notebook)
  if jsonErr == nil {
    WriteNotebook(notebook)
  }
  return notebook, jsonErr
}
