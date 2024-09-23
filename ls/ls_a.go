package ls

import (
	"io/fs"
)

func Ls_a(entries []fs.DirEntry) []string {
	Name := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			Name = append(Name, "\033[34m"+entry.Name()+"\033[0m"+"  ")
		} else {
			Name = append(Name, entry.Name()+"  ")
		}
	}
	// fmt.Println()
	return Name
}
