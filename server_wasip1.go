//go:build wasip1
// +build wasip1

package sftp

import (
	"io/fs"
	"os"
)

func (s *Server) openfile(path string, flag int, mode fs.FileMode) (file, error) {
	return nil, nil
}

func (s *Server) lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (s *Server) stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
