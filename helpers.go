package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func DeleteHiddenFileNames(filesInfo []FileInfo) []FileInfo {
	var result []FileInfo
	for i := 0; i < len(filesInfo); i++ {
		if len((filesInfo)[i].Name) > 0 && (filesInfo)[i].Name[0] != '.' {
			result = append(result, FileInfo{Name: filesInfo[i].Name})
		}
	}
	return result
}

func SortNamesByAlphaOrder(fileNames []FileInfo) {
	sort.Slice(fileNames, func(i, j int) bool {
		x := strings.TrimPrefix(fileNames[i].Name, ".")
		y := strings.TrimPrefix(fileNames[j].Name, ".")
		return x < y
	})
}

func GetFilesInfos(path string, fileInfos []FileInfo) ([]FileInfo, error) {
	for i := 0; i < len(fileInfos); i++ {
		completePath := ""
		if path == "." {
			completePath = "./" + fileInfos[i].Name
		} else {
			completePath = path + "/" + fileInfos[i].Name
		}
		temp, err := os.Stat(completePath)
		fileInfos[i].Detailes = temp
		if err != nil {
			return nil, err
		}
	}
	return fileInfos, nil
}

// func SortBy(fileNames []string, infos []fs.FileInfo, by string) ([]string, []fs.FileInfo) {
// 	sort.Slice(infos, func(i, j int) bool {
// 		var less bool
// 		if by == "size" {
// 			less = infos[i].Size() > infos[j].Size()
// 		} else if by == "modTime" {
// 			less = infos[i].ModTime().After(infos[j].ModTime())
// 		}

// 		if less {
// 			fileNames[i], fileNames[j] = fileNames[j], fileNames[i]
// 		}
// 		return less
// 	})
// 	return fileNames, infos
// }

func SortBySize(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Detailes.Size() > files[j].Detailes.Size()
	})

	fmt.Println(files)
}

func SortByModTime(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Detailes.ModTime().After(files[j].Detailes.ModTime())
	})
}
