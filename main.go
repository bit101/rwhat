package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	longArgs  = map[string]string{}
	shortArgs = map[string]string{}
)

func main() {
	makeArgs()
	args := os.Args
	for i := 1; i < len(args); i++ {
		splain(args[i])
	}
}

func splain(arg string) {
	if strings.HasPrefix(arg, "--") {
		fmt.Printf("%s: %s\n", arg, longDesc(arg))
	} else if strings.HasPrefix(arg, "-") {
		a := strings.TrimPrefix(arg, "-")
		for _, r := range a {
			s := string(r)
			fmt.Printf("-%s: %s\n", s, shortDesc(s))
		}
	}
}

func longDesc(arg string) string {
	desc, ok := longArgs[strings.TrimPrefix(arg, "--")]
	if ok {
		return desc
	}
	return "** invalid argument **"
}

func shortDesc(arg string) string {
	desc, ok := shortArgs[arg]
	if ok {
		return desc
	}
	return "** invalid argument **"
}

func makeArgs() {
	makeLongArgs()
	makeShortArgs()
}

func makeLongArgs() {
	longArgs["verbose"] = "increase verbosity"
	longArgs["quiet"] = "suppress non-error messages"
	longArgs["no-motd"] = "suppress daemon-mode MOTD (see caveat)"
	longArgs["checksum"] = "skip based on checksum, not mod-time & size"
	longArgs["archive"] = "archive mode; equals -rlptgoD (no -H,-A,-X)\n    -H: preserve hard links\n    -A: preserve ACLs (implies -p)\n    -X: preserve extended attributes"
	longArgs["no-OPTION"] = "turn off an implied OPTION (e.g. --no-D)"
	longArgs["recursive"] = "recurse into directories"
	longArgs["relative"] = "use relative path names"
	longArgs["no-implied-dirs"] = "don't send implied dirs with --relative\n    --relative: use relative path names"
	longArgs["backup"] = "make backups (see --suffix & --backup-dir)\n    --suffix: backup suffix (default ~ w/o --backup-dir)\n    --backup-dir: make backups into hierarchy based in DIR"
	longArgs["backup-dir"] = "make backups into hierarchy based in DIR"
	longArgs["suffix"] = "backup suffix (default ~ w/o --backup-dir)"
	longArgs["update"] = "skip files that are newer on the receiver"
	longArgs["inplace"] = "update destination files in-place"
	longArgs["append"] = "append data onto shorter files"
	longArgs["append-verify"] = "--append w/old data in file checksum\n    --append: append data onto shorter files"
	longArgs["dirs"] = "transfer directories without recursing"
	longArgs["links"] = "copy symlinks as symlinks"
	longArgs["copy-links"] = "transform symlink into referent file/dir"
	longArgs["copy-unsafe-links"] = "only 'unsafe' symlinks are transformed"
	longArgs["safe-links"] = "ignore symlinks that point outside the tree"
	longArgs["copy-dirlinks"] = "transform symlink to dir into referent dir"
	longArgs["keep-dirlinks"] = "treat symlinked dir on receiver as dir"
	longArgs["hard-links"] = "preserve hard links"
	longArgs["perms"] = "preserve permissions"
	longArgs["executability"] = "preserve executability"
	longArgs["chmod"] = "affect file and/or directory permissions"
	longArgs["acls"] = "preserve ACLs (implies -p)"
	longArgs["xattrs"] = "preserve extended attributes"
	longArgs["owner"] = "preserve owner (super-user only)"
	longArgs["group"] = "preserve group"
	longArgs["devices"] = "preserve device files (super-user only)"
	longArgs["specials"] = "preserve special files"
	longArgs["times"] = "preserve modification times"
	longArgs["omit-dir-times"] = "omit directories from --times\n    --times: preserve modification times"
	longArgs["super"] = "receiver attempts super-user activities"
	longArgs["fake-super"] = "store/recover privileged attrs using xattrs"
	longArgs["sparse"] = "handle sparse files efficiently"
	longArgs["dry-run"] = "perform a trial run with no changes made"
	longArgs["whole-file"] = "copy files whole (w/o delta-xfer algorithm)"
	longArgs["one-file-system"] = "don't cross filesystem boundaries"
	longArgs["block-size"] = "force a fixed checksum block-size"
	longArgs["rsh"] = "specify the remote shell to use"
	longArgs["rsync-path"] = "specify the rsync to run on remote machine"
	longArgs["existing"] = "skip creating new files on receiver"
	longArgs["ignore-existing"] = "skip updating files that exist on receiver"
	longArgs["remove-source-files"] = "sender removes synchronized files (non-dir)"
	longArgs["del"] = "an alias for --delete-during\n    --delete-during: receiver deletes during xfer, not before"
	longArgs["delete"] = "delete extraneous files from dest dirs"
	longArgs["delete-before"] = "receiver deletes before transfer (default)"
	longArgs["delete-during"] = "receiver deletes during xfer, not before"
	longArgs["delete-delay"] = "find deletions during, delete after"
	longArgs["delete-after"] = "receiver deletes after transfer, not before"
	longArgs["delete-excluded"] = "also delete excluded files from dest dirs"
	longArgs["ignore-errors"] = "delete even if there are I/O errors"
	longArgs["force"] = "force deletion of dirs even if not empty"
	longArgs["max-delete"] = "don't delete more than NUM files"
	longArgs["max-size"] = "don't transfer any file larger than SIZE"
	longArgs["min-size"] = "don't transfer any file smaller than SIZE"
	longArgs["partial"] = "keep partially transferred files"
	longArgs["partial-dir"] = "put a partially transferred file into DIR"
	longArgs["delay-updates"] = "put all updated files into place at end"
	longArgs["prune-empty-dirs"] = "prune empty directory chains from file-list"
	longArgs["numeric-ids"] = "don't map uid/gid values by user/group name"
	longArgs["timeout"] = "set I/O timeout in seconds"
	longArgs["contimeout"] = "set daemon connection timeout in seconds"
	longArgs["ignore-times"] = "don't skip files that match size and time"
	longArgs["size-only"] = "skip files that match in size"
	longArgs["modify-window"] = "compare mod-times with reduced accuracy"
	longArgs["temp-dir"] = "create temporary files in directory DIR"
	longArgs["fuzzy"] = "find similar file for basis if no dest file"
	longArgs["compare-dest"] = "also compare received files relative to DIR"
	longArgs["copy-dest"] = "... and include copies of unchanged files"
	longArgs["link-dest"] = "hardlink to files in DIR when unchanged"
	longArgs["compress"] = "compress file data during the transfer"
	longArgs["compress-level"] = "explicitly set compression level"
	longArgs["skip-compress"] = "skip compressing files with suffix in LIST"
	longArgs["cvs-exclude"] = "auto-ignore files in the same way CVS does"
	longArgs["filter"] = "add a file-filtering RULE"
	longArgs["exclude"] = "exclude files matching PATTERN"
	longArgs["exclude-from"] = "read exclude patterns from FILE"
	longArgs["include"] = "don't exclude files matching PATTERN"
	longArgs["include-from"] = "read include patterns from FILE"
	longArgs["files-from"] = "read list of source-file names from FILE"
	longArgs["from0"] = "all *from/filter files are delimited by 0s"
	longArgs["protect-args"] = "no space-splitting; wildcard chars only"
	longArgs["address"] = "bind address for outgoing socket to daemon"
	longArgs["port"] = "specify double-colon alternate port number"
	longArgs["sockopts"] = "specify custom TCP options"
	longArgs["blocking-io"] = "use blocking I/O for the remote shell"
	longArgs["stats"] = "give some file-transfer stats"
	longArgs["8-bit-output"] = "leave high-bit chars unescaped in output"
	longArgs["human-readable"] = "output numbers in a human-readable format"
	longArgs["progress"] = "show progress during transfer"
	longArgs["itemize-changes"] = "output a change-summary for all updates"
	longArgs["out-format"] = "output updates using the specified FORMAT"
	longArgs["log-file"] = "log what we're doing to the specified FILE"
	longArgs["log-file-format"] = "log updates using the specified FMT"
	longArgs["password-file"] = "read daemon-access password from FILE"
	longArgs["list-only"] = "list the files instead of copying them"
	longArgs["bwlimit"] = "limit I/O bandwidth; KBytes per second"
	longArgs["write-batch"] = "write a batched update to FILE"
	longArgs["only-write-batch"] = "like --write-batch but w/o updating dest\n    --write-batch: write a batched update to FILE"
	longArgs["read-batch"] = "read a batched update from FILE"
	longArgs["protocol"] = "force an older protocol version to be used"
	longArgs["iconv"] = "request charset conversion of filenames"
	longArgs["checksum-seed"] = "set block/file checksum seed (advanced)"
	longArgs["ipv4"] = "prefer IPv4"
	longArgs["ipv6"] = "prefer IPv6"
	longArgs["version"] = "print version number"
	longArgs["help"] = "show this help"
	longArgs["daemon"] = "run as an rsync daemon"
	longArgs["address"] = "bind to the specified address"
	longArgs["bwlimit"] = "limit I/O bandwidth; KBytes per second"
	longArgs["config"] = "specify alternate rsyncd.conf file"
	longArgs["no-detach"] = "do not detach from the parent"
	longArgs["port"] = "listen on alternate port number"
	longArgs["log-file"] = "override the 'log file' setting"
	longArgs["log-file-format"] = "override the 'log format' setting"
	longArgs["sockopts"] = "specify custom TCP options"
	longArgs["verbose"] = "increase verbosity"
	longArgs["ipv4"] = "prefer IPv4"
	longArgs["ipv6"] = "prefer IPv6"
	longArgs["help"] = "show this help (if used after --daemon)\n--daemon: run as an rsync daemon"
}

