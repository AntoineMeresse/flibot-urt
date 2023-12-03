package utils

func CleanEmptyElements(datas []string) []string { 
	var res []string
	for _, value := range(datas) {
		if value != "" {
			res = append(res, value)
		}
	}
	return res
}