# NT
NT is a quick note taking tool for the CLI. Beside basic commands, it integrates with the clipboard and supports self-hosting using NTS (note taking server).


## Note taking
NT supports a number of commands that can be enriched with certain patterns.

### Commands
All commands are prefixed with nt. So for example: `nt add close the PRs :due:Fri :tag:work`

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
      <td>ls</td>
      <td>Lists the notes in your notebook.</td>
    </tr>
    <tr width="600px">
      <td>rm *id*</td>
      <td>Removes notes (all notes by default!).</td>
    </tr>
    <tr width="600px">
      <td>cmt *id*</td>
      <td>Adds comment to specific note.</td>
    </tr>
    <tr width="600px">
      <td>tags</td>
      <td>Lists the tags in the notebook.</td>
    </tr>
    <tr width="600px">
      <td>mod *id*</td>
      <td>Modifies a selected note. Same args as 'add'.</td>
    </tr>
    <tr width="600px">
      <td>mv *id*</td>
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

## Configuration
