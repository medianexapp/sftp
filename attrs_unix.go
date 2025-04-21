//go:build darwin || dragonfly || freebsd || (!android && linux) || netbsd || openbsd || solaris || aix || js || zos || wasip1
// +build darwin dragonfly freebsd !android,linux netbsd openbsd solaris aix js zos wasip1

package sftp

import (
	"os"
	"syscall"
)

func fileStatFromInfoOs(fi os.FileInfo, flags *uint32, fileStat *FileStat) {
	if statt, ok := fi.Sys().(*syscall.Stat_t); ok {
		*flags |= sshFileXferAttrUIDGID
		fileStat.UID = statt.Uid
		fileStat.GID = statt.Gid
	}
}
