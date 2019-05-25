package merge

type Files []string

func (s Files) Len() int {
	return len(s)
}

func (s Files) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Files) Less(i, j int) bool {
	return s[i] < s[j]
}
