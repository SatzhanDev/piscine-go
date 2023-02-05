package piscine

func Join(strs []string, sep string) string {
	str := strs[0]
	for _, v := range strs[1:] {
		str = str + sep + v
	}
	return str
}
