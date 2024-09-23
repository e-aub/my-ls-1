package ls

import (
	"io/fs"
)

func Ls_a(entries []fs.DirEntry) []string {
	// var fullPaths []string
	Name := []string{}
	Name = append(Name, boldBlue+"."+reset+"  ", boldBlue+".."+reset+"  ")
	for _, entry := range entries {
		// fullPath, _ := filepath.Abs(entry.Name())
		// fullPaths = append(fullPaths, fullPath)
		if entry.IsDir() {
			Name = append(Name, boldBlue+entry.Name()+reset+"  ")
		} else {
			Name = append(Name, entry.Name()+"  ")
		}
	}

	// fmt.Println(fullPaths)
	return Name
}
