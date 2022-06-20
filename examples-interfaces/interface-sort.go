package revstrings

type RevStr []string

func (s RevStr) Len() int           { return len(s) }
func (s RevStr) Less(i, j int) bool { return s[i] < s[j] }
func (s RevStr) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
