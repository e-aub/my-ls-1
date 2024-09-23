package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"syscall"
)

// type Group struct {
// 	gid  uint32
// 	name string
// 	errr error
// }

func main() {
	// var gg *Group
	arr := os.Args[1:]
	if len(arr) < 1 {
		log.Fatal("CHeck your argss")
		return
	}
	command := arr[0]
	// flag := arr[1]
	// Read the current directory
	if command == "ls" {
		entries, err := os.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("total", len(entries))

		// Print the names of the files and directories
		for _, entry := range entries {
			fmt.Println(entry.Name())
			fmt.Println("--------------")
			info, err := entry.Info()
			if err != nil {
				fmt.Println("ERRROROROR")
				return
			}
			fmt.Println("Name:", info.Name())
			fmt.Println("Size:", info.Size())
			fmt.Println("modification Time:", info.ModTime())
			mode := info.Mode()
			fmt.Println("Permissions:", mode)
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok {
				fmt.Println("error no ok !")
				continue
			}
			fmt.Println("hard links number", stat.Nlink)
			// Retrieve user and group IDs
			uid := stat.Uid
			// gid := stat.Gid

			// Convert UID to username
			user, err := user.LookupId(fmt.Sprint(uid))
			if err != nil {
				fmt.Printf("Failed to lookup user for UID %d: %v\n", uid, err)
				continue
			}

			// Convert GID to group name
			// group, err := user.LookupGroupId(fmt.Sprint(gid))
			// gg, gg.errr := user.LookupGroup()
			// if gErr != nil {
			// 	fmt.Printf("Failed to lookup group for GID %d: %v\n", gid, err)
			// 	continue
			// }

			// Print file name, username, and group name
			fmt.Printf("%s: user=%s, group=%s\n", entry.Name(), user.Username)

			fmt.Println("----------------------")

		}
	} else {
		fmt.Printf("%s", command+" : Command not found"+"\n")
	}
}
