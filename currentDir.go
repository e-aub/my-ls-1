package main

import (
	"os"
)

func GetDirEntries(dir string, flags *Flags) ([]FileInfo, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	entries, _ := file.ReadDir(0)

	fileInfos := []FileInfo{}

	for _, entry := range entries {
		fileInfos = append(fileInfos, FileInfo{Name: entry.Name()})
	}
	return fileInfos, nil
}
