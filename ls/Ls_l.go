package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os/user"
	"syscall"
)

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
			entryName = boldBlue + entry.Name() + reset
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
