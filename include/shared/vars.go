// Just a submodule to put all global vars so I can find them later on. Nt prefix means
// it's created on setup.

package include

// loaded by fs
var NtPath string
var NtConfig Config
var NtNotes Notebook
var NtNotesErr error
var NtPathErr error
var NtConfigErr error
var DefaultConfig Config = Config{
  Server: ServerConfig {
    Url: "",
    Port: ":8282",
  },
  Notebook: NotebookConfig {
    Width: 30,
    DateFormats: []string{"2006-01-02T15:04", "2006-01-02", "Jan 02", "2", "Mon"},
    LsDefault: "--all",
  },
}

// colors
var Reset = "\033[0m"
var Red = "\033[31m"
var Yellow = "\033[33m"
var Magenta = "\033[35m"
var Cyan = "\033[34m"

// other
var TestMode bool = false
var HelpString string = `Command not found.
  Valid commands :)
  `
