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

  Valid notebook commands:
  - add <<text>>: Adds note to nt.
  - ls Lists the notes in nt.
  - rm Removes notes.
  - cmt <<id>>: Adds comment to specific note.
  - tags Lists the tags in the notebook.
  - mod <<id>>: Modifies a selected note. Same args as 'add'.
  - mv <<id>>: Checks/Unchecks a note.
  - s <<query>>: Filters the notebook on a query.

  Valid server commands:
  - run: Creates prompt that starts server.
  - push: Pushes notes to server.
  - pull: Pulls notes from server.
  - ping: Checks if server in config is running.
  `

