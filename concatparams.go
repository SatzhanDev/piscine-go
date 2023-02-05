package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, v := range args {
		str += string(v)
		if i != len(args)-1 {
			str += "\n"
		}
	}
	return str
}
