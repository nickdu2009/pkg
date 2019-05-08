package xstring


func Include(strs []string, str string) bool  {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}