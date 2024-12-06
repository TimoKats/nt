package include

import (
  "github.com/atotto/clipboard"

  "strings"
  "strconv"
  "errors"
)

func readClipboard() string {
  var clipboardText string
  var readErr error
  clipboardText, readErr = clipboard.ReadAll()
  if readErr != nil { return "" }
  return string(clipboardText)
}

func isFlag(argument string) bool {
  if len(argument) > 2 {
    return string(argument[0:2]) == "--"
  }
  return false
}

func parseNoteIds(argument string) []int {
  var noteIds []int
  if len(argument) == 1 {
    noteId, convErr := strconv.Atoi(argument)
    if convErr != nil { return noteIds }
    noteIds = []int{ noteId }
  } else if strings.Contains(argument, "-") && len(argument) == 3 {
    startId, convErr1 := strconv.Atoi(string(argument[0]))
    endId, convErr2 := strconv.Atoi(string(argument[2]))
    if err := errors.Join(convErr1, convErr2); err != nil { return noteIds }
    for i := startId; i <= endId; i++ {
      noteIds = append(noteIds, i)
    }
  }
  return noteIds
}

func getCommand(argument string) CommandType {
  switch argument {
    // notebook
    case "add", "a":
      return Add
    case "ls", "l":
      return List
    case "remove", "rm":
      return Clear
    case "move", "mv":
      return Move
    case "comment", "cmt":
      return Comment
    case "tags", "tag":
      return Tags
    case "modify", "mod":
      return Modify
    case "s", "search":
      return Search
    // server
    case "run":
      return Server
    case "pull":
      return Pull
    case "push":
      return Push
    case "ping":
      return Ping
    default:
      return None
  }
}

func enrichArgs(argument string, parsedArgs *Arguments) {
  if string(argument[0]) == "@" {
    parsedArgs.Tags = append(parsedArgs.Tags, argument)
  } else if argument == ":c" {
    parsedArgs.Text += readClipboard()
  } else if strings.HasPrefix(argument, ":due:") {
    parsedArgs.Deadline = ParseDate(argument)
  } else if isFlag(argument) {
    parsedArgs.Flags = append(parsedArgs.Flags, argument)
  } else {
    parsedArgs.Text += argument + " "
  }
}

func ParseArgs(arguments []string) Arguments {
  var parsedArgs Arguments
  for index, multiArgument := range arguments { // sometimes "" protect args on windows
    for _, argument := range strings.Fields(multiArgument) {
      if index == 1 { parsedArgs.Command = getCommand(argument) }
      if index > 1 { enrichArgs(argument, &parsedArgs) }
      if index == 2 { parsedArgs.NoteIds = parseNoteIds(argument) }
    }
  }
  if len(parsedArgs.Flags) == 0 && len(parsedArgs.NoteIds) == 0 {
    parsedArgs.Flags = []string{ NtConfig.Notebook.LsDefault }
  }
  return parsedArgs
}

