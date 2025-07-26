// Package data contains lists of option descriptions
package data

import (
	"fmt"
	"strings"
)

var (
	argDescriptions = map[string]string{}
	nestedArgList   = map[string][]string{}
)

func init() {
	makeArgDescriptions()
	makeAliases()
	makenestedArgList()
}

// GetDesc gets the description of a given argument.
func GetDesc(arg string) string {
	desc, ok := argDescriptions[arg]
	if ok {
		return desc
	}
	if strings.HasPrefix(arg, "--no-") {
		return "turn off an implied OPTION (e.g. --no-D)"
	}
	return "\033[31m** invalid argument **\033[0m"
}

// AddNestedArgs adds indented arg descriptions for any descriptions that refer to other args.
func AddNestedArgs(arg string, indent int) {
	prefix := ""
	for range indent {
		prefix += " "
	}
	args, ok := nestedArgList[arg]
	if ok {
		for _, nestedArg := range args {
			fmt.Printf("\033[32m%s%s: %s\033[0m\n", prefix, nestedArg, GetDesc(nestedArg))
		}
	}
}

func makeArgDescriptions() {
	argDescriptions["--8-bit-output"] = "leave high-bit chars unescaped in output"
	argDescriptions["--acls"] = "preserve ACLs (implies --perms)"
	argDescriptions["--address"] = "bind to the specified address, or bind address for outgoing socket to daemon"
	argDescriptions["--append"] = "append data onto shorter files"
	argDescriptions["--append-verify"] = "--append w/old data in file checksum"
	argDescriptions["--archive"] = "archive mode is -rlptgoD (no -A,-X,-U,-N,-H)"
	argDescriptions["--atimes"] = "preserve access (use) times"
	argDescriptions["--backup"] = "make backups (see --suffix & --backup-dir)"
	argDescriptions["--backup-dir"] = "make backups into hierarchy based in DIR"
	argDescriptions["--block-size"] = "force a fixed checksum block-size"
	argDescriptions["--blocking-io"] = "use blocking I/O for the remote shell"
	argDescriptions["--bwlimit"] = "limit socket I/O bandwidth"
	argDescriptions["--checksum"] = "skip based on checksum, not mod-time & size"
	argDescriptions["--checksum-choice"] = "choose the checksum algorithm (aka --cc)"
	argDescriptions["--checksum-seed"] = "set block/file checksum seed (advanced)"
	argDescriptions["--chmod"] = "affect file and/or directory permissions"
	argDescriptions["--chown"] = "simple username/groupname mapping"
	argDescriptions["--compare-dest"] = "also compare destination files relative to DIR"
	argDescriptions["--compress"] = "compress file data during the transfer"
	argDescriptions["--compress-choice"] = "choose the compression algorithm (aka --zc)"
	argDescriptions["--compress-level"] = "explicitly set compression level (aka --zl)"
	argDescriptions["--config"] = "specify alternate rsyncd.conf file"
	argDescriptions["--contimeout"] = "set daemon connection timeout in seconds"
	argDescriptions["--copy-as"] = "specify user & optional group for the copy"
	argDescriptions["--copy-dest"] = "... and include copies of unchanged files"
	argDescriptions["--copy-devices"] = "copy device contents as a regular file"
	argDescriptions["--copy-dirlinks"] = "transform symlink to dir into referent dir"
	argDescriptions["--copy-links"] = "transform symlink into referent file/dir"
	argDescriptions["--copy-unsafe-links"] = "only 'unsafe' symlinks are transformed"
	argDescriptions["--crtimes"] = "preserve create times (newness)"
	argDescriptions["--cvs-exclude"] = "auto-ignore files in the same way CVS does"
	argDescriptions["--daemon"] = "run as an rsync daemon"
	argDescriptions["--debug"] = "fine-grained debug verbosity"
	argDescriptions["--del"] = "an alias for --delete-during"
	argDescriptions["--delay-updates"] = "put all updated files into place at end"
	argDescriptions["--delete"] = "delete extraneous files from dest dirs"
	argDescriptions["--delete-after"] = "receiver deletes after transfer, not during"
	argDescriptions["--delete-before"] = "receiver deletes before xfer, not during"
	argDescriptions["--delete-delay"] = "find deletions during, delete after"
	argDescriptions["--delete-during"] = "receiver deletes during the transfer"
	argDescriptions["--delete-excluded"] = "also delete excluded files from dest dirs"
	argDescriptions["--delete-missing-args"] = "delete missing source args from destination"
	argDescriptions["--devices"] = "preserve device files (super-user only)"
	argDescriptions["--dirs"] = "transfer directories without recursing"
	argDescriptions["--dparam"] = "override global daemon config parameter"
	argDescriptions["--dry-run"] = "perform a trial run with no changes made"
	argDescriptions["--early-input"] = "use FILE for daemon's early exec input"
	argDescriptions["--exclude-from"] = "read exclude patterns from FILE"
	argDescriptions["--exclude"] = "exclude files matching PATTERN"
	argDescriptions["--executability"] = "preserve executability"
	argDescriptions["--existing"] = "skip creating new files on receiver"
	argDescriptions["--fake-super"] = "store/recover privileged attrs using xattrs"
	argDescriptions["--files-from"] = "read list of source-file names from FILE"
	argDescriptions["--filter"] = "add a file-filtering RULE"
	argDescriptions["--force"] = "force deletion of dirs even if not empty"
	argDescriptions["--from0"] = "all *-from/filter files are delimited by 0s"
	argDescriptions["--fsync"] = "fsync every written file"
	argDescriptions["--fuzzy"] = "find similar file for basis if no dest file"
	argDescriptions["--group"] = "preserve group"
	argDescriptions["--groupmap"] = "custom groupname mapping"
	argDescriptions["--hard-links"] = "preserve hard links"
	argDescriptions["--help"] = "show rsync help"
	argDescriptions["--human-readable"] = "output numbers in a human-readable format"
	argDescriptions["--iconv"] = "request charset conversion of filenames"
	argDescriptions["--ignore-errors"] = "delete even if there are I/O errors"
	argDescriptions["--ignore-existing"] = "skip updating files that exist on receiver"
	argDescriptions["--ignore-missing-args"] = "ignore missing source args without error"
	argDescriptions["--ignore-times"] = "don't skip files that match size and time"
	argDescriptions["--include-from"] = "read include patterns from FILE"
	argDescriptions["--include"] = "don't exclude files matching PATTERN"
	argDescriptions["--info"] = "fine-grained informational verbosity"
	argDescriptions["--inplace"] = "update destination files in-place"
	argDescriptions["--ipv4"] = "prefer IPv4"
	argDescriptions["--ipv6"] = "prefer IPv6"
	argDescriptions["--itemize-changes"] = "output a change-summary for all updates"
	argDescriptions["--keep-dirlinks"] = "treat symlinked dir on receiver as dir"
	argDescriptions["--link-dest"] = "hardlink to files in DIR when unchanged"
	argDescriptions["--links"] = "copy symlinks as symlinks"
	argDescriptions["--list-only"] = "list the files instead of copying them"
	argDescriptions["--log-file-format"] = "log updates using the specified FMT, or override the 'log format' setting"
	argDescriptions["--log-file"] = "log what we're doing to the specified FILE, or override the 'log file' setting"
	argDescriptions["--max-alloc"] = "change a limit relating to memory alloc"
	argDescriptions["--max-delete"] = "don't delete more than NUM files"
	argDescriptions["--max-size"] = "don't transfer any file larger than SIZE"
	argDescriptions["--min-size"] = "don't transfer any file smaller than SIZE"
	argDescriptions["--mkpath"] = "create destination's missing path components"
	argDescriptions["--modify-window"] = "set the accuracy for mod-time comparisons"
	argDescriptions["--munge-links"] = "munge symlinks to make them safe & unusable"
	argDescriptions["--no-detach"] = "do not detach from the parent"
	argDescriptions["--no-implied-dirs"] = "don't send implied dirs with --relative"
	argDescriptions["--no-motd"] = "suppress daemon-mode MOTD"
	argDescriptions["--numeric-ids"] = "don't map uid/gid values by user/group name"
	argDescriptions["--old-args"] = "disable the modern arg-protection idiom"
	argDescriptions["--old-dirs"] = "old-d      works like --dirs when talking to old rsync"
	argDescriptions["--omit-dir-times"] = "omit directories from --times"
	argDescriptions["--omit-link-times"] = "omit symlinks from --times"
	argDescriptions["--one-file-system"] = "don't cross filesystem boundaries"
	argDescriptions["--only-write-batch"] = "like --write-batch but w/o updating dest"
	argDescriptions["--open-noatime"] = "avoid changing the atime on opened files"
	argDescriptions["--out-format"] = "output updates using the specified FORMAT"
	argDescriptions["--outbuf"] = "set out buffering to None, Line, or Block"
	argDescriptions["--owner"] = "preserve owner (super-user only)"
	argDescriptions["--partial"] = "keep partially transferred files"
	argDescriptions["--partial-dir"] = "put a partially transferred file into DIR"
	argDescriptions["--password-file"] = "read daemon-access password from FILE"
	argDescriptions["--perms"] = "preserve permissions"
	argDescriptions["--port"] = "listen on alternate port number"
	argDescriptions["--port"] = "specify double-colon alternate port number"
	argDescriptions["--preallocate"] = "allocate dest files before writing them"
	argDescriptions["--progress"] = "show progress during transfer"
	argDescriptions["--protocol"] = "force an older protocol version to be used"
	argDescriptions["--prune-empty-dirs"] = "prune empty directory chains from file-list"
	argDescriptions["--quiet"] = "suppress non-error messages"
	argDescriptions["--read-batch"] = "read a batched update from FILE"
	argDescriptions["--recursive"] = "recurse into directories"
	argDescriptions["--relative"] = "use relative path names"
	argDescriptions["--remote-option"] = "send OPTION to the remote side only"
	argDescriptions["--remove-source-files"] = "sender removes synchronized files (non-dir)"
	argDescriptions["--rsh"] = "specify the remote shell to use"
	argDescriptions["--rsync-path"] = "specify the rsync to run on remote machine"
	argDescriptions["--safe-links"] = "ignore symlinks that point outside the tree"
	argDescriptions["--secluded-args"] = "use the protocol to safely send the args"
	argDescriptions["--size-only"] = "skip files that match in size"
	argDescriptions["--skip-compress"] = "skip compressing files with suffix in LIST"
	argDescriptions["--sockopts"] = "specify custom TCP options"
	argDescriptions["--sparse"] = "turn sequences of nulls into sparse blocks"
	argDescriptions["--specials"] = "preserve special files"
	argDescriptions["--stats"] = "give some file-transfer stats"
	argDescriptions["--stderr"] = "change stderr output mode (default: errors)"
	argDescriptions["--stop-after"] = "Stop rsync after MINS minutes have elapsed"
	argDescriptions["--stop-at"] = "Stop rsync at the specified point in time"
	argDescriptions["--suffix"] = "backup suffix (default ~ w/o --backup-dir)"
	argDescriptions["--super"] = "receiver attempts super-user activities"
	argDescriptions["--temp-dir"] = "create temporary files in directory DIR"
	argDescriptions["--timeout"] = "set I/O timeout in seconds"
	argDescriptions["--times"] = "preserve modification times"
	argDescriptions["--trust-sender"] = "trust the remote sender's file list"
	argDescriptions["--update"] = "skip files that are newer on the receiver"
	argDescriptions["--usermap"] = "custom username mapping"
	argDescriptions["--verbose"] = "increase verbosity"
	argDescriptions["--version"] = "print the version + other info and exit"
	argDescriptions["--whole-file"] = "copy files whole (w/o delta-xfer algorithm)"
	argDescriptions["--write-batch"] = "write a batched update to FILE"
	argDescriptions["--write-devices"] = "write to devices as files (implies --inplace)"
	argDescriptions["--xattrs"] = "preserve extended attributes"
	argDescriptions["-D"] = "same as --devices --specials"
	argDescriptions["-F"] = "same as --filter /.rsync-filter'"
	argDescriptions["-P"] = "same as --partial --progress"
	argDescriptions["-h"] = "show rsync help if used alone, otherwise output numbers in a human-readable format"
}

