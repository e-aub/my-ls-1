package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"syscall"
)

type LS_l struct {
	Permission fs.FileMode
	Hard_Links uint64
	UserName   string
	GroupName  string
	Size       int64
	ModTime    string
	Name       string
}

func main() {
	// Check for command-line arguments
	arr := os.Args[1:]
	if len(arr) < 1 {
		log.Fatal("Check your args")
		return
	}
	command := arr[0]
	flag := ""
	if len(arr) > 1 {
		flag = arr[1]
	}

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// Read the current directory
	if command == "ls" && flag == "-l" {
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
	} else if command == "ls" && len(arr) == 1 {
		for _, entry := range entries {
			// Exclude hidden files and directories
			if entry.Name()[0] == '.' {
				continue
			}

			if entry.IsDir() {
				fmt.Print("\033[34m" + entry.Name() + "\033[0m" + "  ")
			} else {
				fmt.Print(entry.Name() + "  ")
			}
		}
		fmt.Println()
	} else {
		fmt.Printf("%s: Command not found\n", command)
	}
}

func Ls_l(entries []fs.DirEntry) []LS_l {
	var list []LS_l // A slice to hold all the directory entries
	fmt.Println("total", len(entries))

	// Iterate over the directory entries
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println(err)
			continue
		}

		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			fmt.Println("error not ok!")
			continue
		}

		uid := stat.Uid
		currentUser, _ := user.Current()
		gid := currentUser.Gid

		// Convert UID to username
		usr, err := user.LookupId(fmt.Sprint(uid))
		if err != nil {
			fmt.Printf("Failed to lookup user for UID %d: %v\n", uid, err)
			continue
		}

		group, err := user.LookupGroupId(gid)
		if err != nil {
			log.Fatalf("Error looking up group: %v", err)
		}

		// Convert the time format
		LastTime := info.ModTime()
		LastTimeMod := LastTime.Format("Jan 02 15:04")

		// Collect the data necessary for the ls -l output
		userName := usr.Username
		groupName := group.Name
		size := info.Size()
		entryName := ""
		if entry.IsDir() {
			if entry.Name()[0] == '.' {
				continue
			}
			entryName = "\033[34m" + entry.Name() + "\033[0m"
		} else {
			entryName = info.Name()
		}
		mode := info.Mode()

		// Append the data to the list slice
		list = append(list, LS_l{
			Permission: mode,
			Hard_Links: stat.Nlink,
			UserName:   userName,
			GroupName:  groupName,
			Size:       size,
			ModTime:    LastTimeMod,
			Name:       entryName,
		})
	}

	// Return the slice containing all the directory entries
	return list
}
