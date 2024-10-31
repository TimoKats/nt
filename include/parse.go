package include

import (
  "golang.design/x/clipboard"
)

func readClipboard() string {
  initErr := clipboard.Init()
  if initErr != nil {
    return ""
  }
  clipboardText := clipboard.Read(clipboard.FmtText)
  return string(clipboardText)
}

func GetCommand(argument string) CommandType {
  switch argument {
  case "add", "a":
    return Add
  case "ls", "l":
    return List
  case "clear", "c":
    return Clear
  case "move", "mv":
    return Move
  default:
    return None
  }
}

func ParseArgs(arguments []string) Arguments {
  var parsedArgs Arguments
  for index, argument := range arguments {
    if index == 1 {
      parsedArgs.Command = GetCommand(argument)
    } else if index > 1 {
      if string(argument[0]) == "@" {
        parsedArgs.Tags = append(parsedArgs.Tags, argument)
      }
      if string(argument) == "*c" {
        parsedArgs.Text += readClipboard()
      } else {
        parsedArgs.Text += argument + " "
      }
    }
  }
  return parsedArgs
}