func makeAliases() {
	argDescriptions["--old-d"] = argDescriptions["--old-dirs"]
	argDescriptions["-0"] = argDescriptions["--from0"]
	argDescriptions["-4"] = argDescriptions["--ipv4"]
	argDescriptions["-6"] = argDescriptions["--ipv6"]
	argDescriptions["-8"] = argDescriptions["--8-bit-output"]
	argDescriptions["-@"] = argDescriptions["--modify-window"]
	argDescriptions["-A"] = argDescriptions["--acls"]
	argDescriptions["-B"] = argDescriptions["--block-size"]
	argDescriptions["-C"] = argDescriptions["--cvs-exclude"]
	argDescriptions["-E"] = argDescriptions["--executability"]
	argDescriptions["-H"] = argDescriptions["--hard-links"]
	argDescriptions["-I"] = argDescriptions["--ignore-times"]
	argDescriptions["-J"] = argDescriptions["--omit-link-times"]
	argDescriptions["-K"] = argDescriptions["--keep-dirlinks"]
	argDescriptions["-L"] = argDescriptions["--copy-links"]
	argDescriptions["-M"] = argDescriptions["--dparam"]
	argDescriptions["-M"] = argDescriptions["--remote-option"]
	argDescriptions["-N"] = argDescriptions["--crtimes"]
	argDescriptions["-O"] = argDescriptions["--omit-dir-times"]
	argDescriptions["-R"] = argDescriptions["--relative"]
	argDescriptions["-S"] = argDescriptions["--sparse"]
	argDescriptions["-T"] = argDescriptions["--temp-dir"]
	argDescriptions["-U"] = argDescriptions["--atimes"]
	argDescriptions["-V"] = argDescriptions["--version"]
	argDescriptions["-W"] = argDescriptions["--whole-file"]
	argDescriptions["-X"] = argDescriptions["--xattrs"]
	argDescriptions["-a"] = argDescriptions["--archive"]
	argDescriptions["-b"] = argDescriptions["--backup"]
	argDescriptions["-c"] = argDescriptions["--checksum"]
	argDescriptions["-d"] = argDescriptions["--dirs"]
	argDescriptions["-e"] = argDescriptions["--rsh"]
	argDescriptions["-f"] = argDescriptions["--filter"]
	argDescriptions["-g"] = argDescriptions["--group"]
	argDescriptions["-h"] = argDescriptions["--human-readable"]
	argDescriptions["-i"] = argDescriptions["--itemize-changes"]
	argDescriptions["-k"] = argDescriptions["--copy-dirlinks"]
	argDescriptions["-l"] = argDescriptions["--links"]
	argDescriptions["-m"] = argDescriptions["--prune-empty-dirs"]
	argDescriptions["-n"] = argDescriptions["--dry-run"]
	argDescriptions["-o"] = argDescriptions["--owner"]
	argDescriptions["-p"] = argDescriptions["--perms"]
	argDescriptions["-q"] = argDescriptions["--quiet"]
	argDescriptions["-r"] = argDescriptions["--recursive"]
	argDescriptions["-s"] = argDescriptions["--secluded-args"]
	argDescriptions["-t"] = argDescriptions["--times"]
	argDescriptions["-u"] = argDescriptions["--update"]
	argDescriptions["-v"] = argDescriptions["--verbose"]
	argDescriptions["-x"] = argDescriptions["--one-file-system"]
	argDescriptions["-y"] = argDescriptions["--fuzzy"]
	argDescriptions["-z"] = argDescriptions["--compress"]
}

