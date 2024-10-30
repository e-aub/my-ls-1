package main

import (
	"fmt"
	"log"
	"os"
)

type FileInfo struct {
	Name     string
	Detailes os.FileInfo
}

func ExecuteFlags(flags *Flags, results *Results) {
	for _, path := range flags.Paths {
		filesInfo, err := GetDirEntries(path, flags)

		if err != nil {
			results.Fail = append(results.Fail, fmt.Sprintf("ls: cannot access '%s': No such file or directory", path))
			continue
		}

		if !flags.Dash_a && !flags.Dash_A {
			filesInfo = DeleteHiddenFileNames(filesInfo)
		}

		if flags.Dash_a {
			filesInfo = append(filesInfo, FileInfo{Name: "."}, FileInfo{Name: ".."})
		}

		if !flags.Dash_t && !flags.Dash_S {
			SortNamesByAlphaOrder(filesInfo)
		}

		if flags.Dash_S || flags.Dash_t || flags.Dash_l {
			filesInfo, err = GetFilesInfos(path, filesInfo)
			if err != nil {
				log.Fatalln(err)
			}
		}

		if flags.Dash_t {
			SortBySize(filesInfo)
		} else if flags.Dash_S {
			SortByModTime(filesInfo)

		}

		for _, info := range filesInfo {
			fmt.Println(info.Name, info.Detailes.Size())
		}
	}
}
