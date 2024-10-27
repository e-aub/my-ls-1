package main

import (
	"os"
	"sort"
	"strings"
)

func GetDirEntries(dir string, flags *Flags) ([]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	entries, _ := file.ReadDir(0)

	fileNames := []string{}

	for _, entry := range entries {
		fileNames = append(fileNames, entry.Name())
	}
	if !flags.Dash_a && !flags.Dash_A {
		fileNames = DeleteHiddenFileNames(fileNames)
	}

	if flags.Dash_a {
		fileNames = append(fileNames, ".", "..")
	}
	SortNamesByAlphaOrder(fileNames)
	return fileNames, nil
}

func DeleteHiddenFileNames(fileNames []string) []string {
	var result []string
	for i := 0; i < len(fileNames); i++ {
		if len((fileNames)[i]) > 0 && (fileNames)[i][0] != '.' {
			result = append(result, (fileNames)[i])
		}
	}
	return result
}

func SortNamesByAlphaOrder(fileNames []string) {
	sort.Slice(fileNames, func(i, j int) bool {
		x := strings.TrimPrefix(fileNames[i], ".")
		y := strings.TrimPrefix(fileNames[j], ".")
		return x < y
	})
}