func makenestedArgList() {
	nestedArgList["--acls"] = []string{"--perms"}
	nestedArgList["--append-verify"] = []string{"--append"}
	nestedArgList["--archive"] = []string{"-A", "-X", "-U", "-N", "-H"}
	nestedArgList["--backup"] = []string{"--suffix", "--backup-dir"}
	nestedArgList["--del"] = []string{"--delete-during"}
	nestedArgList["--no-implied-dirs"] = []string{"--relative"}
	nestedArgList["--old-dirs"] = []string{"--dirs"}
	nestedArgList["--omit-dir-times"] = []string{"--times"}
	nestedArgList["--omit-link-times"] = []string{"--times"}
	nestedArgList["--only-write-batch"] = []string{"--write-batch"}
	nestedArgList["--write-devices"] = []string{"--inplace"}
	nestedArgList["-D"] = []string{"--devices", "--specials"}
	nestedArgList["-F"] = []string{"--filter"}
	nestedArgList["-P"] = []string{"--progress", "--partial"}

	nestedArgList["-A"] = nestedArgList["--acls"]
	nestedArgList["-O"] = nestedArgList["--omit-dir-times"]
	nestedArgList["-a"] = nestedArgList["--archive"]
	nestedArgList["-b"] = nestedArgList["--backup"]
}
