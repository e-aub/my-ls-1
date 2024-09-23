package ls

import (
	"fmt"
	"io/fs"
)

func CheckFlags(arr []string) {
	content := []fs.DirEntry{}
	command := arr[0]
	// flag := ""
	dirr := ""
	fileName := ""
	if len(arr) == 2 {
		dirr = arr[1]
		if command == "ls" && dirr != "" {
			content = CurrentDir(dirr)
			Newcontent := Ls(content)
			for i, con := range Newcontent {
				if content[i].IsDir() {
					fmt.Print(con)
				} else {
					fileName = arr[1]
					fmt.Println(fileName)
				}
			}
			fmt.Println()
		}
		// if command == "ls" && fileName != "" {
		// }
	}
}
