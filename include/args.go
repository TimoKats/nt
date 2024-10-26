package include

func ParseArgs(arguments []string) Arguments {
  var parsedArgs Arguments
  for _, argument := range arguments {
    switch {
    case argument == "-w" || argument == "w":
      parsedArgs.Command = "write"
      parsedArgs.CurrentType = Write
    case argument == "-r" || argument == "r":
      parsedArgs.Command = "read"
      parsedArgs.CurrentType = Read
    case argument == "-t" || argument == "t":
      parsedArgs.CurrentType = Tags
    case parsedArgs.CurrentType == Write:
      parsedArgs.Text += argument + " "
    case parsedArgs.CurrentType == Tags:
      parsedArgs.Tags = append(parsedArgs.Tags, argument)
    }
  }
  return parsedArgs
}

