package utils

func GetMapKeys(m map[string]bool) []string {
	res := make([]string, len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}
	return res
}
