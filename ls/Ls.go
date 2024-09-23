package ls

import (
	"io/fs"
)

func Ls(entries []fs.DirEntry) []string {
	Name := []string{}
	for _, entry := range entries {
		// Exclude hidden files and directories
		if entry.Name()[0] == '.' {
			continue
		}

		if entry.IsDir() {
			Name = append(Name, boldBlue+entry.Name()+reset+"  ")
		} else {
			Name = append(Name, entry.Name()+"  ")
		}
	}
	// fmt.Println()
	return Name
}
