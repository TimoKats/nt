package include

import (
  "github.com/atotto/clipboard"

  "strings"
  "strconv"
)

func readClipboard() string {
  clipboardText, readErr := clipboard.ReadAll()
  if readErr != nil {
    return ""
  }
  return string(clipboardText)
}

func GetCommand(argument string) CommandType {
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

func isFlag(argument string) bool {
  if len(argument) > 2 {
    return string(argument[0:2]) == "--"
  }
  return false
}

func ParseArgs(arguments []string) Arguments {
  var parsedArgs Arguments
  // set some defaults
  parsedArgs.NoteId = -1
  for index, multiArgument := range arguments { // sometimes "" protect args on windows
    for _, argument := range strings.Fields(multiArgument) {
      if index == 1 {
        parsedArgs.Command = GetCommand(argument)
      } else if index > 1 {
        if string(argument[0]) == "@" {
          parsedArgs.Tags = append(parsedArgs.Tags, argument)
        } else if string(argument) == "*c" {
          parsedArgs.Text += readClipboard()
        } else if isFlag(argument) {
          parsedArgs.Flags = append(parsedArgs.Flags, argument)
        } else {
          parsedArgs.Text += argument + " "
        }
      }
      if index == 2 {
        noteId, convErr := strconv.Atoi(argument)
        if convErr == nil {
          parsedArgs.NoteId = noteId
        }
      }
    }
  }
  return parsedArgs
}

