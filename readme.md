# NT
![example workflow](https://github.com/timokats/nt/actions/workflows/test.yaml/badge.svg)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-red.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![GitHub tag](https://img.shields.io/github/tag/TimoKats/nt?include_prereleases=&sort=semver&color=cyan)](https://github.com/TimoKats/nt/releases/)
[![stars - nt](https://img.shields.io/github/stars/TimoKats/nt?style=social)](https://github.com/TimoKats/nt)
[![forks - nt](https://img.shields.io/github/forks/TimoKats/nt?style=social)](https://github.com/TimoKats/nt) 

NT is a quick note taking tool for the CLI. Beside basic commands, it integrates with the clipboard and supports self-hosting using NTS (note taking server). It can be installed using go.

```
go install github.com/TimoKats/nt@latest
```

## Note taking
NT supports a number of commands that can be enriched with certain patterns. All commands are prefixed with nt. So for example: `nt add close the PRs :due:Fri :tag:work`

### Commands

<table>
  <thead>
    <tr>
      <th width="500px">Command</th>
      <th width="500px">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr width="600px">
      <td>add *text*</td>
      <td>Adds a note to your notebook.</td>
    </tr>
    <tr width="600px">
      <td>ls *IDs*</td>
      <td>Lists the notes in your notebook.</td>
    </tr>
    <tr width="600px">
      <td>rm *IDs*</td>
      <td>Removes notes (all notes by default!).</td>
    </tr>
    <tr width="600px">
      <td>cmt *IDs*</td>
      <td>Adds comment to specific note.</td>
    </tr>
    <tr width="600px">
      <td>tags</td>
      <td>Lists the tags in the notebook.</td>
    </tr>
    <tr width="600px">
      <td>mod *IDs*</td>
      <td>Modifies a selected note. Same args as 'add'.</td>
    </tr>
    <tr width="600px">
      <td>mv *IDs*</td>
      <td>Checks/Unchecks a note.</td>
    </tr>
    <tr width="600px">
      <td>s *query*</td>
      <td>Filters the notebook based on a query. Query applied to text.</td>
    </tr>
  </tbody>
</table>

### Enrichments

<table>
  <thead>
    <tr>
      <th width="250px">Enrichment</th>
      <th width="250px">Code</th>
      <th width="500px">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr width="500px">
      <td>Tags</td>
      <td>@tagname / :tag:tagname</td>
      <td>Can be inserted into *text* when using 'add'. Adds a tag to a note.</td>
    </tr>
    <tr width="500px">
      <td>Deadlines</td>
      <td>:due:Sun / :due:14-12</td>
      <td>Can be inserted into *text* when using 'add'. Finds the first datetime associated with a (week)day. Formats can be configured.</td>
    </tr>
    <tr width="500px">
      <td>Clipboard</td>
      <td>:c</td>
      <td>Inserts the content of your clipboard into a note when using add. Listing a single note also sends data to clipboard.</td>
    </tr>
    <tr width="500px">
      <td>ID flags</td>
      <td>--done / --old / --today / --all</td>
      <td>Adding these flags will select notes you can apply your command to. Note, 'old' means older than today. </td>
    </tr>
    <tr width="500px">
      <td>ID ranges/selection</td>
      <td> {1,2..n} / 1-4 </td>
      <td>Adding these right after your command will select a single note or a range of notes based on their id.</td>
    </tr>
  </tbody>
</table>

## Note hosting
Any installation of NT can double as a server (NTS). It uses a REST API with basic authentication to send information between the server and the client. To setup/interact with your server, you can use the following commands.

> [!IMPORTANT]
> Upon setup (run command) you define your username and password for the basic authentication. This information isn't stored anywhere on your system. However, you can terminate/restart your server safely without losing the contents of your notes.


### Commands

<table>
  <thead>
    <tr>
      <th width="250px">Command</th>
      <th width="250px">Endpoint</th>
      <th width="500px">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr width="600px">
      <td>run</td>
      <td></td>
      <td>Creates a prompt that starts NTS. By default it runs on port 8282.</td>
    </tr>
    <tr width="600px">
      <td>push</td>
      <td>/push</td>
      <td>Pushes a notebook from the client to the server.</td>
    </tr>
    <tr width="600px">
      <td>pull</td>
      <td>/pull</td>
      <td>Pulls a notebook from the server to the client.</td>
    </tr>
    <tr width="600px">
      <td>ping</td>
      <td>/ping</td>
      <td>Health check from the client to the server.</td>
    </tr>
  </tbody>
</table>

Note, the `run` command is used wherever you want your server to exist. It will also store the notebooks there. The other commands are used where you take your notes.

## Configuration
No configuration is nessesary to use NT. However, you can add a `config.toml` file in `~/.nt` that supports the following options.

```toml
[server]
url = "http://000.000.00.000"
port = ":8282"

[notebook]
width = 40
ls_default = "--all"
# I recommend against changing the supported date formats for now...
date_format = ["2006-01-02T15:04", "2006-01-02", "Jan 02", "2", "Mon"] 
```


