package main

import (
	"log"
	"os"

	"go.mod/ls"
)

func main() {
	// Check for command-line arguments
	arr := os.Args[1:]
	if len(arr) < 1 {
		log.Fatal("Check your args")
		return
	}
	command := arr[0]
	flag := ""
	// dirr := ""
	// fileName := ""
	if len(arr) == 2 {
		flag = arr[1]
		// dirr = arr[2]
	}
	entries := ls.CurrentDir(".")
	ls.OneFlag(command, flag, arr, entries)
	ls.CheckFlags(arr)
}

// if command == "ls" {
// 	if command == "ls" && len(arr) == 1 { // ls
// 		list1 := ls.Ls(entries)
// 		for _, ls := range list1 {
// 			fmt.Print(ls)
// 		}
// 		fmt.Println()
// 	} else if flag == "-l" { // ls -l
// 		list := ls.Ls_l(entries)
// 		for _, item := range list {
// 			fmt.Printf("%s %d %s %s %d %s %s\n",
// 				item.Permission,
// 				item.Hard_Links,
// 				item.UserName,
// 				item.GroupName,
// 				item.Size,
// 				item.ModTime,
// 				item.Name)
// 		}
// 	} else if flag == "-a" { // ls -a
// 		list2 := ls.Ls_a(entries)
// 		for _, lsa := range list2 {
// 			fmt.Print(lsa)
// 		}
// 		fmt.Println()
// 	} else if flag == "-r" { // ls -r
// 		Names := ls.Ls(entries)
// 		RevNames := ls.Ls_r(Names)
// 		for _, rev := range RevNames {
// 			fmt.Print(rev)
// 		}
// 		fmt.Println()

// 	} else if flag == "-R" {
// 		ls.LS_R("ls", entries)
// 	} else {
// 		fmt.Printf("%s: Command not found\n", command)
// 	}
// }
