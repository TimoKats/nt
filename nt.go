// Main control flow of the program. Loads the config file and notebook concurrently.
// Reads command line input and calls the corresponding program after files have been
// loaded succesfully.

package main

import (
  notebook "github.com/TimoKats/nt/include/notebook"
  server "github.com/TimoKats/nt/include/server"
  . "github.com/TimoKats/nt/include/shared"

  "errors"
  "sync"
  "os"
)

func run(arguments Arguments) error {
  var pullErr error
  switch arguments.Command {
    // notebook related
    case Add:
      return notebook.AddNote(arguments)
    case List:
      return notebook.ListNotebook(arguments)
    case Clear:
      return notebook.ClearNotebook(arguments)
    case Move:
      return notebook.MoveNote(arguments)
    case Comment:
      return notebook.AddComment(arguments)
    case Tags:
      return notebook.ReadTags(arguments)
    case Modify:
      return notebook.ModifyNote(arguments)
    case Search:
      return notebook.SearchNote(arguments)
    // server related
    case Server:
      return server.RunServer()
    case Push:
      return server.PushNotebook(NtNotes)
    case Ping:
      return server.PingServer()
    case Pull:
      NtNotes, pullErr = server.PullNotebook()
      return pullErr
    default:
      return errors.New(HelpString)
  }
}

func setup() error {
  var wg sync.WaitGroup
  wg.Add(2)
  SetNtDir()
  go LoadConfig(&wg)
  go LoadNotebook(&wg)
  wg.Wait()
  return errors.Join(NtPathErr, NtConfigErr, NtNotesErr)
}

func main() {
  if err := setup(); err != nil { Error.Println(err); return }
  var arguments Arguments = ParseArgs(os.Args)
  var runErr error = run(arguments)
  if runErr != nil { Error.Println(runErr) }
}
