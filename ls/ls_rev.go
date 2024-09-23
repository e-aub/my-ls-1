package ls

func Ls_r(lst []string) []string {
	RevNames := []string{}
	// RevNames := sort.Sort(sort.Reverse(sort.StringSlice(lst)))
	for i := len(lst) - 1; i >= 0; i-- {
		RevNames = append(RevNames, lst[i])
	}
	// fmt.Println(fullPaths)
	return RevNames
}
