package ls

import (
	"io/fs"
	"log"
	"os"
)

func CurrentDir(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	return entries
}
