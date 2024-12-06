package include

import (
  toml "github.com/pelletier/go-toml/v2"

  "encoding/json"
  "errors"
  "sync"
  "os"
)

// nt home dir (~/.nt)

func SetNtDir() {
  homedir, dirErr := os.UserHomeDir()
  if dirErr != nil { NtPathErr = dirErr }
  if _, fileErr := os.Stat(homedir + "/.nt/"); os.IsNotExist(fileErr) {
    NtPathErr = os.MkdirAll(homedir + "/.nt/", 0755)
  }
  NtPath = homedir + "/.nt/"
}

// notebook

func WriteNotebook(notebook Notebook) error {
  jsonData, jsonErr := json.Marshal(&notebook)
  writeErr := os.WriteFile(NtPath + "notebook.json", jsonData, 0644)
  if err := errors.Join(jsonErr, writeErr); err != nil {
    return err
  }
  return nil
}

func LoadNotebook(wg *sync.WaitGroup) {
  defer wg.Done()
  jsonFile, fileErr := os.ReadFile(NtPath + "notebook.json")
  if fileErr == nil {
    NtNotesErr = json.Unmarshal(jsonFile, &NtNotes)
  } else if errors.Is(fileErr, os.ErrNotExist) {
    Warn.Println("No notebook found. Will create new file on save.")
  }
}

// config

func defaultConfig() Config {
  return Config {
    Server: ServerConfig {
      Url: "",
      Port: ":8282",
    },
    Notebook: NotebookConfig {
      Width: 75,
      DateFormats: []string{"2006-01-02T15:04", "2006-01-02", "Jan 02", "2", "Mon"},
      LsDefault: "--all",
    },
  }
}

func LoadConfig(wg *sync.WaitGroup) {
  defer wg.Done()
  NtConfig = defaultConfig()
  tomlFile, fileErr := os.ReadFile(NtPath + "config.toml")
  if fileErr == nil {
    tomlErr := toml.Unmarshal(tomlFile, &NtConfig)
    NtConfigErr = tomlErr
  }
  NtConfigErr = fileErr
}

