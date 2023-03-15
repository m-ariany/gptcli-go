package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

var (
	// DefaultLogFile represents the default chatgpt log file.
	DefaultLogFile string = filepath.Join(os.TempDir(), fmt.Sprintf("chatgpt-%s.log", mustWhoami()))
)

const (
	// DefaultDirMod default unix perms for chatgpt directory.
	DefaultDirMod os.FileMode = 0755
	// DefaultFileMod default unix perms for chatgpt files.
	DefaultFileMod os.FileMode = 0600
	// DefaultLogLevel for chatgpt logs.
	DefaultLogLevel zerolog.Level = zerolog.DebugLevel
	// DefaultMaxToken to limit model's response token
	DefaultMaxToken int = 500
)
