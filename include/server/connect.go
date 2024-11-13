package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "encoding/json"
  "net/http"
  "strconv"
  "errors"
  "bytes"
)

func PushNotebook(notebook Notebook) error {
  if len(NtConfig.Server.Url) == 0 {
    return errors.New("No URL provided in config.")
  }
  jsonData, jsonErr := json.Marshal(notebook)
  url := NtConfig.Server.Url + ":" + strconv.Itoa(NtConfig.Server.Port)
  req, requestErr := http.NewRequest("POST", url + "/push", bytes.NewBuffer(jsonData))
  if err := errors.Join(jsonErr, requestErr); err != nil {
    return err
  }
  req.Header.Set("Content-Type", "application/json")
  req.SetBasicAuth("username", "password")
  client := &http.Client{}
  resp, responseErr := client.Do(req)
  Info.Printf("Pushed %d notes to server. %s.", len(notebook.Notes), resp.Status)
  return responseErr
}

func PullNotebook() (Notebook, error) {
  var notebook Notebook
  if len(NtConfig.Server.Url) == 0 {
    return notebook, errors.New("No URL provided in config.")
  }
  url := NtConfig.Server.Url + ":" + strconv.Itoa(NtConfig.Server.Port)
  client := &http.Client{}
  req, _ := http.NewRequest("GET", url + "/pull", nil)
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
    Info.Printf("Pulled %d notes from server. %s.", len(notebook.Notes), resp.Status)
    WriteNotebook(notebook)
  }
  return notebook, jsonErr
}

