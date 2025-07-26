`rwhat` is a command line utility to explain what a given rsync command will do.

## Use cases:

- You read a tutorial or StackOverflow answer suggesting you run rsync with the following arguments.
- You have a script you wrote weeks, months, years ago that has an rsync command in it and you don't remember what it does.

## Problem / Solutions
rsync has several dozen commands and it's not always obvious what they do. Some solutions:

- Search through the help or man file for what each argument does. Ouch.
- Use [https://explainshell.com/](https://explainshell.com/) which is actually very good, but requires you to switch to a browser. It's also very verbose - may require scrolling through multiple screens of deep explanations.
- Replace `rsync` with `rwhat` in your command.

## Example

```
$ rwhat -urlmDb --progress --stats --delete /foo /bar
```

Result:

```
-u: skip files that are newer on the receiver
-r: recurse into directories
-l: copy symlinks as symlinks
-m: prune empty directory chains from file-list
-D: same as --devices --specials
    --devices: preserve device files (super-user only)
    --specials: preserve special files
-b: make backups (see --suffix & --backup-dir)
    --suffix: backup suffix (default ~ w/o --backup-dir)
    --backup-dir: make backups into hierarchy based in DIR
--progress: show progress during transfer
--stats: give some file-transfer stats
--delete: delete extraneous files from dest dirs
```

Note that if other arguments are referenced, those are also explained, indented.

## Caveats

- Data was obtained through the man file and processed. It is not dynamic and will need to be updated if anything changes.
- Created based on rsync 3.4.1 (released Jan 15, 2025) - Linux version.
- This isn't necessarily a good tool for discovering *how* to write an rsync command that does what you want, just for explaining what an existing one does.
- `rwhat` will not run `rsync` or any other external program. It merely parses the arguments you pass it and shows the help description for each one.
- As the [license](https://github.com/bit101/rwhat/blob/main/LICENSE) says, no warranty of any kind. This is just a tool to help you understand the rsync arguments. Please make sure you FULLY understand what you are doing before running any command on real data. 

## Install
- If you have Go installed, you can build with `go build` and put the `rwhat` executable in your path.
- Download an executable from the [releases](https://github.com/bit101/rwhat/releases) page and put that in your path.



