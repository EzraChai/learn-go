package model

//	Private type because the first letter of the struct is in small letter
type student struct {
	Name  string `json:"name"`
	score float64
}

//	Looks like getter and setter in JAVA
func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}

//	Set the score of the student struct
func (student *student) SetScore(score float64) {
	student.score = score
}

func (student *student) GetScore() float64 {
	return student.score
}
