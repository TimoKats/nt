package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "encoding/json"
  "errors"
  "os"
)

var CONFIGPATH string
var CONFIGERR error

func ConfigDir() (string, error) {
  homedir, dirErr := os.UserHomeDir()
  var createErr error = nil
  if dirErr != nil {
    return "", dirErr
  }
  if _, fileErr := os.Stat(homedir + "/.nt/"); os.IsNotExist(fileErr) {
    createErr = os.MkdirAll(homedir + "/.nt/", 0755)
  }
  return homedir + "/.nt/", createErr
}

func WriteNotebook() error {
  jsonData, jsonErr := json.Marshal(&notebook)
  writeErr := os.WriteFile(CONFIGPATH + "notebook.json", jsonData, 0644)
  if err := errors.Join(jsonErr, writeErr); err != nil {
    return err
  }
  return nil
}

func LoadNotebook() error {
  jsonFile, fileErr := os.ReadFile(CONFIGPATH + "notebook.json")
  if fileErr == nil {
    jsonErr := json.Unmarshal(jsonFile, &notebook)
    if jsonErr == nil {
      return nil
    }
    return jsonErr
  } else if errors.Is(fileErr, os.ErrNotExist) {
    Warn.Println("No notebook found. Will create new file on save.")
    return nil
  }
  return fileErr
}

func init() {
  CONFIGPATH, CONFIGERR = ConfigDir()
}

