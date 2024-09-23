package ls

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func LS_R(dir string, entries []fs.DirEntry) {
	err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Println(p)

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
