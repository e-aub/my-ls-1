package ls

import (
	"fmt"
	"io/fs"
)

func ParseFlags(entries []fs.DirEntry, flags string) {
	// Collect results from each flag
	var result []string
	var list []LS_l

	// if flags[0] == '-' {
	for _, flag := range flags[1:] {
		switch flag {
		case 'l':
			// Call the function for '-l' (detailed info)
			list = Ls_l(entries)
			result = nil // Replace with actual function
			// continue
		case 'a':
			list = nil
			result = append(result, Ls_a(entries)...) // Replace with actual function
			// continue
		case 'r':
			list = nil
			result = append(result, Ls_r(Ls(entries))...) // Replace with actual function
			// continue
		default:
			fmt.Println("Unknown flag:", string(flag))
			// continue
		}
	}
	// }
	// Print the combined result of all the flags processed
	fmt.Println(result, list)
}
