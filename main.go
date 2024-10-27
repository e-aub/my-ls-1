package main

import (
	"fmt"
	"os"
)

type Flags struct {
	ModeOnlyPathsActivated bool
	Paths                  []string
	Dash_t                 bool
	Dash_S                 bool
	Dash_l                 bool
	Dash_R                 bool
	Dash_a                 bool
	Dash_r                 bool
	Dash_A                 bool
}

type Results struct {
	Succeded []string
	Fail     []string
}

type File struct {
	Name string
}

func main() {
	args := os.Args[1:]
	flags := Flags{}
	results := Results{}
	ArgsParser(&flags, &args)
	if len(flags.Paths) == 0 {
		flags.Paths = append(flags.Paths, ".")
	}
	ExecuteFlags(&flags, &results)
	fmt.Println(flags)
	fmt.Println(results)
}

// if command == "ls" {
// 	if command == "ls" && len(args) == 1 { // ls
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
