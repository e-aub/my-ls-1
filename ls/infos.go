package ls

import "io/fs"

type LS_l struct {
	Permission fs.FileMode
	Hard_Links uint64
	UserName   string
	GroupName  string
	Size       int64
	ModTime    string
	Name       string
}

var (
	boldBlue = "\033[1m\033[34m"
	reset    = "\033[0m"
)
