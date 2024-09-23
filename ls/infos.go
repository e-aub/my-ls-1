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
