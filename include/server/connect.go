// Interacts with the server you can host (with host.go). Pushing, pulling notebooks and
// doing health checks.

package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "encoding/json"
  "net/http"
  "errors"
  "bytes"
)

func formatUrl(path string) string {
  return NtConfig.Server.Url + NtConfig.Server.Port + "/" + path
}

func promptAuth() (string, string) {
  var username []byte = InsecureInput("[1/2] Set username for nts: ")
  var password []byte = SecureInput("[2/2] Set password for nts: ")
  return string(username), string(password)
}

func PushNotebook(notebook Notebook) error {
  // check prerequisites
  if len(NtConfig.Server.Url) == 0 { return errors.New("No URL in config.") }
  jsonData, jsonErr := json.Marshal(notebook)
  if jsonErr != nil { return jsonErr }

  // create request
  req, requestErr := http.NewRequest("POST", formatUrl("push"), bytes.NewBuffer(jsonData))
  if requestErr != nil || req == nil { return requestErr }
  req.Header.Set("Content-Type", "application/json")
  username, password := promptAuth()
  req.SetBasicAuth(username, password)
  client := &http.Client{}

  // execute the request and handle response
  resp, responseErr := client.Do(req)
  if resp == nil {
    Error.Println("No internet connection.")
  } else if responseErr != nil || resp.StatusCode != 200 {
    Error.Printf("%s", resp.Status)
  } else {
    Info.Printf("[%d] Pushed %d notes to server.", resp.StatusCode, len(notebook.Notes))
  }
  return responseErr
}

func PullNotebook() (Notebook, error) {
  // check prerequisites
  var notebook Notebook
  var writeErr error
  if len(NtConfig.Server.Url) == 0 {
    return notebook, errors.New("No URL provided in config.")
  }

  // create request
  client := &http.Client{}
  req, _ := http.NewRequest("GET", formatUrl("pull"), nil)
  username, password := promptAuth()
  req.SetBasicAuth(username, password)

  // execute the request and handle response
  resp, requestErr := client.Do(req)
  if requestErr != nil { return notebook, requestErr }
  if resp.StatusCode != http.StatusOK {
    Info.Println(resp.StatusCode, resp)
    return notebook, errors.New("Status code implies failed request.")
  }
  jsonErr := json.NewDecoder(resp.Body).Decode(&notebook)
  if jsonErr == nil {
    Info.Printf("Pulled %d notes from server. %s.", len(notebook.Notes), resp.Status)
    writeErr = WriteNotebook(notebook)
  }
  return notebook, errors.Join(jsonErr, writeErr)
}

func PingServer() error {
  // check prerequisites
  if len(NtConfig.Server.Url) == 0 {
    return errors.New("No URL provided in config.")
  }

  // create request
  req, requestErr := http.NewRequest("POST", formatUrl("ping"), nil)
  client := &http.Client{}

  // execute the request and handle response
  resp, responseErr := client.Do(req)
  if resp == nil {
    Error.Println("No internet connection.")
  } else {
    Info.Printf("%s.", resp.Status)
  }
  return errors.Join(responseErr, requestErr);
}

