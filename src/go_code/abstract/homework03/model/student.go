package model

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
