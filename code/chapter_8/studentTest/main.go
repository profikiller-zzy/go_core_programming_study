package main

import "fmt"

// 编写一个学生考试系统

// Student 学生共有属性
type Student struct {
	Name  string
	Age   int
	Score int
}

// ShowInfo 显示学生的成绩
func (s *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score int) {
	//业务判断
	s.Score = score
}

// Pupil 小学生
type Pupil struct {
	Student //嵌入了 Student 匿名结构体
	Name    string
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

// Graduate 大学生, 研究生。。
type Graduate struct {
	Student //嵌入了 Student 匿名结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}

func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Name = "jack"
	fmt.Println(pupil.Name)
	fmt.Println(pupil.Student.Name)
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
}