func makeShortArgs() {
	shortArgs["v"] = "increase verbosity"
	shortArgs["q"] = "suppress non-error messages"
	shortArgs["c"] = "skip based on checksum, not mod-time & size"
	shortArgs["a"] = "archive mode; equals -rlptgoD (no -H,-A,-X)\n    -H: preserve hard links\n    -A: preserve ACLs (implies -p)\n    -X: preserve extended attributes"
	shortArgs["r"] = "recurse into directories"
	shortArgs["R"] = "use relative path names"
	shortArgs["b"] = "make backups (see --suffix & --backup-dir)\n    --suffix: backup suffix (default ~ w/o --backup-dir)\n    --backup-dir: make backups into hierarchy based in DIR"
	shortArgs["u"] = "skip files that are newer on the receiver"
	shortArgs["d"] = "transfer directories without recursing"
	shortArgs["l"] = "copy symlinks as symlinks"
	shortArgs["L"] = "transform symlink into referent file/dir"
	shortArgs["k"] = "transform symlink to dir into referent dir"
	shortArgs["K"] = "treat symlinked dir on receiver as dir"
	shortArgs["H"] = "preserve hard links"
	shortArgs["p"] = "preserve permissions"
	shortArgs["E"] = "preserve executability"
	shortArgs["A"] = "preserve ACLs (implies -p)\n    -p: preserve permissions"
	shortArgs["X"] = "preserve extended attributes"
	shortArgs["o"] = "preserve owner (super-user only)"
	shortArgs["g"] = "preserve group"
	shortArgs["D"] = "same as --devices --specials\n    --devices: preserve device files (super-user only)\n    --specials: preserve special files"
	shortArgs["t"] = "preserve modification times"
	shortArgs["O"] = "omit directories from --times\n    --times: preserve modification times"
	shortArgs["S"] = "handle sparse files efficiently"
	shortArgs["n"] = "perform a trial run with no changes made"
	shortArgs["W"] = "copy files whole (w/o delta-xfer algorithm)"
	shortArgs["x"] = "don't cross filesystem boundaries"
	shortArgs["B"] = "force a fixed checksum block-size"
	shortArgs["e"] = "specify the remote shell to use"
	shortArgs["m"] = "prune empty directory chains from file-list"
	shortArgs["I"] = "don't skip files that match size and time"
	shortArgs["T"] = "create temporary files in directory DIR"
	shortArgs["y"] = "find similar file for basis if no dest file"
	shortArgs["z"] = "compress file data during the transfer"
	shortArgs["C"] = "auto-ignore files in the same way CVS does"
	shortArgs["f"] = "add a file-filtering RULE"
	shortArgs["F"] = "same as --filter='dir-merge /.rsync-filter'\n    --filter: add a file-filtering RULE"
	shortArgs["0"] = "all *from/filter files are delimited by 0s"
	shortArgs["s"] = "no space-splitting; wildcard chars only"
	shortArgs["8"] = "leave high-bit chars unescaped in output"
	shortArgs["h"] = "output numbers in a human-readable format"
	shortArgs["P"] = "same as --partial --progress\n    --partial: keep partially transferred files\n    --progress: show progress during transfer"
	shortArgs["i"] = "output a change-summary for all updates"
	shortArgs["4"] = "prefer IPv4"
	shortArgs["6"] = "prefer IPv6"
	shortArgs["h"] = "show this help"
	shortArgs["v"] = "increase verbosity"
	shortArgs["4"] = "prefer IPv4"
	shortArgs["6"] = "prefer IPv6"
	shortArgs["h"] = "show this help (if used after --daemon)\n    --daemon: run as an rsync daemon"
}
