package model

type student struct {
	Name  string
	grade float64 // 假设学生的分数不对外导出
}

func (s *student) getScore() float64 {
	return s.grade
}

func NewStudent(name string, grade float64) *student {
	return &student{
		Name:  name,
		grade: grade,
	}
}
