package include

func GetCommand(argument string) CommandType {
  switch argument {
  case "add", "a", "-a":
    return Add
  case "get", "g", "-g":
    return Get
  case "clear", "c", "-c":
    return Clear
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
      parsedArgs.Text += argument + " "
    }
  }
  return parsedArgs
}

