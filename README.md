Simple list comparison tool

```
Usage:
  lstcmp [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  inboth      list lines that appear in both file A and file B
  notin       list lines in file A that are NOT IN file B

Flags:
  -h, --help   help for lstcmp

Use "lstcmp [command] --help" for more information about a command.
```

#### find lines in file A that are not in file B
```
Usage:
  lstcmp notin [FileA] [FileB]
```

#### fine lines that appear in both files
```
Usage:
  lstcmp inboth [FileA] [FileB]
```