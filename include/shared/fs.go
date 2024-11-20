package include

import (
  toml "github.com/pelletier/go-toml/v2"

  "encoding/json"
  "errors"
  "os"
)

var NtPath string
var NtPathErr error
var NtConfig Config
var NtConfigErr error

func NtDir() (string, error) {
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

func WriteNotebook(notebook Notebook) error {
  jsonData, jsonErr := json.Marshal(&notebook)
  writeErr := os.WriteFile(NtPath + "notebook.json", jsonData, 0644)
  if err := errors.Join(jsonErr, writeErr); err != nil {
    return err
  }
  return nil
}

func LoadNotebook() (Notebook, error) {
  var notebook Notebook
  jsonFile, fileErr := os.ReadFile(NtPath + "notebook.json")
  if fileErr == nil {
    jsonErr := json.Unmarshal(jsonFile, &notebook)
    return notebook, jsonErr
  } else if errors.Is(fileErr, os.ErrNotExist) {
    Warn.Println("No notebook found. Will create new file on save.")
    return notebook, nil
  }
  return notebook, fileErr
}

func defaultConfig() Config {
  return Config {
    Server: ServerConfig {
      Url: "",
      Port: ":8282",
    },
    Notebook: NotebookConfig {
      Width: 75,
      DateFormats: []string{"2006-01-02T15:04", "2006-01-02", "Jan 02", "2", "Mon"},
    },
  }
}

func LoadConfig() (Config, error) {
  NtConfig = defaultConfig()
  tomlFile, fileErr := os.ReadFile(NtPath + "config.toml")
  if fileErr == nil {
    tomlErr := toml.Unmarshal(tomlFile, &NtConfig)
    return NtConfig, tomlErr
  } else if errors.Is(fileErr, os.ErrNotExist) {
    return NtConfig, nil
  }
  return NtConfig, fileErr
}

func init() {
  NtPath, NtPathErr = NtDir()
  NtConfig, NtConfigErr = LoadConfig()
}

