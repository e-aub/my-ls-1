package ls

import (
	"fmt"
	"io/fs"
)

func OneFlag(command, flag string, arr []string, entries []fs.DirEntry) {
	// arr := []string{}
	if command == "ls" {
		if command == "ls" && len(arr) == 1 { // ls
			list1 := Ls(entries)
			for _, ls := range list1 {
				fmt.Print(ls)
			}
			fmt.Println()
		} else if flag == "-l" { // ls -l
			list := Ls_l(entries)
			for _, item := range list {
				fmt.Printf("%s %d %s %s %d %s %s\n",
					item.Permission,
					item.Hard_Links,
					item.UserName,
					item.GroupName,
					item.Size,
					item.ModTime,
					item.Name)
			}
		} else if flag == "-a" { // ls -a
			list2 := Ls_a(entries)
			for _, lsa := range list2 {
				fmt.Print(lsa)
			}
			fmt.Println()
		} else if flag == "-r" { // ls -r
			Names := Ls(entries)
			RevNames := Ls_r(Names)
			for _, rev := range RevNames {
				fmt.Print(rev)
			}
			fmt.Println()

		} else if flag == "-R" {
			LS_R("ls", entries)
		} else {
			fmt.Printf("%s: Command not found\n", command)
		}
	}
}
