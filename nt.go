package main

import (
  notebook "github.com/TimoKats/nt/include/notebook"
  server "github.com/TimoKats/nt/include/server"
  . "github.com/TimoKats/nt/include/shared"

  "errors"
  "os"
)

func run(arguments Arguments) error {
  switch arguments.Command {

  // notebook related
  case Add:
    return notebook.AddNote(arguments)
  case List:
    return notebook.ReadNotebook(arguments)
  case Clear:
    return notebook.ClearNotebook(arguments)
  case Move:
    return notebook.MoveNote(arguments)
  case Comment:
    return notebook.AddComment(arguments)
  case Tags:
    return notebook.ReadTags(arguments)

  // server related
  case Server:
    return server.RunServer()
  case Push:
    return server.PushNotebook(notebook.Notes)
  case Pull:
    var pullErr error
    notebook.Notes, pullErr = server.PullNotebook()
    return pullErr
  default:
    return errors.New("No valid command found. Use <<ls, add, mv>>")
  }
}

func InitErrs() bool {
  return errors.Join(NtPathErr, NtConfigErr) != nil
}

func main() {
  if initErr := errors.Join(NtPathErr, NtConfigErr); initErr != nil {
    Error.Println(initErr)
    return
  }
  arguments := ParseArgs(os.Args)
  runErr := run(arguments)
  if runErr != nil {
    Error.Println(runErr)
  }
}
