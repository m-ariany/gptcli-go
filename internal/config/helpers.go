package config

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

type writeCloser struct {
	io.Writer
}

func (w writeCloser) Close() error {
	return nil
}

// LogOutput returns a WriteCloser where log outputs can be redirected to.
func LogOutput() (io.WriteCloser, error) {
	if err := ensureDirPath(DefaultLogFile, DefaultFileMod); err != nil {
		fmt.Println("Oops, could not create a log file. Logs will be discarded.")
		return writeCloser{io.Discard}, nil
	}

	mod := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(DefaultLogFile, mod, DefaultFileMod)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// ensureDirPath ensures a directory exist from the given path.
func ensureDirPath(path string, mod os.FileMode) error {
	return ensureFullPath(filepath.Dir(path), mod)
}

// ensureFullPath ensures a directory exist from the given path.
func ensureFullPath(path string, mod os.FileMode) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, mod)
	}

	return err
}

// mustWhoami retrieves current user identity or fails.
func mustWhoami() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.Username
}
