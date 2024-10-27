package main

import (
	"fmt"
)

func ExecuteFlags(flags *Flags, results *Results) {
	for _, path := range flags.Paths {
		entries, err := GetDirEntries(path, flags)
		if err != nil {
			results.Fail = append(results.Fail, fmt.Sprintf("ls: cannot access '%s': No such file or directory", path))
			continue
		}

		fmt.Println(entries)
	}
}
