//go:build !windows && !wasip1
// +build !windows,!wasip1

package sftp

import (
	"io/fs"
	"os"
)

func (s *Server) openfile(path string, flag int, mode fs.FileMode) (file, error) {
	return os.OpenFile(path, flag, mode)
}

func (s *Server) lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (s *Server) stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
