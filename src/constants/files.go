package constants

import "os"

const (
	DirPerm  os.FileMode = 0777
	FilePerm os.FileMode = 0666
)

var (
	ApiBotDir = ".apibot"
)
